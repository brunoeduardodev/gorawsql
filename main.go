package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/brunoeduardodev/go-raw-sql/internal/infra/database"
	"github.com/brunoeduardodev/go-raw-sql/internal/services/health"
	"github.com/brunoeduardodev/go-raw-sql/internal/services/products"
	"github.com/brunoeduardodev/go-raw-sql/repositories"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type App struct {
	conn *pgx.Conn
}

type AppConfig struct {
	databaseUrl string
}

func (A *App) Setup(config AppConfig) {
	A.conn = database.GetConnection(config.databaseUrl)

	productRepository := repositories.MakePgProductRepository(A.conn)

	http.HandleFunc("/products", products.ListProducts(productRepository))
	http.HandleFunc("/health", health.HealthHandler)
}

func (A *App) Run(port int) {
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	database.CloseConnection(A.conn)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	app := App{}
	app.Setup(AppConfig{
		databaseUrl: os.Getenv("DATABASE_URL"),
	})

	app.Run(8090)
}
