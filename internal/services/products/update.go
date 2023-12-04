package products

import (
	"github.com/brunoeduardodev/go-raw-sql/internal/helpers"
	"github.com/brunoeduardodev/go-raw-sql/repositories"
)

type UpdateProductResponse struct {
	Product repositories.Product `json:"product"`
}

func UpdateProduct(repository repositories.ProductRepository, id int, input repositories.UpdateProductInput) (*UpdateProductResponse, *helpers.RequestError) {
	product, err := repository.Update(id, input)

	if err != nil {
		return nil, &helpers.RequestError{
			Message: "Could not update product",
			Status:  400,
			Err:     err,
		}
	}

	return &UpdateProductResponse{
		Product: *product,
	}, nil
}
