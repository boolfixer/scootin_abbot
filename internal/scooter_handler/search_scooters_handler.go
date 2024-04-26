package scooter_handler

import (
	"main/internal/model"
	"main/internal/repository"
)

type SearchScootersHandler struct {
	scooterRepository repository.ScooterRepository
}

func (h *SearchScootersHandler) Handle(latitude int, longitude int) []model.Scooter {
	return h.scooterRepository.FindScootersByStatusAndLocation(latitude, longitude)
}

func NewSearchScootersHandler(scooterRepository repository.ScooterRepository) *SearchScootersHandler {
	return &SearchScootersHandler{scooterRepository: scooterRepository}
}
