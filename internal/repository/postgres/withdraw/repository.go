package auth

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/withdraw"
	"github.com/SiyovushAbdulloev/gophermart/pkg/postgres"
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
		if err = rows.Scan(&w.Id, &w.UserId, &w.OrderId, &w.Points, &w.CreatedAt, &w.UpdatedAt); err != nil {
			return []withdraw.WithDraw{}, err
		}

		withdraws = append(withdraws, w)
	}

	if err = rows.Err(); err != nil {
		return []withdraw.WithDraw{}, err
	}

	return withdraws, nil
}
