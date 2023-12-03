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

	product, err := productRepository.Update(1, repositories.UpdateProductInput{
		Name:  "New Name",
		Price: 3333,
	})

	if err != nil {
		fmt.Printf("Error while updating product %v\n", err)
	} else {
		fmt.Println(product.ToString())
	}

	products, err := productRepository.List()

	if err != nil {
		fmt.Printf("Error while listing product %v\n", err)
	} else {
		for _, product := range *products {
			fmt.Println(product.ToString())
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Reached here!\n")
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
