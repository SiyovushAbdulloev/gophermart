package config

import (
	"flag"
	"github.com/caarlos0/env/v11"
)

type Server struct {
	Addr string
}

type DB struct {
	Username    string
	Password    string
	URI         string
	MaxConn     string
	MaxAttempts string
	MaxDelay    string
}

type App struct {
	AccrualAddr string
}

type Config struct {
	ServerAddr          string `env:"RUN_ADDRESS"`
	AccrualAddr         string `env:"ACCRUAL_SYSTEM_ADDRESS"`
	DatabaseURI         string `env:"DATABASE_URI"`
	DatabaseMaxConn     int    `env:"DB_MAX_CONN"`
	DatabaseMaxAttempts int    `env:"DB_MAX_ATTEMPTS"`
	DatabaseMaxDelay    int    `env:"DB_MAX_DELAY"`
	JWTSecretKey        string `env:"JWT_SECRET_KEY"`
	JWTExpire           int    `env:"JWT_EXPIRE"`
}

func New() (*Config, error) {
	cfg := &Config{
		DatabaseMaxAttempts: 3,
		DatabaseMaxDelay:    7,
		DatabaseMaxConn:     10,
		JWTSecretKey:        "secret",
		JWTExpire:           10,
	}

	// Подгрузка переменных окружения (если флаги не заданы)
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	// Значения по умолчанию, если флаг не передан
	flag.StringVar(&cfg.ServerAddr, "a", cfg.ServerAddr, "address and port to run the service (RUN_ADDRESS)")
	flag.StringVar(&cfg.DatabaseURI, "d", cfg.DatabaseURI, "database connection string (DATABASE_URI)")
	flag.StringVar(&cfg.AccrualAddr, "r", cfg.AccrualAddr, "accrual system address (ACCRUAL_SYSTEM_ADDRESS)")
	flag.Parse()

	return cfg, nil
}
