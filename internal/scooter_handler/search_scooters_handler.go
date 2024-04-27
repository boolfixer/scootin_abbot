package scooter_handler

import (
	"main/internal/model"
	"main/internal/repository"
)

const offset = 3

type SearchScootersHandler struct {
	scooterRepository repository.ScooterRepository
}

func (h *SearchScootersHandler) Handle(latitude int, longitude int) []model.Scooter {
	// calculate start and end coordinates of search area
	latitudeStart := latitude - offset
	longitudeStart := longitude - offset
	latitudeEnd := latitude + offset
	longitudeEnd := longitude + offset

	return h.scooterRepository.FindScootersByArea(latitudeStart, longitudeStart, latitudeEnd, longitudeEnd)
}

func NewSearchScootersHandler(scooterRepository repository.ScooterRepository) *SearchScootersHandler {
	return &SearchScootersHandler{scooterRepository: scooterRepository}
}
