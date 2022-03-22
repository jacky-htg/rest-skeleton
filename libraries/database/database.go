package database

import (
	"database/sql"
	"os"
)

func OpenDB() (*sql.DB, error) {
	return sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_SOURCE"))
}
