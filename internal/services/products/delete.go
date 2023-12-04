package products

import (
	"github.com/brunoeduardodev/go-raw-sql/internal/helpers"
	"github.com/brunoeduardodev/go-raw-sql/repositories"
)

type DeleteProductResponse struct {
	Deleted bool `json:"deleted"`
}

func DeleteProduct(repository repositories.ProductRepository, id int) (*DeleteProductResponse, *helpers.RequestError) {
	err := repository.Delete(id)

	if err != nil {
		return nil, &helpers.RequestError{
			Message: "Could not delete product",
			Status:  400,
			Err:     err,
		}
	}

	return &DeleteProductResponse{Deleted: true}, nil
}
