package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const PasswordCost = 10

type User struct {
	id       uuid.UUID
	email    string
	username string
	password []byte
}

func NewUser(email, username, password string) (*User, error) {
	id := uuid.New()
	pw, err := hashPassword(password)

	u := &User{
		id:       id,
		username: username,
		email:    email,
		password: pw,
	}
	return u, err
}

func (u *User) ID() uuid.UUID {
	return u.id
}

func (u *User) Email() string {
	return u.email
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) Username() string {
	return u.username
}

func (u *User) Password() []byte {
	return u.password
}

func (u *User) SetPassword(password string) error {
	ep, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	u.password = ep
	return nil
}

func hashPassword(plaintextPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(plaintextPassword), PasswordCost)
}
