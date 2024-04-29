package scooter_handler

import (
	"main/internal/dto"
	"main/internal/model"
	"main/internal/repository"
)

const offset = 3

type SearchScootersHandler struct {
	scooterRepository repository.ScooterRepository
}

func (h *SearchScootersHandler) Handle(userLocation dto.Location) []model.Scooter {
	// calculate start and end coordinates of search area
	latitudeStart := userLocation.Latitude - offset
	longitudeStart := userLocation.Longitude - offset
	latitudeEnd := userLocation.Latitude + offset
	longitudeEnd := userLocation.Longitude + offset

	return h.scooterRepository.FindScootersByArea(latitudeStart, longitudeStart, latitudeEnd, longitudeEnd)
}

func NewSearchScootersHandler(scooterRepository repository.ScooterRepository) *SearchScootersHandler {
	return &SearchScootersHandler{scooterRepository: scooterRepository}
}
