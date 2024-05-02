package scooter_handler

import (
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"main/internal/http_error"
	mock_repository "main/internal/mock/repository"
	"reflect"
	"testing"
)

func TestReleaseScooterHandler(t *testing.T) {
	t.Run("scooter occupation not found", func(t *testing.T) {
		scooterId := uuid.UUID{}
		userId := uuid.UUID{}

		scooterOccupationRepository := mock_repository.NewMockScooterOccupationRepository(gomock.NewController(t))
		scooterOccupationRepository.EXPECT().DeleteByScooterIdAndUserId(scooterId, userId).Return(false)

		handler := NewReleaseScooterHandler(scooterOccupationRepository)
		err := handler.Handle(scooterId, userId)

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

	t.Run("scooter released successfully", func(t *testing.T) {
		scooterId := uuid.UUID{}
		userId := uuid.UUID{}

		scooterOccupationRepository := mock_repository.NewMockScooterOccupationRepository(gomock.NewController(t))
		scooterOccupationRepository.EXPECT().DeleteByScooterIdAndUserId(scooterId, userId).Return(true)

		handler := NewReleaseScooterHandler(scooterOccupationRepository)
		err := handler.Handle(scooterId, userId)

		if err != nil {
			t.Errorf("unexpected error; got %T, wanted %T", err, nil)
		}
	})
}
