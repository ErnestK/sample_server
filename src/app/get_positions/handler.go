package get_positions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const LIMIT = 50
const DEFAULT_SORT_BY = "volume"
const DEFAULT_PAGE = "0"

var DB *sql.DB

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	domain_name := vars["domain_name"]

	query := r.URL.Query()
	sort_by := getOrDefault(query.Get("sort_by"), DEFAULT_SORT_BY)

	page_str := getOrDefault(query.Get("page"), DEFAULT_PAGE)
	page, err := strconv.Atoi(page_str)
	if err != nil {
		log.Fatal("Error read page params!")
	}

	positions := []Position{}

	sql_stmt := "select keyword, position, url, volume, results, updated from positions where domain = '%s' order by %s asc limit %d offset %d"
	rows, err := DB.Query(fmt.Sprintf(sql_stmt, domain_name, sort_by, LIMIT, page))
	for rows.Next() {
		position := Position{}
		err := rows.Scan(&position.Keyword, &position.Position, &position.Url, &position.Volume, &position.Results, &position.Updated)
		if err != nil {
			log.Fatalln(err)
		}
		positions = append(positions, position)
	}

	dm := DomainWithPosition{Domain: domain_name, Positions: positions}
	bytes, err := json.Marshal(dm)
	if err != nil {
		fmt.Println("Can't serialize", dm)
	}

	w.Write(bytes)
}

func getOrDefault(val, fallback string) string {
	if val != "" {
		return val
	}
	return fallback
}
