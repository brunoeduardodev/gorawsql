package main_test

import (
	"testing"

	"github.com/brunoeduardodev/go-raw-sql/internal/helpers"
)

func TestHealthCheck(t *testing.T) {
	response := helpers.TestGetRequest(t, "health", 200)

	if response["message"] != "Service is up and running" {
		t.Errorf("expected message to be \"Service is up and running\", but got %v", response["message"])
	}
}
