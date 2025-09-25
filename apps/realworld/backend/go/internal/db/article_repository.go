package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Sakrafux/stack-experiment-monorepo/internal/domain/article"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
)

type ArticleRepository struct {
	db *sqlx.DB
}

func NewArticleRepository(db *sqlx.DB) *ArticleRepository {
	return &ArticleRepository{db}
}

func (repo *ArticleRepository) FindAllTags(ctx context.Context) ([]string, error) {
	var tags []string
	err := repo.db.SelectContext(ctx, &tags, "SELECT tag FROM tag")
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (repo *ArticleRepository) InsertArticle(ctx context.Context, a *article.NewArticle) (*article.Article, error) {
	input := fromNewArticle(a)
	tagRecords := make([]TagRecord, 0)

	for _, tag := range a.TagList {
		var tagRecord TagRecord
		err := repo.db.Get(&tagRecord, `
			INSERT INTO tag (id, tag) VALUES (nextval('seq_tag_id'), $1)
			ON CONFLICT (tag) DO NOTHING
			RETURNING *
		`, tag)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				err = repo.db.Get(&tagRecord, "SELECT * FROM tag WHERE tag = $1", tag)
				if err != nil {
					return nil, err
				}
			} else {
				return nil, err
			}
		}
		tagRecords = append(tagRecords, tagRecord)
	}

	rows, err := repo.db.NamedQueryContext(ctx, `
		INSERT INTO article (id, slug, title, description, body, fk_author) 
		SELECT nextval('seq_user_id'), :slug, :title, :description, :body, :fk_author
		WHERE NOT EXISTS (SELECT 1 FROM article WHERE slug = :slug)
		RETURNING *
	`, input)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articleRecord ArticleRecord
	if rows.Next() {
		if err := rows.StructScan(&articleRecord); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("no article created")
	}

	for _, tagRecord := range tagRecords {
		_, err = repo.db.ExecContext(ctx, `
			INSERT INTO tag_is_article_to_tag (article_id, tag_id)
			VALUES ($1, $2)
		`, articleRecord.Id, tagRecord.Id)
		if err != nil {
			return nil, err
		}
	}

	model := toArticle(&articleRecord)
	model.TagList = lo.Map(tagRecords, func(item TagRecord, index int) string {
		return item.Tag
	})

	return model, nil
}

func (repo *ArticleRepository) FindArticle(ctx context.Context, slug string) (*article.Article, error) {
	var record ArticleRecord
	err := repo.db.QueryRowxContext(ctx, `
		SELECT * FROM article WHERE slug = $1
	`, slug).StructScan(&record)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	foundArticle := toArticle(&record)

	var favoriteCount int
	err = repo.db.Get(&favoriteCount, `
		SELECT COUNT(*) FROM favorite_is_article_to_user
		WHERE article_id = $1
	`, record.Id)
	if err != nil {
		return nil, err
	}
	foundArticle.FavoritesCount = favoriteCount

	tags, err := repo.getAllTagsForArticleId(ctx, record.Id)
	if err != nil {
		return nil, err
	}
	foundArticle.TagList = tags

	return foundArticle, nil
}

func (repo *ArticleRepository) FindArticleForUser(ctx context.Context, slug string, userId int64) (*article.Article, error) {
	var record ArticleRecord
	err := repo.db.QueryRowxContext(ctx, `
		SELECT * FROM article WHERE slug = $1
	`, slug).StructScan(&record)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	foundArticle := toArticle(&record)

	var favoriteCount int
	err = repo.db.Get(&favoriteCount, `
		SELECT COUNT(*) FROM favorite_is_article_to_user
		WHERE article_id = $1
	`, record.Id)
	if err != nil {
		return nil, err
	}
	foundArticle.FavoritesCount = favoriteCount

	rows, err := repo.db.QueryxContext(ctx, `
		SELECT 1 FROM favorite_is_article_to_user WHERE article_id = $1 AND user_id = $2
	`, record.Id, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		foundArticle.Favorited = true
	}

	tags, err := repo.getAllTagsForArticleId(ctx, record.Id)
	if err != nil {
		return nil, err
	}
	foundArticle.TagList = tags

	return foundArticle, nil
}

func (repo *ArticleRepository) UpdateArticle(ctx context.Context, article *article.NewArticle) (*article.Article, error) {
	result, err := repo.db.NamedExecContext(ctx, `
		UPDATE article 
		SET title=:title, description=:description, body=:body
		WHERE slug = :slug
	`, fromNewArticle(article))
	if err != nil {
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, fmt.Errorf("no article updated")
	}

	var record ArticleRecord
	err = repo.db.QueryRowxContext(ctx, `
		SELECT * FROM article WHERE slug = $1
	`, article.Slug).StructScan(&record)
	if err != nil {
		return nil, err
	}

	a := toArticle(&record)

	var favoriteCount int
	err = repo.db.Get(&favoriteCount, `
		SELECT COUNT(*) FROM favorite_is_article_to_user
		WHERE article_id = $1
	`, record.Id)
	if err != nil {
		return nil, err
	}
	a.FavoritesCount = favoriteCount

	rows, err := repo.db.QueryxContext(ctx, `
		SELECT 1 FROM favorite_is_article_to_user WHERE article_id = $1 AND user_id = $2
	`, record.Id, article.AuthorId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		a.Favorited = true
	}

	tags, err := repo.getAllTagsForArticleId(ctx, record.Id)
	if err != nil {
		return nil, err
	}
	a.TagList = tags

	return a, nil
}

