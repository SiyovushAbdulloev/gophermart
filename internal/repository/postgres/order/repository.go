package auth

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/order"
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/user"
	"github.com/SiyovushAbdulloev/gophermart/pkg/postgres"
)

type OrderRepository struct {
	DB *postgres.Postgres
}

func New(db *postgres.Postgres) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (repo *OrderRepository) Store(id int, u user.User) (*order.Order, error) {
	ctx := context.Background()
	tx, err := repo.DB.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback(ctx)
	var o order.Order

	query := repo.DB.Builder.Insert("orders").
		Columns("id", "user_id", "points", "status").
		Values(id, u.Id, 100, order.NewStatus).
		Suffix("RETURNING id, created_at, updated_at")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = tx.QueryRow(context.Background(), sql, args...).Scan(&o.Id, &o.CreatedAt, &o.UpdatedAt)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &o, nil
}

func (repo *OrderRepository) GetOrderById(id int) (*order.Order, error) {
	ctx := context.Background()
	tx, err := repo.DB.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := repo.DB.Builder.Select("id, user_id, points, status, created_at, updated_at").
		From("orders").
		Where(squirrel.Eq{"id": id})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var o order.Order
	err = tx.QueryRow(ctx, sql, args...).Scan(&o.Id, &o.UserId, &o.Points, &o.Status, &o.CreatedAt, &o.UpdatedAt)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &o, nil
}

func (repo *OrderRepository) List(userId int) ([]order.Order, error) {
	ctx := context.Background()
	tx, err := repo.DB.Pool.Begin(ctx)
	if err != nil {
		return []order.Order{}, err
	}
	defer tx.Rollback(ctx)

	//TODO: add pagination if enough time
	query := repo.DB.Builder.Select("id, user_id, points, status, created_at, updated_at").
		From("orders").
		Where(squirrel.Eq{"user_id": userId}).
		OrderBy("created_at DESC")

	sql, args, err := query.ToSql()
	if err != nil {
		return []order.Order{}, err
	}

	rows, err := tx.Query(ctx, sql, args...)
	if err != nil {
		return []order.Order{}, err
	}
	defer rows.Close()
	var orders []order.Order

	for rows.Next() {
		var o order.Order
		if err = rows.Scan(&o.Id, &o.UserId, &o.Points, &o.Status, &o.CreatedAt, &o.UpdatedAt); err != nil {
			return []order.Order{}, err
		}

		orders = append(orders, o)
	}

	if err = rows.Err(); err != nil {
		return []order.Order{}, err
	}

	return orders, nil
}
