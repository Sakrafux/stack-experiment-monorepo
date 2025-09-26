package db

import (
	"context"
	"database/sql"
	"errors"

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

func (repo *ArticleRepository) FindAllTags(ctx context.Context) []string {
	var tags []string
	err := repo.db.SelectContext(ctx, &tags, "SELECT tag FROM tag")
	if err != nil {
		panic(err)
	}
	return tags
}

func (repo *ArticleRepository) InsertArticle(ctx context.Context, newArticle *article.NewArticle) *article.Article {
	tagRecords := make([]TagRecord, 0)
	for _, tag := range newArticle.TagList {
		var tagRecord TagRecord
		err := repo.db.GetContext(ctx, &tagRecord, `
			INSERT INTO tag (id, tag) VALUES (nextval('seq_tag_id'), $1)
			ON CONFLICT (tag) DO NOTHING
			RETURNING *
		`, tag)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				err = repo.db.GetContext(ctx, &tagRecord, "SELECT * FROM tag WHERE tag = $1", tag)
				if err != nil {
					panic(err)
				}
			} else {
				panic(err)
			}
		}
		tagRecords = append(tagRecords, tagRecord)
	}

	var record ArticleRecord
	err := repo.db.GetContext(ctx, &record, `
		INSERT INTO article (id, slug, title, description, body, fk_author) 
		SELECT nextval('seq_user_id'), $1::varchar, $2, $3, $4, $5
		WHERE NOT EXISTS (SELECT 1 FROM article WHERE slug = $1)
		RETURNING *
	`, newArticle.Slug, newArticle.Title, newArticle.Description, newArticle.Body, newArticle.AuthorId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	for _, tagRecord := range tagRecords {
		_, err = repo.db.ExecContext(ctx, `
			INSERT INTO tag_is_article_to_tag (article_id, tag_id)
			VALUES ($1, $2)
		`, record.Id, tagRecord.Id)
		if err != nil {
			panic(err)
		}
	}

	model := toArticle(&record)
	model.TagList = lo.Map(tagRecords, func(item TagRecord, index int) string {
		return item.Tag
	})

	return model
}

func (repo *ArticleRepository) FindArticle(ctx context.Context, slug string) *article.Article {
	var record ArticleRecord
	err := repo.db.GetContext(ctx, &record, `
		SELECT * FROM article WHERE slug = $1
	`, slug)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	foundArticle := toArticle(&record)
	foundArticle.FavoritesCount = repo.getFavoritesCountForArticleId(ctx, record.Id)
	foundArticle.TagList = repo.getAllTagsForArticleId(ctx, record.Id)

	return foundArticle
}

func (repo *ArticleRepository) FindArticleForUser(ctx context.Context, slug string, userId int64) *article.Article {
	var record ArticleRecord
	err := repo.db.GetContext(ctx, &record, `
		SELECT * FROM article WHERE slug = $1
	`, slug)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	foundArticle := toArticle(&record)
	foundArticle.FavoritesCount = repo.getFavoritesCountForArticleId(ctx, record.Id)
	foundArticle.TagList = repo.getAllTagsForArticleId(ctx, record.Id)
	foundArticle.Favorited = repo.getFavoritedByUserForArticleId(ctx, userId, record.Id)

	return foundArticle
}

func (repo *ArticleRepository) UpdateArticle(ctx context.Context, newArticle *article.NewArticle) *article.Article {
	var record ArticleRecord
	err := repo.db.GetContext(ctx, &record, `
		UPDATE article 
		SET title=$2, description=$3, body=$4
		WHERE slug = $1
		RETURNING *
	`, newArticle.Slug, newArticle.Title, newArticle.Description, newArticle.Body)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	a := toArticle(&record)
	a.FavoritesCount = repo.getFavoritesCountForArticleId(ctx, record.Id)
	a.TagList = repo.getAllTagsForArticleId(ctx, record.Id)
	a.Favorited = repo.getFavoritedByUserForArticleId(ctx, newArticle.AuthorId, record.Id)

	return a
}

func (repo *ArticleRepository) DeleteArticle(ctx context.Context, slug string) {
	_, err := repo.db.ExecContext(ctx, `
		DELETE FROM tag_is_article_to_tag 
		WHERE article_id = (
		    SELECT id FROM article WHERE slug = $1
		)
	`, slug)
	if err != nil {
		panic(err)
	}

	_, err = repo.db.ExecContext(ctx, `
		DELETE FROM article 
		WHERE slug = $1
	`, slug)
	if err != nil {
		panic(err)
	}
}

// TODO optimize
func (repo *ArticleRepository) FindAllArticlesFiltered(ctx context.Context, filter *article.FilterParams) []*article.Article {
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
		panic(err)
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
			panic(err)
		}
		foundArticle.FavoritesCount = favoriteCount

		if filter.UserId != nil {
			rows, err := repo.db.QueryxContext(ctx, `
				SELECT 1 FROM favorite_is_article_to_user WHERE article_id = $1 AND user_id = $2
			`, record.Id, filter.UserId)
			if err != nil {
				panic(err)
			}
			defer rows.Close()

			if rows.Next() {
				foundArticle.Favorited = true
			}
		}

		tags := repo.getAllTagsForArticleId(ctx, record.Id)
		foundArticle.TagList = tags

		articles = append(articles, foundArticle)
	}

	return articles
}

