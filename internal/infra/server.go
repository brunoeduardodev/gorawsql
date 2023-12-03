package infra

import (
	"fmt"
	"net/http"

	"github.com/brunoeduardodev/go-raw-sql/internal/infra/database"
	"github.com/brunoeduardodev/go-raw-sql/internal/services/health"
	"github.com/brunoeduardodev/go-raw-sql/repositories"
)

func StartServer(port int) {
	http.HandleFunc("/health", health.HealthHandler)

	conn := database.GetConnection()
	defer database.CloseConnection(conn)

	productRepository := repositories.MakePgProductRepository(conn)
	products, err := productRepository.List()

	if err != nil {
		fmt.Printf("Error while creating product %v\n", err)
	} else {
		for _, product := range *products {
			fmt.Printf("Product id %d name %s price %d\n", product.Id, product.Name, product.Price)
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Reached here!\n")
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
