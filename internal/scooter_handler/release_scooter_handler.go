package scooter_handler

import (
	"github.com/google/uuid"
	"main/internal/http_error"
	"main/internal/repository"
)

type ReleaseScooterHandler struct {
	scooterOccupationRepository repository.ScooterOccupationRepository
}

func (h ReleaseScooterHandler) Handle(scooterUuid uuid.UUID, userUuid uuid.UUID) error {
	deleted := h.scooterOccupationRepository.DeleteByScooterUuidAndUserUuid(scooterUuid, userUuid)

	if !deleted {
		return http_error.NotFoundError{ModelName: "Scooter occupation"}
	}

	return nil
}

func NewReleaseScooterHandler(
	scooterOccupationRepository repository.ScooterOccupationRepository,
) *ReleaseScooterHandler {
	return &ReleaseScooterHandler{
		scooterOccupationRepository: scooterOccupationRepository,
	}
}
