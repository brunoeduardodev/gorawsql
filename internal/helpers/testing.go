package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

type RouteTest struct {
	Name               string
	Method             string
	Route              string
	ExpectedStatusCode int
	TestResponse       func(response Json)
}

func RunRoutesTests(t *testing.T, tests []RouteTest) {
	for _, test := range tests {
		body := TestRequest(t, test.Route, test.Method, test.ExpectedStatusCode)
		test.TestResponse(body)
	}
}

func TestRequest(t *testing.T, route string, method string, expectedStatusCode int) Json {
	request, err := http.NewRequest(method, fmt.Sprintf("http://localhost:8090/%s", route), http.NoBody)
	if err != nil {
		panic(fmt.Sprintf("Unable to create request %v", err))
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Errorf("expected no errors, but got %v", err)
	}

	if response.StatusCode != expectedStatusCode {
		t.Errorf("expected %d statuscode, but got %d", expectedStatusCode, response.StatusCode)
	}

	responseBody := make(Json)
	json.NewDecoder(response.Body).Decode(&responseBody)
	response.Body.Close()

	return responseBody
}
