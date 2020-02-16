package initialize_config

import (
	positions "app/get_positions"
	summary "app/get_summary"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is a website server by a Go HTTP server.")
	})

	db := InitDB()

	positions.DB = db
	summary.DB = db
	r.HandleFunc("/summary/{domain_name}", summary.Handler).Methods("GET")
	r.HandleFunc("/positions/{domain_name}", positions.Handler).Methods("GET")

	return r
}
