package products

import (
	"net/http"

	"github.com/brunoeduardodev/go-raw-sql/internal/helpers"
	"github.com/brunoeduardodev/go-raw-sql/repositories"
)

func ListProducts(repository repositories.ProductRepository) helpers.RequestHandler {
	return func(w http.ResponseWriter, req *http.Request) {
		products, err := repository.List()

		if err != nil {
			helpers.SendError(w, http.StatusInternalServerError, "Could not list products.", err)
			return
		}

		if err != nil {
			helpers.SendError(w, http.StatusInternalServerError, "Could not return products.", err)
			return
		}

		helpers.SendJson(w, http.StatusOK, helpers.Json{
			"products": products,
		})

	}
}
