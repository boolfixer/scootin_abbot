package middleware

import (
	"github.com/gin-gonic/gin"
	"main/internal/repository"
)

type AuthMiddleware struct {
	userRepository repository.UserRepository
}

func (m AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-KEY")

		if apiKey == "" {
			panic("Authentication failed")
		}

		user := m.userRepository.GetByApiKey(apiKey)

		c.Set("user", user)
		c.Next()
	}
}

func NewAuthMiddleware(repository repository.UserRepository) *AuthMiddleware {
	return &AuthMiddleware{userRepository: repository}
}
