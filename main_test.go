package main_test

import (
	"context"
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
	}
}

func clearTable() {
	app.DB.Exec(context.Background(), "DELETE FROM products")
	app.DB.Exec(context.Background(), "ALTER SEQUENCE products_id_seq RESTART WITH 1")
}

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	app.Setup(infra.AppConfig{
		DatabaseUrl: os.Getenv("DATABASE_URL_TEST"),
	})

	ensureTableExists()

	go app.Run(8090)
	m.Run()

	clearTable()
	os.Interrupt.Signal()
}
