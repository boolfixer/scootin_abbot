package scooter_handler

import (
	"github.com/google/uuid"
	"main/internal/dto"
	"main/internal/repository"
)

type ReleaseScooterHandler struct {
	scooterOccupationRepository repository.ScooterOccupationRepository
	scooterRepository           repository.ScooterRepository
}

func (h ReleaseScooterHandler) Handle(scooterUuid uuid.UUID, userUuid uuid.UUID, scooterLocation dto.Location) {
	h.scooterOccupationRepository.DeleteByScooterUuidAndUserUuid(scooterUuid, userUuid)
	h.scooterRepository.UpdateScooterCoordinatesByScooterId(scooterUuid, scooterLocation.Latitude, scooterLocation.Longitude)
}

func NewReleaseScooterHandler(
	scooterOccupationRepository repository.ScooterOccupationRepository,
	scooterRepository repository.ScooterRepository) *ReleaseScooterHandler {

	return &ReleaseScooterHandler{
		scooterOccupationRepository: scooterOccupationRepository,
		scooterRepository:           scooterRepository,
	}
}
