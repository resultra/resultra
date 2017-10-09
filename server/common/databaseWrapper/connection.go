package databaseWrapper

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var dbHandle *sql.DB

func init() {

	var err error
	dbHandle, err = sql.Open("sqlite3", "./tracker.db")
	if err != nil {
		log.Fatal(err)
	}

	// Configure the maximum number of open connections.
	dbHandle.SetMaxOpenConns(75)

	// Open doesn't directly open the database connection. To verify the connection, the Ping() function
	// is needed.
	if err := dbHandle.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Database connection established")
}

func DBHandle() *sql.DB {
	return dbHandle
}
