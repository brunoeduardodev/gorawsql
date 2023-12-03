package infra

import (
	"fmt"
	"net/http"

	"github.com/brunoeduardodev/go-raw-sql/internal/services/health"
)

func StartServer(port int) {
	http.HandleFunc("/health", health.HealthHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Reached here!\n")
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
