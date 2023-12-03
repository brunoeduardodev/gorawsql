package health

import (
	"net/http"

	"github.com/brunoeduardodev/go-raw-sql/internal/helpers"
)

type HealthCheckResponse struct {
	Message string `json:"message"`
}

func HealthHandler(w http.ResponseWriter, req *http.Request) {
	response := HealthCheckResponse{
		Message: "Service is up and running",
	}
	helpers.SendJson(w, http.StatusOK, response)
}
