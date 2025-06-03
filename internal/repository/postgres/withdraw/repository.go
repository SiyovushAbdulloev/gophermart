package auth

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/user"
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/withdraw"
	"github.com/SiyovushAbdulloev/gophermart/pkg/postgres"
	"time"
)

type WithDrawRepository struct {
	DB *postgres.Postgres
}

func New(db *postgres.Postgres) *WithDrawRepository {
	return &WithDrawRepository{DB: db}
}

func (repo *WithDrawRepository) List(userId int) ([]withdraw.WithDraw, error) {
	ctx := context.Background()
	tx, err := repo.DB.Pool.Begin(ctx)
	if err != nil {
		return []withdraw.WithDraw{}, err
	}
	defer tx.Rollback(ctx)

	//TODO: add pagination if enough time
	query := repo.DB.Builder.Select("id, user_id, order_id, points, created_at, updated_at").
		From("withdraws").
		Where(squirrel.Eq{"user_id": userId}).
		OrderBy("created_at DESC")

	sql, args, err := query.ToSql()
	if err != nil {
		return []withdraw.WithDraw{}, err
	}

	rows, err := tx.Query(ctx, sql, args...)
	if err != nil {
		return []withdraw.WithDraw{}, err
	}
	defer rows.Close()
	var withdraws []withdraw.WithDraw = []withdraw.WithDraw{}

	for rows.Next() {
		var w withdraw.WithDraw
		if err = rows.Scan(&w.Id, &w.UserId, &w.Order, &w.Sum, &w.CreatedAt, &w.UpdatedAt); err != nil {
			return []withdraw.WithDraw{}, err
		}

		withdraws = append(withdraws, w)
	}

	if err = rows.Err(); err != nil {
		return []withdraw.WithDraw{}, err
	}

	return withdraws, nil
}

func (repo *WithDrawRepository) Store(w withdraw.WithDraw, user user.User) (*withdraw.WithDraw, error) {
	ctx := context.Background()
	tx, err := repo.DB.Pool.Begin(ctx)
	if err != nil {
		return &withdraw.WithDraw{}, err
	}
	defer tx.Rollback(ctx)

	query := repo.DB.Builder.Update("balances").
		Set("amount", squirrel.Expr("amount - ?", w.Sum)).
		Where(squirrel.Eq{"user_id": user.Id})

	sql, args, err := query.ToSql()
	if err != nil {
		return &withdraw.WithDraw{}, err
	}

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return &withdraw.WithDraw{}, err
	}

	now := time.Now()
	queryW := repo.DB.Builder.Insert("withdraws").
		Columns("user_id", "order_id", "points", "created_at", "updated_at").
		Values(user.Id, w.Order, w.Sum, now, now).
		Suffix("RETURNING id, user_id, order_id, points, created_at, updated_at")

	sql, args, err = queryW.ToSql()
	if err != nil {
		return &withdraw.WithDraw{}, err
	}

	err = tx.QueryRow(ctx, sql, args...).Scan(&w.Id, &w.UserId, &w.Order, &w.Sum, &w.CreatedAt, &w.UpdatedAt)
	if err != nil {
		return &withdraw.WithDraw{}, err
	}

	if err = tx.Commit(ctx); err != nil {
		return &withdraw.WithDraw{}, err
	}

	return &w, nil
}

func (repo *WithDrawRepository) Sum(id int) (int, error) {
	ctx := context.Background()
	tx, err := repo.DB.Pool.Begin(ctx)
	if err != nil {
		return 0, err
	}

	defer tx.Rollback(ctx)

	query := repo.DB.Builder.Select("SUM(points) as sum").
		From("withdraws").
		Where(squirrel.Eq{"user_id": id})

	sql, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}

	var sum int
	err = tx.QueryRow(ctx, sql, args...).Scan(&sum)
	if err != nil {
		return 0, err
	}

	return sum, nil
}
