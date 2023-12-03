package main_test

import (
	"encoding/json"
	"testing"

	"github.com/brunoeduardodev/go-raw-sql/internal/helpers"
	"github.com/brunoeduardodev/go-raw-sql/internal/services/products"
)

func TestProducts(t *testing.T) {
	tests := []helpers.RouteTest{
		{
			Name:               "Should return empty list without products",
			Route:              "products",
			Method:             "GET",
			ExpectedStatusCode: 200,
			TestResponse: func(decoder *json.Decoder) {
				var response products.ListProductsResponse
				helpers.EnsureResponseDecodesTo(t, decoder, &response)

				if len(response.Products) > 0 {
					t.Errorf("expected empty list of products, but got %v", response.Products)
				}
			},
		},
	}

	helpers.RunRoutesTests(t, tests)

}
