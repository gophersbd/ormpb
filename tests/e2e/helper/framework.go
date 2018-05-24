package helper

import (
	"database/sql"
)

type Framework struct {
	PostgresClient *sql.DB
}

func New(pg *sql.DB) *Framework {
	return &Framework{
		PostgresClient: pg,
	}
}
