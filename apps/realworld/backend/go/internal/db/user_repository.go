package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Sakrafux/stack-experiment-monorepo/internal/domain/profile"
	"github.com/Sakrafux/stack-experiment-monorepo/internal/domain/user"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) Insert(ctx context.Context, user *user.User) *user.User {
	var record UserRecord
	err := repo.db.GetContext(ctx, &record, `
		INSERT INTO app_user (id, username, email, password, bio, image) 
		SELECT nextval('seq_user_id'), $1::varchar, $2::varchar, $3, $4, $5
		WHERE NOT EXISTS (SELECT 1 FROM app_user WHERE username = $1 OR email = $2)
		RETURNING *
	`, user.Username, user.Email, user.Password, user.Bio, user.Image)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	return toUser(&record)
}

func (repo *UserRepository) ExistsByUsernameOrEmail(ctx context.Context, username, email string) bool {
	var exists bool
	err := repo.db.GetContext(ctx, &exists, `
		SELECT true FROM app_user WHERE username = $1 OR email = $2
	`, username, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false
		}
		panic(err)
	}

	return true
}

func (repo *UserRepository) FindByEmail(ctx context.Context, email string) *user.User {
	var record UserRecord
	err := repo.db.GetContext(ctx, &record, `
		SELECT * FROM app_user WHERE email = $1
	`, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	return toUser(&record)
}

func (repo *UserRepository) FindByUsername(ctx context.Context, username string) *user.User {
	var record UserRecord
	err := repo.db.GetContext(ctx, &record, `
		SELECT * FROM app_user WHERE username = $1
	`, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	return toUser(&record)
}

func (repo *UserRepository) FindById(ctx context.Context, id int64) *user.User {
	var record UserRecord
	err := repo.db.GetContext(ctx, &record, `
		SELECT * FROM app_user WHERE id = $1
	`, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	return toUser(&record)
}

func (repo *UserRepository) Update(ctx context.Context, user *user.User) *user.User {
	var record UserRecord
	err := repo.db.GetContext(ctx, &record, `
		UPDATE app_user 
		SET username=$2, email=$3, password=$4, bio=$5, image=$6
		WHERE id = $1
		RETURNING *
	`, user.Id, user.Username, user.Email, user.Password, user.Bio, user.Image)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	return toUser(&record)
}

func (repo *UserRepository) FindProfileById(ctx context.Context, id int64) *profile.Profile {
	var record UserRecord
	err := repo.db.GetContext(ctx, &record, `
		SELECT * FROM app_user WHERE id = $1
	`, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	return toProfile(&record)
}

func (repo *UserRepository) FindProfileByIds(ctx context.Context, sourceId, targetId int64) *profile.Profile {
	var record ProfileRecord
	err := repo.db.GetContext(ctx, &record, `
		SELECT u.id, u.username, u.bio, u.image, f.followed_user_id IS NOT NULL as "following"
		FROM app_user u 
		LEFT JOIN (SELECT * FROM follow_is_user_to_user WHERE following_user_id = $1) f ON u.id = f.followed_user_id
		WHERE u.id = $2
	`, sourceId, targetId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	return fromProfileRecordToProfile(&record)
}

func (repo *UserRepository) FollowProfileByIds(ctx context.Context, sourceId, targetId int64) *profile.Profile {
	var record ProfileRecord
	err := repo.db.GetContext(ctx, &record, `
		SELECT u.id, u.username, u.bio, u.image, f.followed_user_id IS NOT NULL as "following"
		FROM app_user u 
		LEFT JOIN (SELECT * FROM follow_is_user_to_user WHERE following_user_id = $1) f ON u.id = f.followed_user_id
		WHERE u.id = $2
	`, sourceId, targetId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	p := fromProfileRecordToProfile(&record)

	if p.Following {
		return p
	}

	_, err = repo.db.ExecContext(ctx, `
		INSERT INTO follow_is_user_to_user (following_user_id, followed_user_id)
		VALUES ($1, $2)
	`, sourceId, targetId)
	if err != nil {
		panic(err)
	}

	p.Following = true
	return p
}

func (repo *UserRepository) UnfollowProfileByIds(ctx context.Context, sourceId, targetId int64) {
	_, err := repo.db.ExecContext(ctx, `
		DELETE FROM follow_is_user_to_user 
		WHERE following_user_id = $1 AND followed_user_id = $2
	`, sourceId, targetId)
	if err != nil {
		panic(err)
	}
}

func (repo *UserRepository) FindAllProfilesById(ctx context.Context, ids []int64, userId *int64) []*profile.Profile {
	if len(ids) == 0 {
		return make([]*profile.Profile, 0)
	}

	query, args, err := sqlx.In(`
		SELECT u.id, u.username, u.bio, u.image, f.followed_user_id IS NOT NULL as "following"
		FROM app_user u 
		LEFT JOIN (SELECT * FROM follow_is_user_to_user WHERE following_user_id = ?) f ON u.id = f.followed_user_id
		WHERE u.id IN (?)
	`, userId, ids)
	if err != nil {
		panic(err)
	}
	query = repo.db.Rebind(query)

	var records []ProfileRecord
	err = repo.db.SelectContext(ctx, &records, query, args...)
	if err != nil {
		panic(err)
	}

	return lo.Map(records, func(item ProfileRecord, index int) *profile.Profile {
		return fromProfileRecordToProfile(&item)
	})
}
