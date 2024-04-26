package repository

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type ScooterOccupationRepository interface {
	Create(scooterUuid uuid.UUID, userUuid uuid.UUID, occupiedAt time.Time)
	SetReleasedAtByScooterUuidAndUserUuid(releasedAt time.Time, scooterUuid uuid.UUID, userUuid uuid.UUID)
}

type mysqlScooterOccupationRepository struct {
	db *sql.DB
}

func (r mysqlScooterOccupationRepository) Create(scooterUuid uuid.UUID, userUuid uuid.UUID, occupiedAt time.Time) {
	return
}

func (r mysqlScooterOccupationRepository) SetReleasedAtByScooterUuidAndUserUuid(releasedAt time.Time, scooterUuid uuid.UUID, userUuid uuid.UUID) {
	return
}

func NewScooterOccupationRepository(db *sql.DB) ScooterOccupationRepository {
	return &mysqlScooterOccupationRepository{db: db}
}
