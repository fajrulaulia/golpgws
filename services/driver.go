package golpgws

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

func Driver() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}
