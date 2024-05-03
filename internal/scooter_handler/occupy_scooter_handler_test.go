package scooter_handler

import (
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"main/internal/http_error"
	mock_repository "main/internal/mock/repository"
	"main/internal/model"
	"reflect"
	"testing"
)

func TestOccupyScooterHandler(t *testing.T) {
	t.Run("scooter not found", func(t *testing.T) {
		var scooter model.Scooter
		scooterId := uuid.UUID{}
		userId := uuid.UUID{}

		scooterOccupationRepository := mock_repository.NewMockScooterOccupationRepository(gomock.NewController(t))

		scooterRepository := mock_repository.NewMockScooterRepository(gomock.NewController(t))
		scooterRepository.
			EXPECT().
			GetByScooterId(scooterId).
			Return(scooter, false)

		handler := NewOccupyScooterHandler(scooterOccupationRepository, scooterRepository)
		err := handler.Handle(scooterId, userId)

		assertError(t, reflect.TypeOf(http_error.NotFoundError{}).Name(), "Scooter not found.", err)
	})

	t.Run("failed to create scooter", func(t *testing.T) {
		scooterId := uuid.UUID{}
		userId := uuid.UUID{}
		scooter := model.Scooter{Id: scooterId, IsOccupied: true}

		scooterOccupationRepository := mock_repository.NewMockScooterOccupationRepository(gomock.NewController(t))
		scooterRepository := mock_repository.NewMockScooterRepository(gomock.NewController(t))
		scooterRepository.
			EXPECT().
			GetByScooterId(scooterId).
			Return(scooter, true)

		handler := NewOccupyScooterHandler(scooterOccupationRepository, scooterRepository)
		err := handler.Handle(scooterId, userId)

		assertError(t, reflect.TypeOf(http_error.ConflictError{}).Name(), "Scooter has been already occupied.", err)
	})

	t.Run("failed to occupy scooter", func(t *testing.T) {
		scooterId := uuid.UUID{}
		userId := uuid.UUID{}
		scooter := model.Scooter{Id: scooterId, IsOccupied: false}

		scooterRepository := mock_repository.NewMockScooterRepository(gomock.NewController(t))
		scooterRepository.
			EXPECT().
			GetByScooterId(scooterId).
			Return(scooter, true)

		scooterOccupationRepository := mock_repository.NewMockScooterOccupationRepository(gomock.NewController(t))
		scooterOccupationRepository.
			EXPECT().
			Create(scooterId, userId).
			Return(false)

		handler := NewOccupyScooterHandler(scooterOccupationRepository, scooterRepository)
		err := handler.Handle(scooterId, userId)

		assertError(t, reflect.TypeOf(http_error.ConflictError{}).Name(), "Scooter has been already occupied.", err)
	})

	t.Run("scooter successfully occupied", func(t *testing.T) {
		var scooter model.Scooter
		scooterId := uuid.UUID{}
		userId := uuid.UUID{}

		scooterRepository := mock_repository.NewMockScooterRepository(gomock.NewController(t))
		scooterRepository.
			EXPECT().
			GetByScooterId(scooterId).
			Return(scooter, true)

		scooterOccupationRepository := mock_repository.NewMockScooterOccupationRepository(gomock.NewController(t))
		scooterOccupationRepository.
			EXPECT().
			Create(scooterId, userId).
			Return(true)

		handler := NewOccupyScooterHandler(scooterOccupationRepository, scooterRepository)
		err := handler.Handle(scooterId, userId)

		if err != nil {
			t.Errorf("unexpected error; got %T, wanted %T", err, nil)
		}
	})
}

func assertError(t testing.TB, wantErrorType string, wantErrorMessage string, gotError error) {
	t.Helper()

	got := reflect.TypeOf(gotError).Name()
	if wantErrorType != got {
		t.Errorf("failed to assert error; got %q, wanted %q", got, wantErrorMessage)
	}

	got = gotError.Error()
	if wantErrorMessage != got {
		t.Errorf("failed to assert error message; got %q, wanted %q", got, wantErrorMessage)
	}
}
