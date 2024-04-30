package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"main/internal/http_error"
	"net/http"
)

type ErrorMiddleware struct {
}

func (m ErrorMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {
			switch e := err.Err.(type) {

			case validator.ValidationErrors:
				c.AbortWithStatusJSON(http.StatusBadRequest, formatValidationErrors(e))
			case http_error.NotFoundError:
				c.AbortWithStatusJSON(http.StatusNotFound, map[string]string{"message": e.Error()})
			case http_error.ConflictError:
				c.AbortWithStatusJSON(http.StatusConflict, map[string]string{"message": e.Error()})
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"message": "Server error."})
			}
		}
	}
}

func formatValidationErrors(errors validator.ValidationErrors) map[string]string {
	result := make(map[string]string)

	for _, value := range errors {
		switch value.Tag() {
		case "required":
			result[value.Field()] = "Field is required"
		case "gte":
			result[value.Field()] = "Field should be greater or equal " + value.Param()
		case "lte":
			result[value.Field()] = "Field should be less or equal " + value.Param()
		case "number":
			result[value.Field()] = "Provided value should be numeric"
		}
	}

	return result
}

func NewErrorMiddleware() ErrorMiddleware {
	return ErrorMiddleware{}
}
