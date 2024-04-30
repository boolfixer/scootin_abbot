package scooter_handler

import (
	"github.com/google/uuid"
	"main/internal/dto"
	"main/internal/http_error"
	"main/internal/repository"
)

type ReleaseScooterHandler struct {
	scooterOccupationRepository repository.ScooterOccupationRepository
	scooterRepository           repository.ScooterRepository
}

func (h ReleaseScooterHandler) Handle(scooterUuid uuid.UUID, userUuid uuid.UUID, scooterLocation dto.Location) error {
	deleted := h.scooterOccupationRepository.DeleteByScooterUuidAndUserUuid(scooterUuid, userUuid)
	if !deleted {
		return http_error.NotFoundError{ModelName: "Scooter occupation"}
	}

	updated := h.scooterRepository.UpdateScooterCoordinatesByScooterId(
		scooterUuid,
		scooterLocation.Latitude,
		scooterLocation.Longitude,
	)

	if !updated {
		return http_error.NotFoundError{ModelName: "Scooter"}
	}

	return nil
}

func NewReleaseScooterHandler(
	scooterOccupationRepository repository.ScooterOccupationRepository,
	scooterRepository repository.ScooterRepository) *ReleaseScooterHandler {

	return &ReleaseScooterHandler{
		scooterOccupationRepository: scooterOccupationRepository,
		scooterRepository:           scooterRepository,
	}
}
