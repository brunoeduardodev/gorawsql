package products

import (
	"github.com/brunoeduardodev/go-raw-sql/internal/helpers"
	"github.com/brunoeduardodev/go-raw-sql/repositories"
)

type CreateProductResponse struct {
	Product repositories.Product `json:"product"`
}

func CreateProduct(repository repositories.ProductRepository, input repositories.CreateProductInput) (*CreateProductResponse, *helpers.RequestError) {
	product, err := repository.Create(input)

	if err != nil {
		return nil, &helpers.RequestError{
			Message: "Could not create product",
			Status:  400,
			Err:     err,
		}
	}

	return &CreateProductResponse{Product: *product}, nil
}
