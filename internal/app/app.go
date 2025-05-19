package app

import (
	"database/sql"
	"github.com/SiyovushAbdulloev/gophermart/internal/handler/http"
	AuthHandler "github.com/SiyovushAbdulloev/gophermart/internal/handler/http/auth"
	OrderHandler "github.com/SiyovushAbdulloev/gophermart/internal/handler/http/order"
	AuthRepo "github.com/SiyovushAbdulloev/gophermart/internal/repository/postgres/auth"
	OrderRepo "github.com/SiyovushAbdulloev/gophermart/internal/repository/postgres/order"
	AuthUsecase "github.com/SiyovushAbdulloev/gophermart/internal/usecase/auth"
	OrderUsecase "github.com/SiyovushAbdulloev/gophermart/internal/usecase/order"
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

	log.Println(postgresDB.Pool)
	log.Println(postgresDB.Builder)
	log.Println("Migrations started")
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
	log.Println("Migrations finished")

	authRepo := AuthRepo.New(postgresDB)
	authUC := AuthUsecase.New(authRepo, cfg.JWTSecretKey, time.Duration(cfg.JWTExpire)*time.Hour)
	authHl := AuthHandler.New(authUC)

	orderRepo := OrderRepo.New(postgresDB)
	orderUC := OrderUsecase.New(orderRepo)
	orderHl := OrderHandler.New(orderUC)

	httpServer := httpserver.New(httpserver.WithAddress(cfg.ServerAddr))
	http.DefineAuthRoutes(httpServer.App, authHl)
	http.DefineOrderRoutes(httpServer.App, orderHl, cfg.JWTSecretKey, authRepo)

	go func() {
		err := httpServer.Run()

		if err != nil {
			panic(err)
		}
	}()

	select {}
}
