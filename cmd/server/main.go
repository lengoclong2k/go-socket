package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "go-chat-app/docs"
	"go-chat-app/internal/infrastructure/database"
	infraRepo "go-chat-app/internal/infrastructure/repositories"
	"go-chat-app/internal/interfaces/http/handlers"
	"go-chat-app/internal/interfaces/http/middlerware"
	"go-chat-app/internal/interfaces/http/routes"
	"go-chat-app/internal/usecases"
	"go-chat-app/pkg/config"
	"log"
)

// @title			User Service
// @version		1.0
// @description	Swagger for API Chat app.
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:8080
// @BasePath		/v1
// @schemes		http https
func main() {
	cfg := config.Load()

	db, err := database.NewPostgresDB(cfg.DatabaseURL())
	if err != nil {
		log.Fatal("Failed to connect to database::", err)
	}

	// Initialize repositories
	userRepo := infraRepo.NewUserRepository(db)

	// Initialize use cases
	authUseCase := usecases.NewAuthUseCase(userRepo, cfg.JWTSecret)

	//Initialize handlers
	authHandler := handlers.NewAuthHandler(authUseCase)

	// Initialize middleware
	authMiddleware := middlerware.NewAuthMiddleware(authUseCase)

	// Setup router
	router := gin.Default()

	//Setup Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.BasePath = "/api/v1"
	// CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Setup routes
	routes.SetupRoutes(router, authHandler, authMiddleware)

	// Start server
	log.Printf("Server starting on port %s", cfg.ServerPort)
	log.Fatal(router.Run(":" + cfg.ServerPort))
}
