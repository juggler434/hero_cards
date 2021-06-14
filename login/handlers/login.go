package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//standin until routes are written
func handler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not yet implemented", http.StatusBadRequest)
}

type Config struct {
	db *sql.DB
}

func Start(port string) {

	r := mux.NewRouter()
	if port == "" {
		port = ":8080"
	}

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	r.HandleFunc("/users/new", handler).Methods("POST")

	// Register New User
	// Login with Created User

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}

}
