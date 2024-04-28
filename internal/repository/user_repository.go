package repository

import (
	"database/sql"
	"main/internal/model"
)

type UserRepository interface {
	GetByApiKey(apiKey string) model.User
}

type mysqlUserRepository struct {
	db *sql.DB
}

func (r mysqlUserRepository) GetByApiKey(apiKey string) model.User {
	query := "SELECT * FROM users WHERE users.api_key = ?"
	var user model.User

	err := r.db.QueryRow(query, apiKey).Scan(&user.Id, &user.FirstName, &user.LastName, &user.ApiKey)

	if err != nil {
		panic(err)
	}

	return user
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &mysqlUserRepository{db: db}
}
