package db

import (
	"context"
	"fmt"

	"github.com/Sakrafux/stack-experiment-monorepo/internal/domain/user"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) Insert(ctx context.Context, user *user.User) (*user.User, error) {
	rows, err := repo.db.NamedQueryContext(ctx, `
		INSERT INTO app_user (id, username, email, password, bio, image) 
		SELECT nextval('seq_user_id'), :username, :email, :password, :bio, :image
		WHERE NOT EXISTS (SELECT 1 FROM app_user WHERE username = :username OR email = :email)
		RETURNING id
	`, fromUser(user))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id int64
	if rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("no user created")
	}

	var newUser UserRecord
	err = repo.db.GetContext(ctx, &newUser, "SELECT * FROM app_user WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return toUser(&newUser), nil
}

func (repo *UserRepository) ExistsByUsernameOrEmail(ctx context.Context, username, email string) (bool, error) {
	rows, err := repo.db.QueryxContext(ctx, `
		SELECT 1 FROM app_user WHERE username = $1 OR email = $2
	`, username, email)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if rows.Next() {
		return true, nil
	}
	return false, nil
}

func (repo *UserRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	var record UserRecord
	err := repo.db.QueryRowxContext(ctx, `
		SELECT * FROM app_user WHERE email = $1
	`, email).StructScan(&record)
	if err != nil {
		return nil, err
	}

	return toUser(&record), nil
}
