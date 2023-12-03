package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SendJson(w http.ResponseWriter, status int, response map[string]interface{}) {
	responseStr, err := json.Marshal(response)
	if err != nil {
		panic("Unable to write response")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintf(w, "%v", string(responseStr))
}

type Json map[string]interface{}
