package main_test

import (
	"encoding/json"
	"testing"

	"github.com/brunoeduardodev/go-raw-sql/internal/helpers"
	"github.com/brunoeduardodev/go-raw-sql/internal/services/products"
	"github.com/brunoeduardodev/go-raw-sql/repositories"
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
		{
			Name:               "Should be able to create a product",
			Method:             "POST",
			Route:              "products",
			ExpectedStatusCode: 200,
			Body: repositories.CreateProductInput{
				Name:  "Product 1",
				Price: 10000,
			},
			TestResponse: func(decoder *json.Decoder) {
				var response products.CreateProductResponse
				helpers.EnsureResponseDecodesTo(t, decoder, &response)

				if response.Product.Name != "Product 1" && response.Product.Price != 10000 {
					t.Errorf("Expected receive product with name \"Product 1\" and price 10000 but got Name=%s, Price=%d", response.Product.Name, response.Product.Price)
				}
			},
		},
		{
			Name:               "Should be able to list created product",
			Method:             "GET",
			Route:              "products",
			ExpectedStatusCode: 200,
			TestResponse: func(decoder *json.Decoder) {
				var response products.ListProductsResponse
				helpers.EnsureResponseDecodesTo(t, decoder, &response)

				if len(response.Products) != 1 {
					t.Errorf("Expected to receive only 1 product, but got %d", len(response.Products))
				}

				if response.Products[0].Name != "Product 1" {
					t.Errorf("Expected to receive a product with name \"Product 1\", but got %s", response.Products[0].Name)
				}

			},
		},
		{
			Name:               "Should be able to find created product by id",
			Method:             "GET",
			Route:              "products/1",
			ExpectedStatusCode: 200,
			TestResponse: func(decoder *json.Decoder) {
				var response products.FindProductByIdResponse
				helpers.EnsureResponseDecodesTo(t, decoder, &response)

				if response.Product.Id != 1 {
					t.Errorf("Expected to find product with id 1 but got %d", response.Product.Id)
				}

				if response.Product.Name != "Product 1" {
					t.Errorf("Expected to receive a product with name \"Product 1\", but got %s", response.Product.Name)
				}
			},
		},
		{
			Name:               "Should be able to update product",
			Method:             "PUT",
			Route:              "products/1",
			ExpectedStatusCode: 200,
			Body: repositories.UpdateProductInput{
				Name:  "Updated 1",
				Price: 10000,
			},
			TestResponse: func(decoder *json.Decoder) {
				var response products.UpdateProductResponse
				helpers.EnsureResponseDecodesTo(t, decoder, &response)

				if response.Product.Id != 1 {
					t.Errorf("Expected to find product with id 1 but got %d", response.Product.Id)
				}

				if response.Product.Name != "Updated 1" {
					t.Errorf("Expected to receive a product with name \"Updated 1\", but got %s", response.Product.Name)
				}
			},
		},
		{
			Name:               "Should be able to retrieve product with updated name",
			Method:             "GET",
			Route:              "products/1",
			ExpectedStatusCode: 200,
			TestResponse: func(decoder *json.Decoder) {
				var response products.FindProductByIdResponse
				helpers.EnsureResponseDecodesTo(t, decoder, &response)

				if response.Product.Id != 1 {
					t.Errorf("Expected to find product with id 1 but got %d", response.Product.Id)
				}

				if response.Product.Name != "Updated 1" {
					t.Errorf("Expected to receive a product with name \"Updated 1\", but got %s", response.Product.Name)
				}
			},
		},
	}

	helpers.RunRoutesTests(t, tests)

}
