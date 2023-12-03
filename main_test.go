package main_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/brunoeduardodev/go-raw-sql/internal/infra"
	"github.com/joho/godotenv"
)

var app infra.App

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  price INTEGER NOT NULL
)`

func ensureTableExists() {
	if _, err := app.DB.Query(context.Background(), tableCreationQuery); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Table exists!")
	}
}

func TestMain(m *testing.M) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	app.Setup(infra.AppConfig{
		DatabaseUrl: os.Getenv("DATABASE_URL"),
	})

	ensureTableExists()

	go app.Run(8090)
	m.Run()

	os.Interrupt.Signal()
}
