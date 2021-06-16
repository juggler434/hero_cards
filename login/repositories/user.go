package repositories

import (
	"database/sql"

	"github.com/juggler434/marvelServer/login/models"
)

type UserRepository interface {
	GetUser(username string) (models.User, error)
	CreateUser(username, email, password string) error
	UpdateUser(username, email, password string) error
	DeleteUser(username string) error
}

type userRepository struct {
	db sql.DB
}

func (u *userRepository) GetUser(username string) (models.User, error) {
	panic("implement me")
}

func (u *userRepository) CreateUser(username, email, password string) error {
	panic("implement me")
}

func (u *userRepository) UpdateUser(username, email, password string) error {
	panic("implement me")
}

func (u *userRepository) DeleteUser(username string) error {
	panic("implement me")
}

// NewUserRepository returns an instance of UserRepository
func NewUserRepository(db sql.DB) *userRepository {
	return &userRepository{db: db}
}
