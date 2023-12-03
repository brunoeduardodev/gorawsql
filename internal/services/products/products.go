package products

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/brunoeduardodev/go-raw-sql/internal/helpers"
	"github.com/brunoeduardodev/go-raw-sql/repositories"
)

type ListProductsResponse struct {
	Products []repositories.Product `json:"products"`
}

func ProductHandler(repository repositories.ProductRepository) helpers.RequestHandler {
	return func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			segments := strings.Split(req.URL.Path, "/")[2:]
			if len(segments) > 1 {
				http.NotFound(w, req)
			}

			if len(segments) == 0 {
				response, err := ListProducts(repository)
				if err != nil {
					helpers.SendError(w, *err)
				}
				helpers.SendJson(w, 200, response)
				return
			}

			id, convError := strconv.Atoi(segments[0])
			if convError != nil {
				helpers.SendNotFound(w)
			}

			response, err := FindProductById(repository, id)
			if err != nil {
				helpers.SendError(w, *err)
			}

			helpers.SendJson(w, 200, response)
			return
		default:
			helpers.SendNotFound(w)
		}

	}
}

func ListProducts(repository repositories.ProductRepository) (*ListProductsResponse, *helpers.RequestError) {
	products, err := repository.List()

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
