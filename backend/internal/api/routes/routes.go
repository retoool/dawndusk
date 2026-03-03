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
	sleepScheduleRepo := repositories.NewSleepScheduleRepository(db)
	groupRepo := repositories.NewGroupRepository(db)
	decorationRepo := repositories.NewDecorationRepository(db)
	messageRepo := repositories.NewMessageRepository(db)

	// Initialize services
	authService := services.NewAuthService(userRepo, cfg)
	petService := services.NewPetService(petRepo, db)
	checkInService := services.NewCheckInService(checkInRepo, petService)
	sleepScheduleService := services.NewSleepScheduleService(sleepScheduleRepo)
	groupService := services.NewGroupService(groupRepo)
	decorationService := services.NewDecorationService(decorationRepo, petRepo)
	messageService := services.NewMessageService(messageRepo, userRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	checkInHandler := handlers.NewCheckInHandler(checkInService)
	petHandler := handlers.NewPetHandler(petService)
	sleepScheduleHandler := handlers.NewSleepScheduleHandler(sleepScheduleService)
	userHandler := handlers.NewUserHandler(userRepo)
	groupHandler := handlers.NewGroupHandler(groupService)
	decorationHandler := handlers.NewDecorationHandler(decorationService)
	messageHandler := handlers.NewMessageHandler(messageService)

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
				users.GET("/me", userHandler.GetProfile)
				users.PUT("/me", userHandler.UpdateProfile)
			}

			// Sleep schedule routes
			protected.GET("/sleep-schedule", sleepScheduleHandler.Get)
			protected.PUT("/sleep-schedule", sleepScheduleHandler.Update)

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

			// Group routes
			groups := protected.Group("/groups")
			{
				groups.POST("/", groupHandler.Create)
				groups.GET("/", groupHandler.List)
				groups.GET("/my", groupHandler.GetUserGroups)
				groups.POST("/join", groupHandler.Join)
				groups.GET("/:id", groupHandler.Get)
				groups.PUT("/:id", groupHandler.Update)
				groups.DELETE("/:id", groupHandler.Delete)
				groups.POST("/:id/leave", groupHandler.Leave)
				groups.GET("/:id/members", groupHandler.GetMembers)
				groups.DELETE("/:id/members/:userId", groupHandler.RemoveMember)
			}

			// Decoration routes
			decorations := protected.Group("/decorations")
			{
				decorations.GET("/", decorationHandler.ListDecorations)
				decorations.GET("/my", decorationHandler.GetUserDecorations)
				decorations.POST("/unlock", decorationHandler.UnlockDecoration)
				decorations.POST("/:id/equip", decorationHandler.EquipDecoration)
			}

			// Message routes
			messages := protected.Group("/messages")
			{
				messages.POST("/", messageHandler.SendMessage)
				messages.GET("/", messageHandler.GetMessages)
				messages.GET("/unread-count", messageHandler.GetUnreadCount)
				messages.GET("/conversation/:userId", messageHandler.GetConversation)
				messages.POST("/conversation/:userId/read", messageHandler.MarkConversationAsRead)
				messages.POST("/:id/read", messageHandler.MarkAsRead)
				messages.DELETE("/:id", messageHandler.DeleteMessage)
			}
		}
	}
}
