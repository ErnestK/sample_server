package positions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"database/sql"
)

const limit = 50
const defaultSortBy = "volume"
const defaultPage = "0"

// DB set when init db and set it to handler
var DB *sql.DB

// Handler public since in router we set handler to route, when init router
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	domainName := vars["domain_name"]

	query := r.URL.Query()
	sortBy := getOrDefault(query.Get("sort_by"), defaultSortBy)

	pageStr := getOrDefault(query.Get("page"), defaultPage)
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		log.Fatal("Error read page params!")
	}

	positions := []Position{}

	sqlStmt := "select keyword, position, url, volume, results, updated from positions where domain = '%s' order by %s asc limit %d offset %d"
	rows, err := DB.Query(fmt.Sprintf(sqlStmt, domainName, sortBy, limit, page))
	for rows.Next() {
		position := Position{}
		err := rows.Scan(&position.Keyword, &position.Position, &position.URL, &position.Volume, &position.Results, &position.Updated)
		if err != nil {
			log.Fatalln(err)
		}
		positions = append(positions, position)
	}

	dm := DomainWithPosition{Domain: domainName, Positions: positions}
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
