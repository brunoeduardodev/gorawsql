package main_test

import (
	"encoding/json"
	"testing"

	"github.com/brunoeduardodev/go-raw-sql/internal/helpers"
	"github.com/brunoeduardodev/go-raw-sql/internal/services/health"
)

func TestHealthCheck(t *testing.T) {
	tests := []helpers.RouteTest{
		{
			Name:               "Should return Service is up and running",
			Method:             "GET",
			Route:              "health",
			ExpectedStatusCode: 200,
			TestResponse: func(decoder *json.Decoder) {
				var response health.HealthCheckResponse
				helpers.EnsureResponseDecodesTo(t, decoder, &response)

				if response.Message != "Service is up and running" {
					t.Errorf("expected message to be \"Service is up and running\", but got %v", response.Message)
				}

			},
		},
	}

	helpers.RunRoutesTests(t, tests)

}
