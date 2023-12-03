package health

import (
	"net/http"

	"github.com/brunoeduardodev/go-raw-sql/internal/helpers"
)

func HealthHandler(w http.ResponseWriter, req *http.Request) {
	helpers.SendJson(w, http.StatusOK, helpers.Json{
		"message": "Service is up and running",
	})
}
