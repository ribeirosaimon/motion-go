package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	connStr := "postgres://postgres:frajolinha202@localhost:5432/motion?sslmode=mode"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
