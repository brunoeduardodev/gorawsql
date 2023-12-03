package infra

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/brunoeduardodev/go-raw-sql/internal/infra/database"
	"github.com/brunoeduardodev/go-raw-sql/internal/services/health"
)

func StartServer(port int) {
	http.HandleFunc("/health", health.HealthHandler)

	conn := database.GetConnection()
	defer database.CloseConnection(conn)

	var id int
	var name string
	var price int

	err := conn.QueryRow(context.Background(), "SELECT id, name, price FROM PRODUCTS").Scan(&id, &name, &price)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("ID=%d Name=%s Price=%d\n", id, name, price)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Reached here!\n")
	})

	fmt.Println("Started server!")

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	fmt.Println("Stopped server!")
}
