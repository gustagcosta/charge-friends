package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func BootstrapMysql(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	fmt.Printf("mysql database connected")

	return db, nil
}

func CloseMysql(db *sql.DB) {
	fmt.Printf("mysql database disconnected")

	db.Close()
}
