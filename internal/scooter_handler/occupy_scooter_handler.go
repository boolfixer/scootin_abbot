package scooter_handler

import (
	"github.com/google/uuid"
	"main/internal/repository"
	"time"
)

type OccupyScooterHandler struct {
	scooterOccupationRepository repository.ScooterOccupationRepository
}

func (h OccupyScooterHandler) Handle(scooterUuid uuid.UUID, userUuid uuid.UUID) {
	h.scooterOccupationRepository.Create(scooterUuid, userUuid, time.Now())
}

func NewOccupyScooterHandler(scooterOccupationRepository repository.ScooterOccupationRepository) OccupyScooterHandler {
	return OccupyScooterHandler{scooterOccupationRepository: scooterOccupationRepository}
}
