package get_summary

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"

    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB 

func Handler(w http.ResponseWriter, r *http.Request) {
    v := r.URL.Query()
    fmt.Println(v.Get("name"))

    db, err := sql.Open("sqlite3", "/Users/ernestkhasanzhinov/work/go/sample_server/db/positions.db")
    if err != nil {
        log.Fatalln(err)
    }

    summary := CountByDomain{}
    err = db.QueryRow("select domain, count(*) count from positions group by domain having domain = 'apostrophied.co.uk'").Scan(&summary.Domain, &summary.Count)

    fmt.Println(summary)

    bytes, err := json.Marshal(summary)
    if err != nil {
        fmt.Println("Can't serialize", summary)
    }

    w.Write(bytes)
}