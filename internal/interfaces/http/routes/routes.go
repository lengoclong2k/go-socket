package routes

import (
	"github.com/gin-gonic/gin"
	"go-chat-app/internal/interfaces/http/handlers"
	"go-chat-app/internal/interfaces/http/middlerware"
)

func SetupRoutes(router *gin.Engine, authHandler *handlers.AuthHandler, authMiddleware *middlerware.AuthMiddleware) {

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	v1 := router.Group("/api/v1")
	{
		// Auth routes (public)
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/refresh", authHandler.RefreshToken)
		}

		//Protected routes
		protected := v1.Group("")
		protected.Use(authMiddleware.RequireAuth())
		{
			//Room routes
			rooms := protected.Group("/rooms")
			{
				rooms.POST("")
			}
		}
	}
}
