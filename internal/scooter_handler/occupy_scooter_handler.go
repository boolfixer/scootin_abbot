package scooter_handler

import (
	"github.com/google/uuid"
	"main/internal/http_error"
	"main/internal/repository"
)

type OccupyScooterHandler struct {
	scooterOccupationRepository repository.ScooterOccupationRepository
	scooterRepository           repository.ScooterRepository
}

func (h OccupyScooterHandler) Handle(scooterUuid uuid.UUID, userUuid uuid.UUID) error {
	scooter, exists := h.scooterRepository.GetByScooterId(scooterUuid)

	if !exists {
		return http_error.NotFoundError{ModelName: "Scooter"}
	}

	if scooter.IsOccupied || !h.scooterOccupationRepository.Create(scooterUuid, userUuid) {
		return http_error.ConflictError{Message: "Scooter has been already occupied."}
	}

	return nil
}

func NewOccupyScooterHandler(
	scooterOccupationRepository repository.ScooterOccupationRepository,
	scooterRepository repository.ScooterRepository,
) *OccupyScooterHandler {
	return &OccupyScooterHandler{
		scooterOccupationRepository: scooterOccupationRepository,
		scooterRepository:           scooterRepository,
	}
}
