package initialize_config

import (
	"log"
	
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
    db, err := sql.Open("sqlite3", "/Users/ernestkhasanzhinov/work/go/sample_server/db/positions.db")
    if err != nil {
        log.Fatalln(err)
    }

    return db
}
