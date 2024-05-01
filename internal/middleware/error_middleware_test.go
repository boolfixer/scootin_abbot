package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"main/internal/http_error"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockFieldError struct {
	validator.FieldError
	tag   string
	field string
	param string
}

func (e mockFieldError) Tag() string { return e.tag }

func (e mockFieldError) Field() string { return e.field }

func (e mockFieldError) Param() string { return e.param }

func TestErrorMiddleware(t *testing.T) {
	cases := []struct {
		testCaseName string
		error        error
		statusCode   int
		errorMessage string
	}{
		{
			"NotFoundError received",
			http_error.NotFoundError{ModelName: "Test model"},
			http.StatusNotFound,
			"{\"message\":\"Test model not found.\"}",
		},
		{
			"ConflictError received",
			http_error.ConflictError{Message: "Test conflict message."},
			http.StatusConflict,
			"{\"message\":\"Test conflict message.\"}",
		},
		{
			"ValidationErrors received",
			validator.ValidationErrors{
				mockFieldError{tag: "required", field: "field1"},
				mockFieldError{tag: "gte", field: "field2", param: "20"},
				mockFieldError{tag: "lte", field: "field3", param: "25"},
				mockFieldError{tag: "number", field: "field4"},
			},
			http.StatusBadRequest,
			"{\"field1\":\"Field is required\",\"field2\":\"Field should be greater or equal 20\",\"field3\":\"Field should be less or equal 25\",\"field4\":\"Provided value should be numeric\"}",
		},
		{
			"Unknown error received",
			errors.New("Unknown error."),
			http.StatusInternalServerError,
			"{\"message\":\"Server error.\"}",
		},
	}

	for _, c := range cases {
		t.Run(c.testCaseName, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(recorder)
			context.Error(c.error)

			errorMiddleware := NewErrorMiddleware()
			errorMiddleware.Handle()(context)

			if recorder.Code != c.statusCode {
				t.Fatalf("Unexpected status code; wanted %d, but got %d", c.statusCode, recorder.Code)
			}

			got := recorder.Body.String()
			if got != c.errorMessage {
				t.Fatalf("Unexpected error message; wanted %q, but got %q", c.errorMessage, got)
			}
		})
	}

	t.Run("No errors received", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(recorder)

		errorMiddleware := NewErrorMiddleware()
		errorMiddleware.Handle()(context)

		if recorder.Code != http.StatusOK {
			t.Fatalf("Unexpected status code; wanted %d, but got %d", http.StatusOK, recorder.Code)
		}
	})
}
