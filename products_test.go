package main_test

import (
	"testing"

	"github.com/brunoeduardodev/go-raw-sql/internal/helpers"
)

func TestProducts(t *testing.T) {
	tests := []helpers.RouteTest{
		{
			Name:               "Should return empty list without products",
			Route:              "products",
			Method:             "GET",
			ExpectedStatusCode: 200,
			TestResponse: func(body helpers.Json) {
				// if body["products"] != "[]" {
				// 	t.Errorf("expected empty list of products, but got %s", body["products"])
				// }
			},
		},
	}

	helpers.RunRoutesTests(t, tests)

}
