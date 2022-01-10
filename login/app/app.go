package app

import (
	"database/sql"
	"net/http"
)

type Server struct {
	router *http.ServeMux
	db     *sql.DB
}

func StartServer(port string) {
	s := &Server{
		router: http.NewServeMux(),
	}

	s.routes()
	http.ListenAndServe(port, s.router)
}
