package models

import (
	"database/sql"
	"log"
)

// User : struct of User
type User struct {
	ID       uint
	Username string
	Password string
	Email    string
	IsActive bool
	Db       *sql.DB
	Log      *log.Logger
}
