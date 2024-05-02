package scooter_handler

import (
	"github.com/google/uuid"
	"main/internal/http_error"
	"main/internal/repository"
)

type ReleaseScooterHandler struct {
	scooterOccupationRepository repository.ScooterOccupationRepository
}

func (h ReleaseScooterHandler) Handle(scooterId uuid.UUID, userId uuid.UUID) error {
	deleted := h.scooterOccupationRepository.DeleteByScooterIdAndUserId(scooterId, userId)

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
