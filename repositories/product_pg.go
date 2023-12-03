package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Create(conn *pgx.Conn) {
	conn.QueryRow(context.Background(), "INSERT INTO")
}
