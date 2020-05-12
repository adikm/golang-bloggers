package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "bloggers"
	conn     *pgx.Conn
)

func Connect() {
	if conn != nil {
		_ = conn.Close(context.Background())
	}
	var err error
	conn, err = pgx.Connect(context.Background(), fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}
