package app

import (
	"database/sql"
	"github.com/SiyovushAbdulloev/gophermart/internal/handler/http"
	AuthHandler "github.com/SiyovushAbdulloev/gophermart/internal/handler/http/auth"
	OrderHandler "github.com/SiyovushAbdulloev/gophermart/internal/handler/http/order"
	WithdrawHandler "github.com/SiyovushAbdulloev/gophermart/internal/handler/http/withdraw"
	AuthRepo "github.com/SiyovushAbdulloev/gophermart/internal/repository/postgres/auth"
	BalanceRepo "github.com/SiyovushAbdulloev/gophermart/internal/repository/postgres/balance"
	OrderRepo "github.com/SiyovushAbdulloev/gophermart/internal/repository/postgres/order"
	WithdrawRepo "github.com/SiyovushAbdulloev/gophermart/internal/repository/postgres/withdraw"
	AuthUsecase "github.com/SiyovushAbdulloev/gophermart/internal/usecase/auth"
	BalanceUsecase "github.com/SiyovushAbdulloev/gophermart/internal/usecase/balance"
	OrderUsecase "github.com/SiyovushAbdulloev/gophermart/internal/usecase/order"
	WithdrawUsecase "github.com/SiyovushAbdulloev/gophermart/internal/usecase/withdraw"
	"github.com/SiyovushAbdulloev/gophermart/pkg/config"
	"github.com/SiyovushAbdulloev/gophermart/pkg/httpserver"
	"github.com/SiyovushAbdulloev/gophermart/pkg/postgres"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"log"
	"time"
)

func Main(cfg *config.Config) {
	postgresDB, err := postgres.New(cfg.DatabaseURI, postgres.ConnAttempts(cfg.DatabaseMaxConn), postgres.ConnDelay(cfg.DatabaseMaxDelay), postgres.MaxPoolSize(cfg.DatabaseMaxConn))
	if err != nil {
		panic(err)
	}

	dbMigration, err := sql.Open("postgres", cfg.DatabaseURI)
	if err != nil {
		log.Fatalf("❌ goose.Open error: %v", err)
	}
	defer dbMigration.Close()

	if err = goose.SetDialect("postgres"); err != nil {
		log.Fatalf("❌ goose.Select error: %v", err)
	}

	if err = goose.Up(dbMigration, "./migrations"); err != nil {
		log.Fatalf("❌ goose.Up error: %v", err)
	}

	authRepo := AuthRepo.New(postgresDB)
	authUC := AuthUsecase.New(authRepo, cfg.JWTSecretKey, time.Duration(cfg.JWTExpire)*time.Hour)
	authHl := AuthHandler.New(authUC)

	orderRepo := OrderRepo.New(postgresDB)
	orderUC := OrderUsecase.New(orderRepo)
	orderHl := OrderHandler.New(orderUC)

	balanceRepo := BalanceRepo.New(postgresDB)
	balanceUC := BalanceUsecase.New(balanceRepo)

	withdrawRepo := WithdrawRepo.New(postgresDB)
	withdrawUC := WithdrawUsecase.New(withdrawRepo)
	withdrawHl := WithdrawHandler.New(withdrawUC, balanceUC)

	httpServer := httpserver.New(httpserver.WithAddress(cfg.ServerAddr))
	http.DefineAuthRoutes(httpServer.App, authHl)
	http.DefineOrderRoutes(httpServer.App, orderHl, cfg.JWTSecretKey, authRepo)
	http.DefineWithdrawRoutes(httpServer.App, withdrawHl, cfg.JWTSecretKey, authRepo)

	go func() {
		err := httpServer.Run()

		if err != nil {
			panic(err)
		}
	}()

	select {}
}
