package databaseWrapper

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var dbHandle *sql.DB

func init() {

	var err error
	dbHandle, err = sql.Open("postgres", "user=devuser dbname=datasheet password=here4dev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

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
