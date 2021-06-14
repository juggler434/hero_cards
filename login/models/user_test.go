package models

import (
	"bytes"
	"reflect"
	"testing"

	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
)

var testUser = User{
	id:       uuid.New(), // This will be different every time, don't test against a specific value
	email:    "test@test.com",
	username: "SteveRodgers",
	password: []byte("notverysecret"),
}

func TestNewUser(t *testing.T) {
	type args struct {
		email    string
		username string
		password string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "base case",
			args: args{
				email:    "test@test.com",
				username: "SteveRodgers",
				password: "not@goodpA$$word",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			u, err := NewUser(tt.args.email, tt.args.username, tt.args.password)
			if err != nil {
				t.Errorf("Unexptected error when creating user: %s", err)
			}

			if u.id == uuid.Nil {
				t.Error("expected: NewUser to set id, got: nil")
			}

			if u.username != tt.args.username {
				t.Errorf("expected: NewUser to set username to: %s, got %s", tt.args.username, u.username)
			}

			if u.email != tt.args.email {
				t.Errorf("expected: NewUser to set email to: %s, got %s", tt.args.email, u.email)
			}

			if bytes.Equal(u.password, []byte(tt.args.password)) {
				t.Errorf("exptected: NewUser to encrypt password got %s", u.password)
			}

			if bcrypt.CompareHashAndPassword(u.password, []byte(tt.args.password)) != nil {
				t.Errorf("expected: password to match got: %s", err)
			}

		})
	}
}

func TestUser_Email(t *testing.T) {
	tests := []struct {
		name string
		user *User
		want string
	}{
		{
			name: "base case",
			user: &testUser,
			want: testUser.email,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := tt.user
			if got := u.Email(); got != tt.want {
				t.Errorf("Email() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Password(t *testing.T) {
	tests := []struct {
		name string
		user *User
		want []byte
	}{
		{
			name: "base case",
			user: &testUser,
			want: testUser.password,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.user.Password(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Password() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_SetEmail(t *testing.T) {
	tests := []struct {
		name  string
		user  *User
		email string
	}{
		{
			name:  "base case",
			user:  &testUser,
			email: "newEmail@test.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := tt.user
			u.SetEmail(tt.email)
			if u.email != tt.email {
				t.Errorf("SetEmail() set email to %v, want %v", u.email, tt.email)
			}
		})
	}
}

func TestUser_SetPassword(t *testing.T) {
	tests := []struct {
		name     string
		user     *User
		password string
	}{
		{
			name:     "base case",
			user:     &testUser,
			password: "differentPassword!1234",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := tt.user

			err := u.SetPassword(tt.password)
			if err != nil {
				t.Errorf("unexpected error when setting password %s", err)
			}
			if bcrypt.CompareHashAndPassword(u.password, []byte(tt.password)) != nil {
				t.Errorf("SetPassword() set password to %v, want %v", u.password, tt.password)
			}

		})
	}
}

func TestUser_Username(t *testing.T) {
	tests := []struct {
		name string
		user *User
		want string
	}{
		{
			name: "base case",
			user: &testUser,
			want: testUser.username,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := tt.user
			if got := u.Username(); got != tt.want {
				t.Errorf("Username() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_ID(t *testing.T) {
	tests := []struct {
		name string
		user *User
	}{
		{
			name: "base case",
			user: &testUser,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := tt.user
			if u.ID() == uuid.Nil {
				t.Error("ID() == nil, want not nil")
			}
		})
	}
}
