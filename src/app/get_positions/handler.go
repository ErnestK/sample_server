package get_positions

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
    vars := mux.Vars(r)
    fmt.Println(vars["domain_name"])

    positions := []Position{}
    sql_stmt := "select keyword, position, url, volume, results, updated from positions where domain = 'apostrophied.co.uk' order by volume asc limit 10 offset 10"
    rows, err := DB.Query(sql_stmt)
    for rows.Next() {
        position := Position{}
        err := rows.Scan(&position.Keyword, &position.Position, &position.Url, &position.Volume, &position.Results, &position.Updated)
        if err != nil {
            log.Fatalln(err)
        }
        positions = append(positions, position)
    }

    dm := DomainWithPosition{Domain: "apostrophied.co.uk", Positions: positions}
    bytes, err := json.Marshal(dm)
    if err != nil {
        fmt.Println("Can't serialize", dm)
    }

    w.Write(bytes)
}