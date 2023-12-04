package products

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/brunoeduardodev/go-raw-sql/internal/helpers"
	"github.com/brunoeduardodev/go-raw-sql/repositories"
)

func getIdFromUrl(url string) (int, error) {
	segments := strings.Split(url, "/")[2:]
	if len(segments) != 1 {
		return 0, fmt.Errorf("invalid segments amount")
	}

	id, err := strconv.Atoi(segments[0])
	if err != nil {
		return 0, err
	}

	return id, nil
}

func ProductHandler(repository repositories.ProductRepository) helpers.RequestHandler {
	return func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			segments := strings.Split(req.URL.Path, "/")[2:]
			if len(segments) > 1 {
				http.NotFound(w, req)
				return
			}

			if len(segments) == 0 {
				response, err := ListProducts(repository)
				if err != nil {
					helpers.SendError(w, *err)
					return
				}
				helpers.SendJson(w, 200, response)
				return
			}

			id, convError := strconv.Atoi(segments[0])
			if convError != nil {
				helpers.SendNotFound(w)
				return
			}

			response, err := FindProductById(repository, id)
			if err != nil {
				helpers.SendError(w, *err)
				return
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
				return
			}

			response, responseError := CreateProduct(repository, input)
			if responseError != nil {
				helpers.SendError(w, *responseError)
				return
			}

			helpers.SendJson(w, 200, response)
			return
		case "PUT":
			id, err := getIdFromUrl(req.URL.Path)
			if err != nil {
				helpers.SendNotFound(w)
				return
			}

			var input repositories.UpdateProductInput
			err = json.NewDecoder(req.Body).Decode(&input)

			if err != nil {
				helpers.SendError(w, helpers.RequestError{
					Message: "Invalid request body",
					Status:  400,
					Err:     err,
				})
				return
			}

			response, requestError := UpdateProduct(repository, id, input)
			if requestError != nil {
				helpers.SendError(w, *requestError)
				return
			}

			helpers.SendJson(w, 200, response)
			return

		case "DELETE":
			id, err := getIdFromUrl(req.URL.Path)
			if err != nil {
				helpers.SendNotFound(w)
				return
			}

			response, requestError := DeleteProduct(repository, id)
			if requestError != nil {
				helpers.SendError(w, *requestError)
				return
			}

			helpers.SendJson(w, 200, response)
			return

		default:
			helpers.SendNotFound(w)
		}

	}
}
