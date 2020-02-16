package main

import (
    summary "app/get_summary" 
    "fmt"
    "context"
    "flag"
    "log"
    "net/http"
    "os"
    "os/signal"
    "time"
    "encoding/json"

	"github.com/gorilla/mux"

    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

const Timeout = 15
const IdleTimeout = 15
const Host = "0.0.0.0:3000"


func main() {
    var wait time.Duration
    flag.DurationVar(&wait, "graceful-timeout", time.Second * Timeout, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
    flag.Parse()

    r := mux.NewRouter()
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "This is a website server by a Go HTTP server.")
    })

    r.HandleFunc("/summary", summary.Handler).Methods("GET")
    r.HandleFunc("/positions/{domain_name}", positions).Methods("GET")

    srv := &http.Server{
        Addr:         Host,
        // Good practice to set timeouts to avoid Slowloris attacks.
        WriteTimeout: time.Second * Timeout,
        ReadTimeout:  time.Second * Timeout,
        IdleTimeout:  time.Second * IdleTimeout,
        Handler: r, // Pass our instance of gorilla/mux in.
    }

    // Run our server in a goroutine so that it doesn't block.
    go func() {
        if err := srv.ListenAndServe(); err != nil {
            log.Println(err)
        }
    }()

    c := make(chan os.Signal, 1)
    // We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
    // SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
    signal.Notify(c, os.Interrupt)

    // Block until we receive our signal.
    <-c

    // Create a deadline to wait for.
    ctx, cancel := context.WithTimeout(context.Background(), wait)
    defer cancel()
    // Doesn't block if no connections, but will otherwise wait
    // until the timeout deadline.
    srv.Shutdown(ctx)
    // Optionally, you could run srv.Shutdown in a goroutine and block on
    // <-ctx.Done() if your application should wait for other services
    // to finalize based on context cancellation.
    log.Println("shutting down")
    os.Exit(0)
}

type Position struct {
    Keyword   string    `json:"keyword"`
    Position  int       `json:"position"`
    Url       string    `json:"url"`
    Volume    int       `json:"volume"`
    Results   int       `json:"results"`
    Updated   time.Time `json:"updated"`
}

type DomainWithPosition struct {
    Domain    string        `json:"keyword"`
    Positions []Position    `json:"positions"`
}


func positions(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    fmt.Println(vars["domain_name"])

    db, err := sql.Open("sqlite3", "/Users/ernestkhasanzhinov/work/go/sample_server/db/positions.db")
    if err != nil {
        log.Fatalln(err)
    }

    positions := []Position{}
    sql_stmt := "select keyword, position, url, volume, results, updated from positions where domain = 'apostrophied.co.uk' order by volume asc limit 10 offset 10"
    rows, err := db.Query(sql_stmt)
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
