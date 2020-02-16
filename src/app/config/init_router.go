package config

import (
	positions "app/positions"
	summary "app/summary"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// InitRouter we used in main when create router and all routings and also call InitDb here and set DB to all handlers
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
