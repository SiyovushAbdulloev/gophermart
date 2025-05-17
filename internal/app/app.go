package app

import (
	"github.com/SiyovushAbdulloev/gophermart/pkg/config"
	"github.com/SiyovushAbdulloev/gophermart/pkg/httpserver"
	"github.com/SiyovushAbdulloev/gophermart/pkg/postgres"
	"log"
)

func Main(cfg *config.Config) {
	postgresDB, err := postgres.New(cfg.DatabaseURI)
	if err != nil {
		panic(err)
	}

	log.Println(postgresDB)
	httpServer := httpserver.New(httpserver.WithAddress(cfg.ServerAddr))

	go func() {
		err := httpServer.Run()

		if err != nil {
			panic(err)
		}
	}()

	select {}
}
