package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func BootstrapPG(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ClosePG(db *sql.DB) {
	fmt.Printf("pg database disconnected")

	db.Close()
}
