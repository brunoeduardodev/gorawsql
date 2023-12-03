package infra

import (
	"fmt"
	"net/http"

	"github.com/brunoeduardodev/go-raw-sql/internal/infra/database"
	"github.com/brunoeduardodev/go-raw-sql/internal/services/health"
	"github.com/brunoeduardodev/go-raw-sql/internal/services/products"
	"github.com/brunoeduardodev/go-raw-sql/repositories"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	DB *pgxpool.Pool
}

type AppConfig struct {
	DatabaseUrl string
}

func (A *App) Setup(config AppConfig) {
	A.DB = database.GetConnection(config.DatabaseUrl)

	productRepository := repositories.MakePgProductRepository(A.DB)

	http.HandleFunc("/products", products.ProductHandler(productRepository))
	http.HandleFunc("/products/", products.ProductHandler(productRepository))

	http.HandleFunc("/health", health.HealthHandler)
}

func (A *App) Run(port int) {
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	database.CloseConnection(A.DB)
}
