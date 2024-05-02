package repository

import (
	"database/sql"
	"github.com/google/uuid"
	"main/internal/model"
)

type ScooterOccupationRepository interface {
	Create(scooterId uuid.UUID, userId uuid.UUID) (created bool)
	DeleteByScooterIdAndUserId(scooterId uuid.UUID, userId uuid.UUID) (recordDeleted bool)
	GetByScooterIdAndUserId(scooterId uuid.UUID, userId uuid.UUID) (model.ScooterOccupation, bool)
}

type mysqlScooterOccupationRepository struct {
	db *sql.DB
}

func (r mysqlScooterOccupationRepository) Create(scooterId uuid.UUID, userId uuid.UUID) (created bool) {
	scooterIdAsBinary, _ := scooterId.MarshalBinary()
	userIdAsBinary, _ := userId.MarshalBinary()

	query := "INSERT INTO scooters_occupations (scooter_id, user_id) VALUES (?, ?)"
	_, err := r.db.Exec(query, scooterIdAsBinary, userIdAsBinary)

	if err != nil {
		return false
	}

	return true
}

func (r mysqlScooterOccupationRepository) DeleteByScooterIdAndUserId(
	scooterId uuid.UUID,
	userId uuid.UUID,
) (recordDeleted bool) {

	scooterIdAsBinary, _ := scooterId.MarshalBinary()
	userIdAsBinary, _ := userId.MarshalBinary()

	query := "DELETE FROM scooters_occupations WHERE scooter_id = ? AND user_id = ?"
	result, err := r.db.Exec(query, scooterIdAsBinary, userIdAsBinary)

	if err != nil {
		panic(err)
	}

	deletedRowsCount, _ := result.RowsAffected()

	return deletedRowsCount == 1
}

func (r mysqlScooterOccupationRepository) GetByScooterIdAndUserId(
	scooterId uuid.UUID,
	userId uuid.UUID,
) (model.ScooterOccupation, bool) {
	scooterIdAsBinary, _ := scooterId.MarshalBinary()
	userIdAsBinary, _ := userId.MarshalBinary()

	query := "SELECT * FROM scooters_occupations WHERE scooter_id = ? AND user_id = ?"

	var scooterOccupation model.ScooterOccupation
	err := r.db.QueryRow(query, scooterIdAsBinary, userIdAsBinary).Scan(
		&scooterOccupation.Id,
		&scooterOccupation.ScooterId,
		&scooterOccupation.UserId,
	)

	if err != nil {
		return scooterOccupation, false
	}

	return scooterOccupation, true
}

func NewScooterOccupationRepository(db *sql.DB) ScooterOccupationRepository {
	return &mysqlScooterOccupationRepository{db: db}
}
