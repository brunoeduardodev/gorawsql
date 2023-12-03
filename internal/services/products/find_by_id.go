package products

import (
	"github.com/brunoeduardodev/go-raw-sql/internal/helpers"
	"github.com/brunoeduardodev/go-raw-sql/repositories"
)

type FindProductByIdResponse struct {
	Product repositories.Product `json:"product"`
}

func FindProductById(repository repositories.ProductRepository, id int) (*FindProductByIdResponse, *helpers.RequestError) {
	product, err := repository.FindById(id)

	if err != nil {
		return nil, &helpers.RequestError{
			Message: "Product not found",
			Status:  400,
			Err:     err,
		}
	}

	return &FindProductByIdResponse{Product: *product}, nil
}
