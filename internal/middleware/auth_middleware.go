package middleware

import (
	"github.com/gin-gonic/gin"
	"main/internal/repository"
	"net/http"
)

type AuthMiddleware struct {
	userRepository repository.UserRepository
}

func (m AuthMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-KEY")

		if apiKey == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user, exists := m.userRepository.GetByApiKey(apiKey)
		if !exists {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func NewAuthMiddleware(repository repository.UserRepository) *AuthMiddleware {
	return &AuthMiddleware{userRepository: repository}
}
