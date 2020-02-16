package initialize_config

import (
	"log"
    "os"

    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
    db, err := sql.Open("sqlite3", os.Getenv("PATH_TO_DUMP"))
    if err != nil {
        log.Fatalln(err)
    }

    return db
}
