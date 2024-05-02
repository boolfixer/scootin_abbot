package scooter_handler

import (
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"main/internal/dto"
	mock_repository "main/internal/mock/repository"
	"main/internal/model"
	"testing"
)

func TestSearchScootersHandler(t *testing.T) {
	location := dto.Location{Latitude: 15, Longitude: 13}

	latitudeStart := 12
	longitudeStart := 10
	latitudeEnd := 18
	longitudeEnd := 16

	scooter := model.Scooter{Id: uuid.UUID{}, Name: "Test scooter", Latitude: 14, Longitude: 16, IsOccupied: false}

	scooterRepository := mock_repository.NewMockScooterRepository(gomock.NewController(t))
	scooterRepository.
		EXPECT().
		FindScootersByArea(latitudeStart, longitudeStart, latitudeEnd, longitudeEnd).
		Return([]model.Scooter{scooter})

	handler := NewSearchScootersHandler(scooterRepository)
	scooters := handler.Handle(location)

	if len(scooters) != 1 {
		t.Errorf("failed to assert scooters count; got %d, wanted %d", len(scooters), 1)
	}

	if scooter != scooters[0] {
		t.Errorf("unexpected scooter record returned")
	}
}
