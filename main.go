package main

import (
	"log"
	"os"

	"github.com/brunoeduardodev/go-raw-sql/internal/infra"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	app := infra.App{}
	app.Setup(infra.AppConfig{
		DatabaseUrl: os.Getenv("DATABASE_URL"),
	})

	app.Run(8090)
}
