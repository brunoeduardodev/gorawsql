package main

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	app := App{}

	app.Setup(AppConfig{
		databaseUrl: os.Getenv("DATABASE_URL"),
	})

	go app.Run(8090)
	m.Run()

	os.Interrupt.Signal()
}
