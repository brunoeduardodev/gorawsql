package main_test

import (
	"testing"

	"github.com/brunoeduardodev/go-raw-sql/internal/helpers"
)

func TestHealthCheck(t *testing.T) {
	tests := []helpers.RouteTest{
		{
			Name:               "Should return Service is up and running",
			Method:             "GET",
			Route:              "health",
			ExpectedStatusCode: 200,
			TestResponse: func(response helpers.Json) {
				if response["message"] != "Service is up and running" {
					t.Errorf("expected message to be \"Service is up and running\", but got %v", response["message"])
				}

			},
		},
	}

	helpers.RunRoutesTests(t, tests)

}
