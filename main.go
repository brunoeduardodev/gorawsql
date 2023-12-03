package main

import (
	"log"

	"github.com/brunoeduardodev/go-raw-sql/internal/infra"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	infra.StartServer(8090)
}
