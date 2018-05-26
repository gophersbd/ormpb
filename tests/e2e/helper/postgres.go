package helper

import (
	"database/sql"
)

func GetPostgresClient() (*sql.DB, error) {
	connStr := "postgres://postgres@127.0.0.1/?sslmode=disable"
	return sql.Open("postgres", connStr)
}
