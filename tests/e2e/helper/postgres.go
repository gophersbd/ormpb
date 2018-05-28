package helper

import (
	"database/sql"
)

func GetPostgresClient() (*sql.DB, error) {
	connStr := "postgres://postgres@127.0.0.1:5432/?sslmode=disable"
	return sql.Open("postgres", connStr)
}
