package scooter_handler

import (
	"github.com/google/uuid"
	"main/internal/repository"
)

type OccupyScooterHandler struct {
	scooterOccupationRepository repository.ScooterOccupationRepository
}

func (h OccupyScooterHandler) Handle(scooterUuid uuid.UUID, userUuid uuid.UUID) {
	h.scooterOccupationRepository.Create(scooterUuid, userUuid)
}

func NewOccupyScooterHandler(scooterOccupationRepository repository.ScooterOccupationRepository) *OccupyScooterHandler {
	return &OccupyScooterHandler{scooterOccupationRepository: scooterOccupationRepository}
}
