package scooter_handler

import (
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"main/internal/dto"
	"main/internal/http_error"
	mock_repository "main/internal/mock/repository"
	"reflect"
	"testing"
)

func TestReleaseScooterHandler(t *testing.T) {
	t.Run("scooter occupation not found", func(t *testing.T) {
		location := dto.Location{Latitude: 12, Longitude: 14}
		scooterId := uuid.UUID{}
		userId := uuid.UUID{}

		scooterRepository := mock_repository.NewMockScooterRepository(gomock.NewController(t))

		scooterOccupationRepository := mock_repository.NewMockScooterOccupationRepository(gomock.NewController(t))
		scooterOccupationRepository.EXPECT().DeleteByScooterUuidAndUserUuid(scooterId, userId).Return(false)

		handler := NewReleaseScooterHandler(scooterOccupationRepository, scooterRepository)
		err := handler.Handle(scooterId, userId, location)

		want := reflect.TypeOf(http_error.NotFoundError{}).Name()
		got := reflect.TypeOf(err).Name()
		if want != got {
			t.Errorf("failed to assert error; got %q, wanted %q", got, want)
		}

		want = "Scooter occupation not found."
		got = err.Error()
		if want != got {
			t.Errorf("failed to assert error message; got %q, wanted %q", got, want)
		}
	})

	t.Run("failed to update scooter occupation", func(t *testing.T) {
		location := dto.Location{Latitude: 12, Longitude: 14}
		scooterId := uuid.UUID{}
		userId := uuid.UUID{}

		scooterOccupationRepository := mock_repository.NewMockScooterOccupationRepository(gomock.NewController(t))
		scooterOccupationRepository.EXPECT().DeleteByScooterUuidAndUserUuid(scooterId, userId).Return(true)

		scooterRepository := mock_repository.NewMockScooterRepository(gomock.NewController(t))
		scooterRepository.
			EXPECT().
			UpdateScooterCoordinatesByScooterId(scooterId, location.Latitude, location.Longitude).
			Return(false)

		handler := NewReleaseScooterHandler(scooterOccupationRepository, scooterRepository)
		err := handler.Handle(scooterId, userId, location)

		want := reflect.TypeOf(http_error.NotFoundError{}).Name()
		got := reflect.TypeOf(err).Name()
		if want != got {
			t.Errorf("failed to assert error; got %q, wanted %q", got, want)
		}

		want = "Scooter not found."
		got = err.Error()
		if want != got {
			t.Errorf("failed to assert error message; got %q, wanted %q", got, want)
		}
	})

	t.Run("scooter released successfully", func(t *testing.T) {
		location := dto.Location{Latitude: 12, Longitude: 14}
		scooterId := uuid.UUID{}
		userId := uuid.UUID{}

		scooterOccupationRepository := mock_repository.NewMockScooterOccupationRepository(gomock.NewController(t))
		scooterOccupationRepository.EXPECT().DeleteByScooterUuidAndUserUuid(scooterId, userId).Return(true)

		scooterRepository := mock_repository.NewMockScooterRepository(gomock.NewController(t))
		scooterRepository.
			EXPECT().
			UpdateScooterCoordinatesByScooterId(scooterId, location.Latitude, location.Longitude).
			Return(true)

		handler := NewReleaseScooterHandler(scooterOccupationRepository, scooterRepository)
		err := handler.Handle(scooterId, userId, location)

		if err != nil {
			t.Errorf("unexpected error; got %T, wanted %T", err, nil)
		}
	})
}
