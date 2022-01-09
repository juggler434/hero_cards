package app

import (
	"database/sql"
	"net/http"
)

type server struct {
	router *http.ServeMux
	db     *sql.DB
}