// TODO optimize
func (repo *ArticleRepository) FindAllArticlesFeed(ctx context.Context, filter *article.FilterParams) []*article.Article {
	var records []ArticleRecord
	err := repo.db.SelectContext(ctx, &records, `
		SELECT a.* FROM article a
		JOIN follow_is_user_to_user fiutu ON a.fk_author = fiutu.followed_user_id
		WHERE fiutu.following_user_id = $1
		ORDER BY a.updated_at DESC
		LIMIT $2 OFFSET $3
	`, filter.UserId, filter.Limit, filter.Offset)
	if err != nil {
		panic(err)
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
			panic(err)
		}
		foundArticle.FavoritesCount = favoriteCount

		rows, err := repo.db.QueryxContext(ctx, `
			SELECT 1 FROM favorite_is_article_to_user WHERE article_id = $1 AND user_id = $2
		`, record.Id, filter.UserId)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		if rows.Next() {
			foundArticle.Favorited = true
		}

		tags := repo.getAllTagsForArticleId(ctx, record.Id)
		foundArticle.TagList = tags

		articles = append(articles, foundArticle)
	}

	return articles
}

func (repo *ArticleRepository) CreateArticleFavorite(ctx context.Context, slug string, userId int64) {
	_, err := repo.db.ExecContext(ctx, `
		INSERT INTO favorite_is_article_to_user (article_id, user_id) 
		SELECT a.id, $2
		FROM article a
		WHERE a.slug = $1
	`, slug, userId)
	if err != nil {
		panic(err)
	}
}

func (repo *ArticleRepository) DeleteArticleFavorite(ctx context.Context, slug string, userId int64) {
	_, err := repo.db.ExecContext(ctx, `
		DELETE FROM favorite_is_article_to_user
		WHERE user_id = $2
		AND article_id = (
		    SELECT a.id
			FROM article a
			WHERE a.slug = $1
		)
	`, slug, userId)
	if err != nil {
		panic(err)
	}
}

func (repo *ArticleRepository) FindAllCommentsForArticle(ctx context.Context, slug string) []*article.Comment {
	var records []CommentRecord
	err := repo.db.SelectContext(ctx, &records, `
		SELECT c.* FROM comment c
		JOIN article a ON c.fk_article = a.id
		WHERE a.slug = $1
	`, slug)
	if err != nil {
		panic(err)
	}

	return lo.Map(records, func(item CommentRecord, index int) *article.Comment {
		return toComment(&item)
	})
}

func (repo *ArticleRepository) CreateArticleComment(ctx context.Context, slug string, userId int64, body string) *article.Comment {
	var record CommentRecord
	err := repo.db.GetContext(ctx, &record, `
		INSERT INTO comment (id, body, fk_author, fk_article)
		SELECT nextval('seq_comment_id'), $1, $2, a.id
		FROM article a
		WHERE a.slug = $3
		RETURNING *
	`, body, userId, slug)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}
	return toComment(&record)
}

func (repo *ArticleRepository) DeleteArticleComment(ctx context.Context, slug string, userId, id int64) {
	_, err := repo.db.ExecContext(ctx, `
		DELETE FROM comment
		WHERE id = $1
		  	AND fk_author = $2
			AND fk_article = (
				SELECT a.id
				FROM article a
				WHERE a.slug = $3
		)
	`, id, userId, slug)
	if err != nil {
		panic(err)
	}
}

func (repo *ArticleRepository) getAllTagsForArticleId(ctx context.Context, articleId int64) []string {
	var tags []string
	err := repo.db.SelectContext(ctx, &tags, `
		SELECT t.tag FROM tag t
		JOIN tag_is_article_to_tag tiatt on t.id = tiatt.tag_id
		WHERE tiatt.article_id = $1
		ORDER BY t.tag
	`, articleId)
	if err != nil {
		panic(err)
	}

	return tags
}

func (repo *ArticleRepository) getFavoritesCountForArticleId(ctx context.Context, articleId int64) int {
	var favoritesCount int
	err := repo.db.GetContext(ctx, &favoritesCount, `
		SELECT COUNT(*) FROM favorite_is_article_to_user
		WHERE article_id = $1
	`, articleId)
	if err != nil {
		panic(err)
	}
	return favoritesCount
}

func (repo *ArticleRepository) getFavoritedByUserForArticleId(ctx context.Context, userId, articleId int64) bool {
	var favorited bool
	err := repo.db.GetContext(ctx, &favorited, `
		SELECT 1 FROM favorite_is_article_to_user WHERE article_id = $1 AND user_id = $2
	`, articleId, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false
		}
		panic(err)
	}
	return true
}
