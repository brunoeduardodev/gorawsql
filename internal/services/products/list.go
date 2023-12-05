package products

import (
	"github.com/brunoeduardodev/go-raw-sql/internal/helpers"
	"github.com/brunoeduardodev/go-raw-sql/repositories"
)

type ListProductsResponse struct {
	Products []repositories.Product `json:"products"`
}

func ListProducts(repository repositories.ProductRepository, input repositories.ListProductsInput) (*ListProductsResponse, *helpers.RequestError) {
	products, err := repository.List(input)

	if err != nil {
		return nil, &helpers.RequestError{
			Err:     err,
			Message: "Could not list products",
			Status:  503,
		}
	}

	response := ListProductsResponse{
		Products: *products,
	}

	return &response, nil
}