func (repo *ArticleRepository) DeleteArticle(ctx context.Context, slug string) error {
	_, err := repo.db.ExecContext(ctx, `
		DELETE FROM tag_is_article_to_tag 
		WHERE article_id = (
		    SELECT id FROM article WHERE slug = $1
		)
	`, slug)

	_, err = repo.db.ExecContext(ctx, `
		DELETE FROM article 
		WHERE slug = $1
	`, slug)
	if err != nil {
		return err
	}

	return nil
}

func (repo *ArticleRepository) FindAllArticlesFiltered(ctx context.Context, filter *article.FilterParams) ([]*article.Article, error) {
	var records []ArticleRecord
	err := repo.db.SelectContext(ctx, &records, `
		SELECT DISTINCT a.* FROM article a
		JOIN app_user u ON u.id = a.fk_author
		LEFT JOIN favorite_is_article_to_user fiatu ON fiatu.article_id = a.id
		LEFT JOIN app_user f ON f.id = fiatu.user_id
		JOIN tag_is_article_to_tag tiatt ON tiatt.article_id = a.id
		JOIN tag t ON t.id = tiatt.tag_id
		WHERE ($3::varchar IS NULL OR t.tag = $3)
			AND ($4::varchar IS NULL OR u.username = $4)
			AND ($5::varchar IS NULL OR f.username = $5)
		ORDER BY a.updated_at DESC
		LIMIT $1 OFFSET $2
	`, filter.Limit, filter.Offset, filter.Tag, filter.Author, filter.Favorited)
	if err != nil {
		return nil, err
	}

	var articles []*article.Article
	for _, record := range records {
		foundArticle := toArticle(&record)

		var favoriteCount int
		err = repo.db.Get(&favoriteCount, `
			SELECT COUNT(*) FROM favorite_is_article_to_user
			WHERE article_id = $1
		`, record.Id)
		if err != nil {
			return nil, err
		}
		foundArticle.FavoritesCount = favoriteCount

		if filter.UserId != nil {
			rows, err := repo.db.QueryxContext(ctx, `
				SELECT 1 FROM favorite_is_article_to_user WHERE article_id = $1 AND user_id = $2
			`, record.Id, filter.UserId)
			if err != nil {
				return nil, err
			}
			defer rows.Close()

			if rows.Next() {
				foundArticle.Favorited = true
			}
		}

		tags, err := repo.getAllTagsForArticleId(ctx, record.Id)
		if err != nil {
			return nil, err
		}
		foundArticle.TagList = tags

		articles = append(articles, foundArticle)
	}

	return articles, nil
}

func (repo *ArticleRepository) FindAllArticlesFeed(ctx context.Context, filter *article.FilterParams) ([]*article.Article, error) {
	var records []ArticleRecord
	err := repo.db.SelectContext(ctx, &records, `
		SELECT a.* FROM article a
		JOIN follow_is_user_to_user fiutu ON a.fk_author = fiutu.followed_user_id
		WHERE fiutu.following_user_id = $1
		ORDER BY a.updated_at DESC
		LIMIT $2 OFFSET $3
	`, filter.UserId, filter.Limit, filter.Offset)
	if err != nil {
		return nil, err
	}

	var articles []*article.Article
	for _, record := range records {
		foundArticle := toArticle(&record)

		var favoriteCount int
		err = repo.db.Get(&favoriteCount, `
			SELECT COUNT(*) FROM favorite_is_article_to_user
			WHERE article_id = $1
		`, record.Id)
		if err != nil {
			return nil, err
		}
		foundArticle.FavoritesCount = favoriteCount

		rows, err := repo.db.QueryxContext(ctx, `
			SELECT 1 FROM favorite_is_article_to_user WHERE article_id = $1 AND user_id = $2
		`, record.Id, filter.UserId)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		if rows.Next() {
			foundArticle.Favorited = true
		}

		tags, err := repo.getAllTagsForArticleId(ctx, record.Id)
		if err != nil {
			return nil, err
		}
		foundArticle.TagList = tags

		articles = append(articles, foundArticle)
	}

	return articles, nil
}

func (repo *ArticleRepository) getAllTagsForArticleId(ctx context.Context, articleId int64) ([]string, error) {
	var tags []string
	err := repo.db.SelectContext(ctx, &tags, `
		SELECT t.tag FROM tag t
		JOIN tag_is_article_to_tag tiatt on t.id = tiatt.tag_id
		WHERE tiatt.article_id = $1
		ORDER BY t.tag
	`, articleId)
	if err != nil {
		return nil, err
	}

	return tags, nil
}
