package balance

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/SiyovushAbdulloev/gophermart/pkg/postgres"
)

type BalanceRepository struct {
	DB *postgres.Postgres
}

func New(db *postgres.Postgres) *BalanceRepository {
	return &BalanceRepository{DB: db}
}

func (repo *BalanceRepository) GetAmount(id int) (int, error) {
	ctx := context.Background()
	tx, err := repo.DB.Pool.Begin(ctx)
	if err != nil {
		return 0, err
	}

	defer tx.Rollback(ctx)

	query := repo.DB.Builder.Select("amount").
		From("balances").
		Where(squirrel.Eq{"user_id": id})

	sql, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}

	var amount int
	err = tx.QueryRow(ctx, sql, args...).Scan(&amount)
	if err != nil {
		return 0, err
	}

	if err = tx.Commit(ctx); err != nil {
		return 0, err
	}

	return amount, nil
}

func (repo *BalanceRepository) AddPoints(userID int, points int) error {
	ctx := context.Background()
	query := repo.DB.Builder.
		Update("balances").
		Set("amount", squirrel.Expr("amount + ?", points)).
		Where(squirrel.Eq{"user_id": userID})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = repo.DB.Pool.Exec(ctx, sql, args...)
	return err
}
