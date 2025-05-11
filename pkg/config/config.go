package config

import "github.com/caarlos0/env/v11"

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
	ServerAddr          string `env:"SERVER_ADDR,required"`
	AccrualAddr         string `env:"ACCRUAL_ADDR,required"`
	DatabaseUsername    string `env:"DB_USERNAME,required"`
	DatabasePassword    string `env:"DB_PASSWORD,required"`
	DatabaseURI         string `env:"DB_URI,required"`
	DatabaseMaxConn     string `env:"DB_MAX_CONN,required"`
	DatabaseMaxAttempts string `env:"DB_MAX_ATTEMPTS,required"`
	DatabaseMaxDelay    string `env:"DB_MAX_DELAY,required"`
}

func New() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
