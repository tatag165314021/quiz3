package database

import (
	"database/sql"
)

var (
	DbConnection *sql.DB
)

func ConnectionDatabase(dbParam *sql.DB) {
	DbConnection = dbParam
}
