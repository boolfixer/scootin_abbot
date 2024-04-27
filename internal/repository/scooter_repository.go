package repository

import (
	"database/sql"
	"main/internal/model"
)

type ScooterRepository interface {
	FindScootersByArea(latitudeStart int, longitudeStart int, latitudeEnd int, longitudeEnd int) []model.Scooter
}

type mysqlScooterRepository struct {
	db *sql.DB
}

func (r mysqlScooterRepository) FindScootersByArea(
	latitudeStart int,
	longitudeStart int,
	latitudeEnd int,
	longitudeEnd int,
) []model.Scooter {

	query := "SELECT * " +
		"FROM scooters " +
		"WHERE scooters.latitude >= ? " +
		"AND scooters.longitude >= ? " +
		"AND scooters.latitude <= ? " +
		"AND scooters.longitude <= ?"

	rows, err := r.db.Query(query, latitudeStart, longitudeStart, latitudeEnd, longitudeEnd)

	if err != nil {
		panic(err)
	}

	var scooters []model.Scooter

	for rows.Next() {
		var scooter = model.Scooter{}
		err := rows.Scan(&scooter.Id, &scooter.Name, &scooter.Latitude, &scooter.Longitude)

		if err != nil {
			panic(err)
		}

		scooters = append(scooters, scooter)
	}

	return scooters
}

func NewScooterRepository(db *sql.DB) ScooterRepository {
	return &mysqlScooterRepository{db: db}
}
