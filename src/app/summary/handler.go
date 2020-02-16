package summary

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
)

// DB set when init db and set it to handler
var DB *sql.DB

// Handler public since in router we set handler to route, when init router
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	summary := CountByDomain{}
	sqlStmt := "select domain, count(*) count from positions group by domain having domain = '%s'"

	err := DB.QueryRow(fmt.Sprintf(sqlStmt, vars["domain_name"])).Scan(&summary.Domain, &summary.Count)
	if err != nil {
		log.Println("Error duraing access to db!")
	}

	bytes, err := json.Marshal(summary)
	if err != nil {
		log.Println("Can't serialize", summary)
	}

	w.Write(bytes)
}
