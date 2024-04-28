package repository

import (
	"database/sql"
	"github.com/google/uuid"
)

type ScooterOccupationRepository interface {
	Create(scooterUuid uuid.UUID, userUuid uuid.UUID)
	DeleteBy()
}

type mysqlScooterOccupationRepository struct {
	db *sql.DB
}

func (r mysqlScooterOccupationRepository) Create(scooterUuid uuid.UUID, userUuid uuid.UUID) {
	scooterUuidAsBinary, err := scooterUuid.MarshalBinary()
	if err != nil {
		panic(err)
	}

	userUuidAsBinary, err := userUuid.MarshalBinary()
	if err != nil {
		panic(err)
	}

	query := "INSERT INTO scooters_occupations (scooter_id, user_id) VALUES (?, ?)"
	_, err = r.db.Exec(query, scooterUuidAsBinary, userUuidAsBinary)

	if err != nil {
		panic(err)
	}
}

func (r mysqlScooterOccupationRepository) DeleteBy() {
	return
}

func NewScooterOccupationRepository(db *sql.DB) ScooterOccupationRepository {
	return &mysqlScooterOccupationRepository{db: db}
}
