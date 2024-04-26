package repository

import (
	"database/sql"
	"github.com/google/uuid"
)

type ScooterOccupationRepository interface {
	Create(scooterUuid uuid.UUID, userUuid uuid.UUID)
	SetReleasedAtByScooterUuidAndUserUuid(scooterUuid uuid.UUID, userUuid uuid.UUID)
}

type mysqlScooterOccupationRepository struct {
	db *sql.DB
}

func (r mysqlScooterOccupationRepository) Create(scooterUuid uuid.UUID, userUuid uuid.UUID) {
	return
}

func (r mysqlScooterOccupationRepository) SetReleasedAtByScooterUuidAndUserUuid(scooterUuid uuid.UUID, userUuid uuid.UUID) {
	return
}

func NewScooterOccupationRepository(db *sql.DB) ScooterOccupationRepository {
	return &mysqlScooterOccupationRepository{db: db}
}
