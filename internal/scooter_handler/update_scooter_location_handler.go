package scooter_handler

import (
	"github.com/google/uuid"
	"main/internal/dto"
)

type UpdateScooterLocationHandler struct {
}

func (h UpdateScooterLocationHandler) Handle(
	scooterUuid uuid.UUID,
	userUuid uuid.UUID,
	scooterLocationUpdate dto.ScooterLocationUpdate,
) error {
	// @todo: implement
	return nil
}

func NewUpdateScooterLocation() *UpdateScooterLocationHandler {
	return &UpdateScooterLocationHandler{}
}
