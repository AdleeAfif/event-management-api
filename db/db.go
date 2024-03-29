package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	DB, err := sql.Open("sqlite3", "eventmanager.db")

	if err != nil {
		panic("Unable to access database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetConnMaxIdleTime(5)
}
