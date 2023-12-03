package health_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/brunoeduardodev/go-raw-sql/internal/helpers"
)

func TestHealthCheck(t *testing.T) {
	response, err := http.Get("http://localhost:8090/health")

	if err != nil {
		t.Errorf("expected no errors, but got %v", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("expected 200 statuscode, but got %v", response.StatusCode)
	}

	responseBody := make(helpers.Json)
	json.NewDecoder(response.Body).Decode(&responseBody)
	response.Body.Close()

	if responseBody["message"] != "Service is up and running" {
		t.Errorf("expected message to be \"Service is up and running\", but got %v", responseBody["message"])
	}
}
