package repository

import (
	"database/sql"
	"github.com/google/uuid"
)

type ScooterOccupationRepository interface {
	Create(scooterUuid uuid.UUID, userUuid uuid.UUID)
	DeleteByScooterUuidAndUserUuid(scooterUuid uuid.UUID, userUuid uuid.UUID)
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

func (r mysqlScooterOccupationRepository) DeleteByScooterUuidAndUserUuid(scooterUuid uuid.UUID, userUuid uuid.UUID) {
	scooterUuidAsBinary, err := scooterUuid.MarshalBinary()
	if err != nil {
		panic(err)
	}

	userUuidAsBinary, err := userUuid.MarshalBinary()
	if err != nil {
		panic(err)
	}

	query := "DELETE FROM scooters_occupations WHERE scooter_id = ? AND user_id = ?"

	result, err := r.db.Exec(query, scooterUuidAsBinary, userUuidAsBinary)

	if err != nil {
		panic(err)
	}

	deletedRowsCount, err := result.RowsAffected()

	if err != nil {
		panic(err)
	}

	if deletedRowsCount != 0 {
		panic("No records are affected")
	}
}

func NewScooterOccupationRepository(db *sql.DB) ScooterOccupationRepository {
	return &mysqlScooterOccupationRepository{db: db}
}
