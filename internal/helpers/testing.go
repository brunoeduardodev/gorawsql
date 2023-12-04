package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

type RouteTest struct {
	Name               string
	Method             string
	Route              string
	Body               any
	ExpectedStatusCode int
	TestResponse       func(decoder *json.Decoder)
}

func RunRoutesTests(t *testing.T, tests []RouteTest) {
	for _, test := range tests {
		bodyDecoder := TestRequest(t, test.Route, test.Method, test.Body, test.ExpectedStatusCode)
		test.TestResponse(bodyDecoder)
	}
}

func TestRequest(t *testing.T, route string, method string, body interface{}, expectedStatusCode int) *json.Decoder {
	var requestBody io.Reader

	if body != nil {
		stringified, err := json.Marshal(body)
		if err != nil {
			panic(fmt.Sprintf("Unable to stringify body %v", err))
		}
		requestBody = strings.NewReader(string(stringified))
	} else {
		requestBody = http.NoBody

	}

	request, err := http.NewRequest(method, fmt.Sprintf("http://localhost:8090/%s", route), requestBody)
	if body != nil {
		request.Header.Set("Content-Type", "application-json")
	}

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

	return json.NewDecoder(response.Body)
}

func EnsureResponseDecodesTo(t *testing.T, decoder *json.Decoder, response any) {
	err := decoder.Decode(&response)
	if err != nil {
		t.Errorf("Expected to decode, but got %v", err)
	}
}
