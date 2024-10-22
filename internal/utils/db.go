package utils

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func ConnectDB() (*sql.DB, error) {
	connStr := "postgres://carmaint:example@postgres/carmaint"
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}

	db.Query("SELECT 1")
	return db, err
}
