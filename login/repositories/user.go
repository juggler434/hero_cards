package repositories

import "database/sql"

type UserRepository struct {
	db sql.DB
}
