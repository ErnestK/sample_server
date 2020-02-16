package config

import (
	"log"
	"os"

	"database/sql"
	// need only as driver for database
	_ "github.com/mattn/go-sqlite3"
)

// InitDB we use to create one instance of database connection
func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", os.Getenv("PATH_TO_DUMP"))
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
