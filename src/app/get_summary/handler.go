package get_summary

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"

    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB 

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    vars := mux.Vars(r) 

    summary := CountByDomain{}
    sql_stmt := "select domain, count(*) count from positions group by domain having domain = '%s'"

    err := DB.QueryRow(fmt.Sprintf(sql_stmt, vars["domain_name"])).Scan(&summary.Domain, &summary.Count)
    if err != nil {
        log.Println("Error duraing access to db!")
    }

    bytes, err := json.Marshal(summary)
    if err != nil {
        log.Println("Can't serialize", summary)
    }

    w.Write(bytes)
}