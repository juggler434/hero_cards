package login

import (
	"fmt"
	"log"
	"net/http"
)

func Start(port string) {
	if port == "" {
		port = ":8080"
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}

}
