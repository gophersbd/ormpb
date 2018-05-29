package helper

import (
	"database/sql"
)

type Framework struct {
	PostgresClient *sql.DB
	MySQLClient    *sql.DB
}

func New(pg, ms *sql.DB) *Framework {
	return &Framework{
		PostgresClient: pg,
		MySQLClient:    ms,
	}
}
