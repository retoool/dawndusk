package routes

import (
	"github.com/dawndusk/backend/internal/api/handlers"
	"github.com/dawndusk/backend/internal/api/middlewares"
	"github.com/dawndusk/backend/internal/domain/repositories"
	"github.com/dawndusk/backend/internal/domain/services"
	"github.com/dawndusk/backend/internal/shared/config"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB, redis *redis.Client) {
	// Load config
	cfg := config.Load()

	// Apply global middlewares
	router.Use(middlewares.Logger())
	router.Use(middlewares.CORS())

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "DawnDusk API is running",
		})
	})

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	checkInRepo := repositories.NewCheckInRepository(db)
	petRepo := repositories.NewPetRepository(db)

	// Initialize services
	authService := services.NewAuthService(userRepo, cfg)
	petService := services.NewPetService(petRepo, db)
	checkInService := services.NewCheckInService(checkInRepo, petService)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	checkInHandler := handlers.NewCheckInHandler(checkInService)
	petHandler := handlers.NewPetHandler(petService)

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Auth routes (public)
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/refresh", authHandler.RefreshToken)
			auth.POST("/logout", authHandler.Logout)
		}

		// Protected routes
		protected := v1.Group("")
		protected.Use(middlewares.AuthMiddleware(cfg))
		{
			// User routes
			users := protected.Group("/users")
			{
				users.GET("/me", func(c *gin.Context) {
					userID, _ := middlewares.GetUserID(c)
					c.JSON(200, gin.H{"message": "Get current user", "user_id": userID})
				})
			}

			// Check-in routes
			checkins := protected.Group("/check-ins")
			{
				checkins.POST("/", checkInHandler.Create)
				checkins.GET("/", checkInHandler.GetList)
				checkins.GET("/today", checkInHandler.GetToday)
				checkins.GET("/stats", checkInHandler.GetStats)
			}

			// Pet routes
			pet := protected.Group("/pet")
			{
				pet.GET("/", petHandler.Get)
				pet.PUT("/", petHandler.Update)
				pet.GET("/decorations", petHandler.GetDecorations)
				pet.POST("/decorations/:id/equip", petHandler.EquipDecoration)
			}
		}
	}
}
