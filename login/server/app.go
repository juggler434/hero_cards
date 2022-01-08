package app

import (
	"database/sql"
	"net/http"
)

type server struct {
	router *http.ServeMux
	db     *sql.DB
}

// NewServer returns an instance of the loging server with a router and db
func NewServer(router *http.ServeMux, db *sql.DB) *server {
	return &server{
		router: router,
		db:     db,
	}
}
