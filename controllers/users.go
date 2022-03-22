package controllers

import (
	"database/sql"
	"log"
)

type Users struct {
	Db  *sql.DB
	Log *log.Logger
}
