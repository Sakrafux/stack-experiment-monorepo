package db

import (
	"context"
	"database/sql"
	"errors"
	"strings"

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
	foundArticle.FavoritesCount = repo.getFavoritesCountForArticleIds(ctx, []int64{record.Id})[record.Id]
	foundArticle.TagList = repo.getAllTagsForArticleIds(ctx, []int64{record.Id})[record.Id]

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
	foundArticle.FavoritesCount = repo.getFavoritesCountForArticleIds(ctx, []int64{record.Id})[record.Id]
	foundArticle.TagList = repo.getAllTagsForArticleIds(ctx, []int64{record.Id})[record.Id]
	foundArticle.Favorited = repo.getFavoritedByUserForArticleIds(ctx, userId, []int64{record.Id})[record.Id]

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
	a.FavoritesCount = repo.getFavoritesCountForArticleIds(ctx, []int64{record.Id})[record.Id]
	a.TagList = repo.getAllTagsForArticleIds(ctx, []int64{record.Id})[record.Id]
	a.Favorited = repo.getFavoritedByUserForArticleIds(ctx, newArticle.AuthorId, []int64{record.Id})[record.Id]

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

func (repo *ArticleRepository) FindAllArticlesFiltered(ctx context.Context, filter *article.FilterParams) []*article.Article {
	var records []ArticleRecord
	err := repo.db.SelectContext(ctx, &records, `
		SELECT DISTINCT a.* FROM article a
		JOIN app_user u ON u.id = a.fk_author
		LEFT JOIN favorite_is_article_to_user fiatu ON fiatu.article_id = a.id
		LEFT JOIN app_user f ON f.id = fiatu.user_id
		LEFT JOIN tag_is_article_to_tag tiatt ON tiatt.article_id = a.id
		LEFT JOIN tag t ON t.id = tiatt.tag_id
		WHERE ($3::varchar IS NULL OR t.tag = $3)
			AND ($4::varchar IS NULL OR u.username = $4)
			AND ($5::varchar IS NULL OR f.username = $5)
		ORDER BY a.updated_at DESC
		LIMIT $1 OFFSET $2
	`, filter.Limit, filter.Offset, filter.Tag, filter.Author, filter.Favorited)
	if err != nil {
		panic(err)
	}

	articleIds := lo.Map(records, func(item ArticleRecord, index int) int64 {
		return item.Id
	})

	favoritesCounts := repo.getFavoritesCountForArticleIds(ctx, articleIds)
	var favoriteds map[int64]bool
	if filter.UserId != nil {
		favoriteds = repo.getFavoritedByUserForArticleIds(ctx, *filter.UserId, articleIds)
	}
	tags := repo.getAllTagsForArticleIds(ctx, articleIds)

	return lo.Map(records, func(item ArticleRecord, index int) *article.Article {
		foundArticle := toArticle(&item)

		foundArticle.FavoritesCount = favoritesCounts[item.Id]
		if filter.UserId != nil {
			foundArticle.Favorited = favoriteds[item.Id]
		}
		foundArticle.TagList = tags[item.Id]
		return foundArticle
	})
}

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

	articleIds := lo.Map(records, func(item ArticleRecord, index int) int64 {
		return item.Id
	})

	favoritesCounts := repo.getFavoritesCountForArticleIds(ctx, articleIds)
	favoriteds := repo.getFavoritedByUserForArticleIds(ctx, *filter.UserId, articleIds)
	tags := repo.getAllTagsForArticleIds(ctx, articleIds)

	return lo.Map(records, func(item ArticleRecord, index int) *article.Article {
		foundArticle := toArticle(&item)

		foundArticle.FavoritesCount = favoritesCounts[item.Id]
		foundArticle.Favorited = favoriteds[item.Id]
		foundArticle.TagList = tags[item.Id]
		return foundArticle
	})
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

type tag struct {
	ArticleId  int64  `db:"article_id"`
	TagsString string `db:"tags"`
}

func (repo *ArticleRepository) getAllTagsForArticleIds(ctx context.Context, articleIds []int64) map[int64][]string {
	var tags []tag
	err := repo.db.SelectContext(ctx, &tags, `
		SELECT a.id AS article_id, array_agg(t.tag ORDER BY t.tag) AS tags
		FROM article a
		LEFT JOIN tag_is_article_to_tag tiatt ON a.id = tiatt.article_id
		LEFT JOIN tag t ON t.id = tiatt.tag_id
		WHERE a.id = ANY($1)
		GROUP BY a.id
	`, articleIds)
	if err != nil {
		panic(err)
	}

	return lo.SliceToMap(tags, func(item tag) (int64, []string) {
		tagStrings := make([]string, 0)
		strippedTagsString := item.TagsString[1 : len(item.TagsString)-1]
		if len(strippedTagsString) > 0 && strippedTagsString != "NULL" {
			tagStrings = strings.Split(strippedTagsString, ",")
		}
		return item.ArticleId, tagStrings
	})
}

type favoriteCount struct {
	ArticleId      int64 `db:"article_id"`
	FavoritesCount int   `db:"favorites_count"`
}

func (repo *ArticleRepository) getFavoritesCountForArticleIds(ctx context.Context, articleIds []int64) map[int64]int {
	var favoritesCounts []favoriteCount
	err := repo.db.SelectContext(ctx, &favoritesCounts, `
		SELECT article_id, COUNT(*) AS "favorites_count" FROM favorite_is_article_to_user
		WHERE article_id = ANY($1)
		GROUP BY article_id
	`, articleIds)
	if err != nil {
		panic(err)
	}
	return lo.SliceToMap(favoritesCounts, func(item favoriteCount) (int64, int) {
		return item.ArticleId, item.FavoritesCount
	})
}

type favorited struct {
	ArticleId int64 `db:"article_id"`
	Favorited bool  `db:"favorited"`
}

func (repo *ArticleRepository) getFavoritedByUserForArticleIds(ctx context.Context, userId int64, articleIds []int64) map[int64]bool {
	var favoriteds []favorited
	err := repo.db.SelectContext(ctx, &favoriteds, `
		SELECT a.id AS "article_id", f.article_id IS NOT NULL AS favorited 
		FROM article a
		LEFT JOIN favorite_is_article_to_user f ON a.id = f.article_id
		WHERE a.id = ANY($1) AND f.user_id = $2
	`, articleIds, userId)
	if err != nil {
		panic(err)
	}
	return lo.SliceToMap(favoriteds, func(item favorited) (int64, bool) {
		return item.ArticleId, item.Favorited
	})
}
