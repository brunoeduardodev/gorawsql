package products

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/brunoeduardodev/go-raw-sql/internal/helpers"
	"github.com/brunoeduardodev/go-raw-sql/repositories"
)

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
		case "POST":
			input := repositories.CreateProductInput{}
			err := json.NewDecoder(req.Body).Decode(&input)

			if err != nil {
				helpers.SendError(w, helpers.RequestError{
					Message: "Invalid request body",
					Status:  400,
					Err:     err,
				})
			}

			response, responseError := CreateProduct(repository, input)
			if responseError != nil {
				helpers.SendError(w, *responseError)
				return
			}

			helpers.SendJson(w, 200, response)
			return

		default:
			helpers.SendNotFound(w)
		}

	}
}
