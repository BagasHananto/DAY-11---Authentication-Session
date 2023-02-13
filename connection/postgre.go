package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func DatabaseConnect() {
	var err error
	databaseUrl := "postgres://postgres:sx1475869@localhost:5432/Personal-Web"

	Conn, err = pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to database: %v", err)
		os.Exit(1)
	}
	fmt.Println("Successfully connected to database")
}
