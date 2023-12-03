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

func SendError(w http.ResponseWriter, status int, message string, e error) {
	fmt.Printf("Error: %v\n", e)
	SendJson(w, status, Json{"error": message})
}

type RequestHandler func(w http.ResponseWriter, req *http.Request)
