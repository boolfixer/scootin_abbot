package repository

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"main/internal/model"
)

type ScooterRepository interface {
	FindScootersByStatusAndLocation(status bool, latitude float64, longitude float64) []model.Scooter
	SetScooterStatus(scooterUuid uuid.UUID, status bool)
}

type mysqlScooterRepository struct {
	db *sql.DB
}

func (r *mysqlScooterRepository) FindScootersByStatusAndLocation(status bool, latitude float64, longitude float64) []model.Scooter {
	return []model.Scooter{}
}

func (r *mysqlScooterRepository) SetScooterStatus(scooterUuid uuid.UUID, status bool) {
	return
}

func NewScooterRepository() ScooterRepository {
	// @todo: fetch config from env file
	cfg := mysql.Config{
		//User:   os.Getenv("DBUSER"),
		User: "root",
		//Passwd: os.Getenv("DBPASS"),
		Passwd: "root",
		Net:    "tcp",
		Addr:   "127.0.0.1:33306",
		DBName: "scootin_aboot",
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		panic(err.Error())
	}

	return &mysqlScooterRepository{db: db}
}
