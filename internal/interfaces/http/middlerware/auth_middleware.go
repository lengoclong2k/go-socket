package middlerware

import (
	"github.com/gin-gonic/gin"
	"go-chat-app/internal/usecases"
	"go-chat-app/pkg/utils"
)

type AuthMiddleware struct {
	authUseCase *usecases.AuthUseCase
}

func NewAuthMiddleware(authUseCase *usecases.AuthUseCase) *AuthMiddleware {
	return &AuthMiddleware{
		authUseCase: authUseCase,
	}
}

func (m *AuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			utils.UnAuthorizeResponse(c, "Authorization header require", nil)
			c.Abort()
			return
		}
		token := authHeader[7:] //Remove "Bearer " prefix

		userID, err := m.authUseCase.ValidateToken(token)

		if err != nil {
			utils.UnAuthorizeResponse(c, "UnAuthorize", err)
		}

		c.Set("user_id", userID)
		c.Next()
	}

}
