package repository

import (
	"database/sql"
	"main/internal/model"
)

type UserRepository interface {
	GetByApiKey(apiKey string) (model.User, bool)
}

type mysqlUserRepository struct {
	db *sql.DB
}

func (r mysqlUserRepository) GetByApiKey(apiKey string) (model.User, bool) {
	query := "SELECT * FROM users WHERE users.api_key = ?"
	var user model.User

	err := r.db.QueryRow(query, apiKey).Scan(&user.Id, &user.FirstName, &user.LastName, &user.ApiKey)

	if err != nil {
		return user, false
	}

	return user, true
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &mysqlUserRepository{db: db}
}
