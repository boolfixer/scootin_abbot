package scooter_handler

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"main/internal/dto"
	"main/internal/http_error"
	mock_repository "main/internal/mock/repository"
	"main/internal/model"
	"reflect"
	"testing"
	"time"
)

func TestUpdateScooterLocation(t *testing.T) {
	t.Run("Scooter is not occupied by current user.", func(t *testing.T) {
		scooterId := uuid.UUID{}
		userId := uuid.UUID{}
		scooterLocationUpdate := dto.ScooterLocationUpdate{}
		var scooterOccupation model.ScooterOccupation

		scooterRepository := mock_repository.NewMockScooterRepository(gomock.NewController(t))
		scooterOccupationRepository := mock_repository.NewMockScooterOccupationRepository(gomock.NewController(t))
		scooterOccupationRepository.
			EXPECT().
			GetByScooterIdAndUserId(scooterId, userId).
			Return(scooterOccupation, false)

		handler := NewUpdateScooterLocation(scooterOccupationRepository, scooterRepository)
		err := handler.Handle(scooterId, userId, scooterLocationUpdate)

		want := reflect.TypeOf(http_error.ConflictError{}).Name()
		got := reflect.TypeOf(err).Name()
		if want != got {
			t.Errorf("failed to assert error; got %q, wanted %q", got, want)
		}

		want = "Scooter is not occupied by current user."
		got = err.Error()
		if want != got {
			t.Errorf("failed to assert error message; got %q, wanted %q", got, want)
		}
	})

	t.Run("Scooter not found.", func(t *testing.T) {
		scooterId := uuid.UUID{}
		userId := uuid.UUID{}
		scooterLocationUpdate := dto.ScooterLocationUpdate{}
		var scooterOccupation model.ScooterOccupation
		var scooter model.Scooter

		scooterOccupationRepository := mock_repository.NewMockScooterOccupationRepository(gomock.NewController(t))
		scooterOccupationRepository.
			EXPECT().
			GetByScooterIdAndUserId(scooterId, userId).
			Return(scooterOccupation, true)

		scooterRepository := mock_repository.NewMockScooterRepository(gomock.NewController(t))
		scooterRepository.
			EXPECT().
			GetByScooterId(scooterId).
			Return(scooter, false)

		handler := NewUpdateScooterLocation(scooterOccupationRepository, scooterRepository)
		err := handler.Handle(scooterId, userId, scooterLocationUpdate)

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

	t.Run("Scooter location is outdated.", func(t *testing.T) {
		scooterId := uuid.UUID{}
		userId := uuid.UUID{}
		yesterday := time.Now().Add(-24 * time.Hour)
		scooterLocationUpdate := dto.ScooterLocationUpdate{Time: yesterday}
		var scooterOccupation model.ScooterOccupation
		scooter := model.Scooter{LocationUpdatedAt: time.Now()}

		scooterOccupationRepository := mock_repository.NewMockScooterOccupationRepository(gomock.NewController(t))
		scooterOccupationRepository.
			EXPECT().
			GetByScooterIdAndUserId(scooterId, userId).
			Return(scooterOccupation, true)

		scooterRepository := mock_repository.NewMockScooterRepository(gomock.NewController(t))
		scooterRepository.
			EXPECT().
			GetByScooterId(scooterId).
			Return(scooter, true)

		handler := NewUpdateScooterLocation(scooterOccupationRepository, scooterRepository)
		err := handler.Handle(scooterId, userId, scooterLocationUpdate)

		want := reflect.TypeOf(http_error.ConflictError{}).Name()
		got := reflect.TypeOf(err).Name()
		if want != got {
			t.Errorf("failed to assert error; got %q, wanted %q", got, want)
		}

		want = "Scooter location is outdated."
		got = err.Error()
		if want != got {
			t.Errorf("failed to assert error message; got %q, wanted %q", got, want)
		}
	})

	t.Run("Failed to update scooter coordinates.", func(t *testing.T) {
		scooterId := uuid.UUID{}
		userId := uuid.UUID{}
		tomorrow := time.Now().Add(24 * time.Hour)
		scooterLocationUpdate := dto.ScooterLocationUpdate{Latitude: 10, Longitude: 15, Time: tomorrow}
		var scooterOccupation model.ScooterOccupation
		scooter := model.Scooter{LocationUpdatedAt: time.Now()}

		scooterOccupationRepository := mock_repository.NewMockScooterOccupationRepository(gomock.NewController(t))
		scooterOccupationRepository.
			EXPECT().
			GetByScooterIdAndUserId(scooterId, userId).
			Return(scooterOccupation, true)

		scooterRepository := mock_repository.NewMockScooterRepository(gomock.NewController(t))
		scooterRepository.
			EXPECT().
			GetByScooterId(scooterId).
			Return(scooter, true)
		scooterRepository.
			EXPECT().
			UpdateScooterCoordinatesByScooterId(scooterId, scooterLocationUpdate.Latitude, scooterLocationUpdate.Longitude).
			Return(errors.New("test error"))

		handler := NewUpdateScooterLocation(scooterOccupationRepository, scooterRepository)
		err := handler.Handle(scooterId, userId, scooterLocationUpdate)

		want := reflect.TypeOf(http_error.ConflictError{}).Name()
		got := reflect.TypeOf(err).Name()
		if want != got {
			t.Errorf("failed to assert error; got %q, wanted %q", got, want)
		}

		want = "Failed to update scooter coordinates."
		got = err.Error()
		if want != got {
			t.Errorf("failed to assert error message; got %q, wanted %q", got, want)
		}
	})

	t.Run("Scooter location updated successfully.", func(t *testing.T) {
		scooterId := uuid.UUID{}
		userId := uuid.UUID{}
		tomorrow := time.Now().Add(24 * time.Hour)
		scooterLocationUpdate := dto.ScooterLocationUpdate{Latitude: 10, Longitude: 15, Time: tomorrow}
		var scooterOccupation model.ScooterOccupation
		scooter := model.Scooter{LocationUpdatedAt: time.Now()}

		scooterOccupationRepository := mock_repository.NewMockScooterOccupationRepository(gomock.NewController(t))
		scooterOccupationRepository.
			EXPECT().
			GetByScooterIdAndUserId(scooterId, userId).
			Return(scooterOccupation, true)

		scooterRepository := mock_repository.NewMockScooterRepository(gomock.NewController(t))
		scooterRepository.
			EXPECT().
			GetByScooterId(scooterId).
			Return(scooter, true)
		scooterRepository.
			EXPECT().
			UpdateScooterCoordinatesByScooterId(scooterId, scooterLocationUpdate.Latitude, scooterLocationUpdate.Longitude).
			Return(nil)

		handler := NewUpdateScooterLocation(scooterOccupationRepository, scooterRepository)
		err := handler.Handle(scooterId, userId, scooterLocationUpdate)

		if err != nil {
			t.Errorf("unexpected error; got %T, wanted %T", err, nil)
		}
	})
}
