package app

import (
	"database/sql"
	"net/http"
)

//Server hold all of the necessary components to run the login server
type Server struct {
	router *http.ServeMux
	db     *sql.DB
}

//StartServer starts the server listenting on the passed in port
// ex: login.StartServer(":8080")
func StartServer(port string) {
	s := &Server{
		router: http.NewServeMux(),
	}

	s.routes()
	http.ListenAndServe(port, s.router)
}
