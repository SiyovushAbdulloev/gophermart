package main

import (
	"github.com/SiyovushAbdulloev/gophermart/internal/app"
	"github.com/SiyovushAbdulloev/gophermart/pkg/config"
)

func main() {
	//fmt.Println("FIRST START")
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file", err)
	//}

	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	app.Main(cfg)
}
