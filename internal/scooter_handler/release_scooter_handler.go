package scooter_handler

import (
	"github.com/google/uuid"
	"main/internal/dto"
	"main/internal/repository"
)

type ReleaseScooterHandler struct {
	scooterOccupationRepository repository.ScooterOccupationRepository
}

func (h ReleaseScooterHandler) Handle(scooterUuid uuid.UUID, userUuid uuid.UUID, scooterLocation dto.Location) {
	h.scooterOccupationRepository.DeleteByScooterUuidAndUserUuid(scooterUuid, userUuid)
}

func NewReleaseScooterHandler(scooterOccupationRepository repository.ScooterOccupationRepository) *ReleaseScooterHandler {
	return &ReleaseScooterHandler{scooterOccupationRepository: scooterOccupationRepository}
}
