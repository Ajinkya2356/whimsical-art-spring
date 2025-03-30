package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"backend/database"
	"backend/handlers"
	"backend/middleware"
	"backend/repositories"

	_ "backend/docs" // Import the generated docs

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Backend API
// @version 1.0
// @description This is the API documentation for the backend service.
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	// Set up logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Set log level based on environment (can be overridden with env var)
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logger.Warnf("Invalid log level %s, defaulting to info", logLevel)
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)

	// Load .env file if it exists
	err = godotenv.Load()
	if err != nil {
		logger.Warn("No .env file found or error loading it")
	}

	// Initialize database
	db, err := database.InitializeDatabase()
	if err != nil {
		logger.WithError(err).Fatal("Failed to initialize database")
	}

	// Run migrations
	if err := database.RunMigrations(db); err != nil {
		logger.WithError(err).Fatal("Failed to run database migrations")
	}

	// Create repositories
	promptRepo := repositories.NewGormPromptRepository(db)
	categoryRepo := repositories.NewGormCategoryRepository(db)
	userRepo := repositories.NewGormUserRepository(db)

	logger.Info("Connected to database")

	// Create middleware
	rateLimiter := middleware.NewRateLimiter(db, logger)
	authMiddleware := middleware.NewAuthMiddleware(db)
	loggingMiddleware := middleware.NewLoggingMiddleware(logger)
	corsMiddleware := middleware.NewCORSMiddleware()

	// Create handlers
	healthHandler := handlers.NewHealthHandler()
	promptHandler := handlers.NewPromptHandler(promptRepo, categoryRepo, userRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryRepo)
	userHandler := handlers.NewUserHandler(userRepo)
	authHandler := handlers.NewAuthHandler(userRepo)

	// Create Gin router
	router := gin.Default()

	// Apply global middleware
	router.Use(loggingMiddleware.LogRequest())
	router.Use(corsMiddleware.CORS())

	// Serve static files from docs directory
	router.Static("/docs", "./docs")

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("/docs/swagger.json"),
		ginSwagger.DefaultModelsExpandDepth(1),
	))

	// API routes
	api := router.Group("/api")

	// Apply global rate limiter to all API routes
	api.Use(rateLimiter.LimitGlobalRequests())

	// Authenticated routes
	auth := api.Group("/auth")
	auth.Use(authMiddleware.RequireAuth())
	auth.GET("/me", userHandler.GetCurrentUser)
	auth.GET("/me/stats", userHandler.GetCurrentUserStats)
	auth.PUT("/me", userHandler.UpdateUser)

	// Prompts authenticated routes (only for write operations)
	authPrompts := api.Group("/prompts")
	authPrompts.Use(authMiddleware.RequireAuth())

	// Apply rate limiter only to POST requests (prompt creation)
	createPrompt := authPrompts.Group("")
	createPrompt.Use(rateLimiter.LimitPromptCreation())
	createPrompt.POST("", promptHandler.CreatePrompt)

	// Other prompt operations (not rate limited)
	authPrompts.PUT("/:id", promptHandler.UpdatePrompt)
	authPrompts.DELETE("/:id", promptHandler.DeletePrompt)
	authPrompts.POST("/:id/like", promptHandler.LikePrompt)

	// Categories authenticated routes (admin only in a real app)
	authCategories := api.Group("/categories")
	authCategories.Use(authMiddleware.RequireAuth())
	authCategories.POST("", categoryHandler.CreateCategory)
	authCategories.PUT("/:id", categoryHandler.UpdateCategory)
	authCategories.DELETE("/:id", categoryHandler.DeleteCategory)

	// Public routes with optional authentication
	public := api.Group("")
	public.Use(authMiddleware.Optional())
	public.GET("/health", healthHandler.Check)
	public.POST("/auth/authenticate", authHandler.Authenticate)
	public.POST("/auth/signup", authHandler.Signup)
	public.POST("/auth/login", authHandler.Login)
	public.GET("/categories", categoryHandler.GetCategories)
	public.GET("/categories/:id", categoryHandler.GetCategory)
	public.GET("/prompts", promptHandler.GetPrompts)
	public.GET("/prompts/:id", promptHandler.GetPrompt)
	public.GET("/categories/:id/prompts", promptHandler.GetPromptsByCategory)
	public.GET("/stats", userHandler.GetSystemStats)

	// Set up server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	logger.Infof("Server starting on port %s...", port)
	if err := router.Run("0.0.0.0:" + port); err != nil {
		logger.Fatal(err)
	}
}
