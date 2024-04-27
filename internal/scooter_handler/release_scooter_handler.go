package scooter_handler

import (
	"github.com/google/uuid"
	"main/internal/repository"
	"time"
)

type ReleaseScooterHandler struct {
	scooterOccupationRepository repository.ScooterOccupationRepository
}

func (h ReleaseScooterHandler) Handle(scooterUuid uuid.UUID, userUuid uuid.UUID) {
	h.scooterOccupationRepository.SetReleasedAtByScooterUuidAndUserUuid(time.Now(), scooterUuid, userUuid)
}

func NewReleaseScooterHandler(scooterOccupationRepository repository.ScooterOccupationRepository) *ReleaseScooterHandler {
	return &ReleaseScooterHandler{scooterOccupationRepository: scooterOccupationRepository}
}
