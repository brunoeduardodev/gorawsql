package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetConnection(databaseUrl string) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	return dbpool
}

func CloseConnection(conn *pgxpool.Pool) {
	conn.Close()
}
