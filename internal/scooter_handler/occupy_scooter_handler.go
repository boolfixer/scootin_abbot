package scooter_handler

import (
	"github.com/google/uuid"
	"main/internal/repository"
)

type OccupyScooterHandler struct {
	scooterOccupationRepository repository.ScooterOccupationRepository
	scooterRepository           repository.ScooterRepository
}

func (h OccupyScooterHandler) Handle(scooterUuid uuid.UUID, userUuid uuid.UUID) {
	h.scooterRepository.GetByScooterId(scooterUuid)
	h.scooterOccupationRepository.Create(scooterUuid, userUuid)
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
