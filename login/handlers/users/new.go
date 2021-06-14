package users

import "net/http"

type NewUserRequest struct {
	userName string
	email    string
	password string
}

func NewUserHandler(w http.ResponseWriter, r *http.Request) {

}
