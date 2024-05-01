package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	mock_repository "main/internal/mock/repository"
	"main/internal/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthMiddleware(t *testing.T) {
	t.Run("auth header is not provided", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(recorder)
		context.Request = &http.Request{Header: make(http.Header)}

		userRepository := mock_repository.NewMockUserRepository(gomock.NewController(t))
		authMiddleware := NewAuthMiddleware(userRepository)

		authMiddleware.Handle()(context)

		want := http.StatusUnauthorized
		got := recorder.Code
		if want != got {
			t.Errorf("failed to assert response status code; got %d, wanted %d", got, want)
		}
	})

	t.Run("user not found", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(recorder)

		apiKey := "test api key"
		header := make(http.Header)
		header.Set("X-API-KEY", apiKey)
		context.Request = &http.Request{Header: header}

		userRepository := mock_repository.NewMockUserRepository(gomock.NewController(t))
		userRepository.EXPECT().GetByApiKey(apiKey).Return(model.User{}, false)
		authMiddleware := NewAuthMiddleware(userRepository)

		authMiddleware.Handle()(context)

		want := http.StatusUnauthorized
		got := recorder.Code
		if want != got {
			t.Errorf("failed to assert response status code; got %d, wanted %d", got, want)
		}
	})

	t.Run("user authorized successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(recorder)

		apiKey := "test api key"
		header := make(http.Header)
		header.Set("X-API-KEY", apiKey)
		context.Request = &http.Request{Header: header}

		user := model.User{Id: uuid.UUID{}, FirstName: "Bob", LastName: "Marley", ApiKey: apiKey}
		userRepository := mock_repository.NewMockUserRepository(gomock.NewController(t))
		userRepository.EXPECT().GetByApiKey(apiKey).Return(user, true)

		authMiddleware := NewAuthMiddleware(userRepository)

		authMiddleware.Handle()(context)

		want := http.StatusOK
		got := recorder.Code
		if want != got {
			t.Errorf("failed to assert response status code; got %d, wanted %d", got, want)
		}

		gotUser, exists := context.Get("user")

		if !exists {
			t.Error("no user found")
		}

		if user != gotUser {
			t.Error("failed to assert user")
		}
	})
}
