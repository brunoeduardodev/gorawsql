package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SendJson(w http.ResponseWriter, status int, response any) {
	responseStr, err := json.Marshal(response)
	if err != nil {
		panic("Unable to write response")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintf(w, "%v", string(responseStr))
}

type Json map[string]interface{}

type RequestError struct {
	Message string
	Status  int
	Err     error
}

func SendError(w http.ResponseWriter, err RequestError) {
	fmt.Printf("Error: %v\n", err.Err)
	SendJson(w, err.Status, Json{"error": err.Message})
}

func SendNotFound(w http.ResponseWriter) {
	w.WriteHeader(404)
}

type RequestHandler func(w http.ResponseWriter, req *http.Request)

type ErrorResponse struct {
	Error string `json:"error"`
}
