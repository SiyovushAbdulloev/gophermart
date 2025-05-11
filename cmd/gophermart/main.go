package main

import (
	"github.com/SiyovushAbdulloev/gophermart/internal/app"
	"github.com/SiyovushAbdulloev/gophermart/pkg/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	app.Main(cfg)
}
