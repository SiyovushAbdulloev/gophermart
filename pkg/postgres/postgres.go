package postgres

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

type Postgres struct {
	maxPoolSize  int
	connAttempts int
	connDelay    int

	Builder squirrel.StatementBuilderType
	Pool    *pgxpool.Pool
}

func New(url string, opts ...Option) (*Postgres, error) {
	pg := &Postgres{
		maxPoolSize:  10,
		connAttempts: 0,
		connDelay:    0,
	}

	for _, opt := range opts {
		opt(pg)
	}

	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	var err error

	for i := 0; i < pg.connAttempts; i++ {
		pg.Pool, err = pgxpool.New(context.Background(), url)
		if err == nil {
			break
		}

		log.Printf("Postgres is trying to connect, attempts left: %d", pg.connAttempts-(i+1))

		time.Sleep(time.Second * time.Duration(pg.connDelay))
	}

	if err != nil {
		_ = fmt.Errorf("postgres - NewPostgres - connAttempts == 0: %w", err)
	}

	return pg, nil
}

func (pg *Postgres) Close() {
	if pg.Pool != nil {
		pg.Pool.Close()
	}
}
