package main

import (
	config "app/config"
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	timeout := getTimeoutFor("TIMEOUT")

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*timeout, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	router := config.InitRouter()

	srv := &http.Server{
		Addr: os.Getenv("HOST"),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * timeout,
		ReadTimeout:  time.Second * timeout,
		IdleTimeout:  time.Second * getTimeoutFor("IDLE_TIMEOUT"),
		Handler:      router, // Pass our instance of gorilla/mux in.
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
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}

func getTimeoutFor(str string) time.Duration {
	timeoutInt, err := strconv.Atoi(os.Getenv(str))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return time.Duration(timeoutInt)
}
