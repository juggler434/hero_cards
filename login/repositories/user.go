package repositories

import (
	"context"
	"database/sql"

	"github.com/juggler434/marvelServer/login/models"
)

type UserRepository interface {
	GetUser(username string) (*models.User, error)
	CreateUser(username, email, password string) error
	UpdateUser(username, email, password string) error
	DeleteUser(username string) error
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) GetUser(username string) (*models.User, error) {
	panic("implement me")
}

func (u *userRepository) CreateUser(ctx context.Context, username, email, password string) error {
	user, err := models.NewUser(email, username, password)
	if err != nil {
		return err
	}

	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO users (id, email, username, password) VALUES ($1, $2, $3, $4)", user.ID(), user.Email(), user.Username(), user.Password())
	if err != nil {
		return err
	}
	err = tx.Commit()
	return err
}

func (u *userRepository) UpdateUser(username, email, password string) error {
	panic("implement me")
}

func (u *userRepository) DeleteUser(username string) error {
	panic("implement me")
}

// NewUserRepository returns an instance of UserRepository
func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}
