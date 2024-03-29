package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	fmt.Print("Connecting to database...\n")
	DB, err = sql.Open("sqlite3", "eventmanager.db")

	if err != nil {
		panic("Unable to access database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetConnMaxIdleTime(5)

	fmt.Print("Connection to database success!\n")

	createTables()
}

func createTables() {
	createEventsTable := `CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		userId INTEGER
	)`

	fmt.Print("Creating Event table...")
	_, err := DB.Exec(createEventsTable)

	if err != nil {
		fmt.Print("Failed!\n")
		panic("Unable to create Event table")
	}

	fmt.Print("Success!\n")
}
