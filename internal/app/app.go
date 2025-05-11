package app

import (
	"github.com/SiyovushAbdulloev/gophermart/pkg/config"
	"github.com/SiyovushAbdulloev/gophermart/pkg/httpserver"
)

func Main(cfg *config.Config) {
	httpServer := httpserver.New(httpserver.WithAddress(cfg.ServerAddr))

	go func() {
		err := httpServer.Run()

		if err != nil {
			panic(err)
		}
	}()

	select {}
}
