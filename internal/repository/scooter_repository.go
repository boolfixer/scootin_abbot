package repository

import (
	"database/sql"
	"github.com/google/uuid"
	"main/internal/model"
)

type ScooterRepository interface {
	FindScootersByArea(latitudeStart int, longitudeStart int, latitudeEnd int, longitudeEnd int) []model.Scooter
	UpdateScooterCoordinatesByScooterId(scooterId uuid.UUID, latitude int, longitude int) (updated bool)
	GetByScooterId(scooterId uuid.UUID) (model.Scooter, bool)
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

func (r mysqlScooterRepository) UpdateScooterCoordinatesByScooterId(
	scooterId uuid.UUID,
	latitude int,
	longitude int,
) (updated bool) {
	scooterIdAsBinary, _ := scooterId.MarshalBinary()

	query := "UPDATE scooters SET scooters.latitude = ?, scooters.longitude = ? WHERE scooters.id = ?"
	_, err := r.db.Exec(query, latitude, longitude, scooterIdAsBinary)

	if err != nil {
		panic(err)
	}

	return true
}

func (r mysqlScooterRepository) GetByScooterId(scooterId uuid.UUID) (model.Scooter, bool) {
	scooterIdAsBinary, err := scooterId.MarshalBinary()
	if err != nil {
		panic(err)
	}

	query := "SELECT * FROM scooters WHERE scooters.id = ?"

	var scooter model.Scooter
	err = r.db.QueryRow(query, scooterIdAsBinary).Scan(&scooter.Id, &scooter.Name, &scooter.Latitude, &scooter.Longitude)

	if err != nil {
		return scooter, false
	}

	return scooter, true
}

func NewScooterRepository(db *sql.DB) ScooterRepository {
	return &mysqlScooterRepository{db: db}
}
