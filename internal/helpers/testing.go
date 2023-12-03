package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestGetRequest(t *testing.T, route string, expectedStatusCode int) Json {
	response, err := http.Get(fmt.Sprintf("http://localhost:8090/%s", route))

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
