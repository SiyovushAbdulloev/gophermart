package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/user"
	"github.com/SiyovushAbdulloev/gophermart/pkg/postgres"
	"github.com/SiyovushAbdulloev/gophermart/pkg/utils/hash"
	"github.com/jackc/pgx/v5"
	"time"
)

type AuthRepository struct {
	DB *postgres.Postgres
}

func New(db *postgres.Postgres) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (repo *AuthRepository) Register(user user.User) (*user.User, error) {
	hashPassword, err := hash.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	tx, err := repo.DB.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback(ctx)

	now := time.Now()
	query := repo.DB.Builder.Insert("users").
		Columns("email", "password", "created_at", "updated_at").
		Values(user.Email, hashPassword, now, now).
		Suffix("RETURNING id, created_at, updated_at")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = tx.QueryRow(context.Background(), sql, args...).Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *AuthRepository) GetUserByEmail(email string) (*user.User, error) {
	ctx := context.Background()
	tx, err := repo.DB.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := repo.DB.Builder.
		Select("id, email, password").
		From("users").
		Where(squirrel.Eq{"email": email})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var u user.User

	err = tx.QueryRow(context.Background(), sql, args...).Scan(&u.Id, &u.Email, &u.Password)
	if err != nil {
		fmt.Println(err.Error())
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	return &u, nil
}

func (repo *AuthRepository) GetUserById(id int) (*user.User, error) {
	ctx := context.Background()

	tx, err := repo.DB.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := repo.DB.Builder.
		Select("id, email, password").
		From("users").
		Where(squirrel.Eq{"id": id})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var u user.User
	err = tx.QueryRow(ctx, sql, args...).Scan(&u.Id, &u.Email, &u.Password)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
