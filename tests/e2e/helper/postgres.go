package helper

import (
	"database/sql"
	"fmt"
	"os"
)

const (
	PostgresURL = "postgres_url"
)

func GetPostgresClient() (*sql.DB, error) {
	postgresURL, found := os.LookupEnv(PostgresURL)
	if !found {
		return nil, nil
	}
	connStr := fmt.Sprintf("postgres://postgres@%s/?sslmode=disable", postgresURL)
	return sql.Open("postgres", connStr)
}
