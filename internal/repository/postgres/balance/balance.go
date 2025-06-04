package balance

import (
	"context"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/SiyovushAbdulloev/gophermart/pkg/postgres"
	"github.com/jackc/pgx/v5"
)

type BalanceRepository struct {
	DB *postgres.Postgres
}

func New(db *postgres.Postgres) *BalanceRepository {
	return &BalanceRepository{DB: db}
}

func (repo *BalanceRepository) GetAmount(id int) (float64, error) {
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

	var amount float64
	err = tx.QueryRow(ctx, sql, args...).Scan(&amount)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, nil // просто нет строки — значит, баланс = 0
		}
		return 0, err
	}

	if err = tx.Commit(ctx); err != nil {
		return 0, err
	}

	return amount, nil
}

func (repo *BalanceRepository) AddPoints(userID int, points float64) error {
	ctx := context.Background()
	tx, err := repo.DB.Pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// Обновление или вставка
	insertQuery := repo.DB.Builder.
		Insert("balances").
		Columns("user_id", "amount").
		Values(userID, points).
		Suffix("ON CONFLICT (user_id) DO UPDATE SET amount = balances.amount + EXCLUDED.amount")
	insertSQL, insertArgs, err := insertQuery.ToSql()
	if err != nil {
		return err
	}

	if _, err = tx.Exec(ctx, insertSQL, insertArgs...); err != nil {
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
