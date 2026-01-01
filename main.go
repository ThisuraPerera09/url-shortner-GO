package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"url-shortener/config"
	"url-shortener/handlers"
	"url-shortener/middleware"
	"url-shortener/service"
	"url-shortener/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize storage
	var store storage.Storage
	var err error

	// Check for PostgreSQL connection string first (Railway provides this)
	if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
		fmt.Println("🚀 Using PostgreSQL storage")
		store, err = storage.NewPostgresStorage(dbURL)
		if err != nil {
			log.Fatalf("Failed to initialize PostgreSQL: %v", err)
		}
	} else if cfg.UseInMemory {
		fmt.Println("🚀 Using in-memory storage")
		store = storage.NewInMemoryStorage()
	} else {
		fmt.Printf("🚀 Using SQLite storage: %s\n", cfg.DatabasePath)
		store, err = storage.NewSQLiteStorage(cfg.DatabasePath)
		if err != nil {
			log.Fatalf("Failed to initialize database: %v", err)
		}
	}

	// Initialize service
	urlService := service.NewURLService(store, cfg.ShortCodeLen)

	// Initialize handlers
	urlHandler := handlers.NewURLHandler(urlService, cfg.BaseURL)

	// Setup Gin router
	router := setupRouter(urlHandler)

	// Graceful shutdown
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		fmt.Println("\n🛑 Shutting down server...")
		urlService.Close()
		os.Exit(0)
	}()

	// Start server
	fmt.Printf("🌐 Server running on http://localhost:%s\n", cfg.Port)
	fmt.Printf("📝 Base URL: %s\n", cfg.BaseURL)
	fmt.Println("✨ URL Shortener API ready!")

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupRouter(handler *handlers.URLHandler) *gin.Engine {
	// Set to release mode for production
	// gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// Middleware
	router.Use(middleware.Logger())
	router.Use(middleware.CORS())

	// API routes
	api := router.Group("/api")
	{
		api.GET("/health", handler.HealthCheck)
		api.POST("/shorten", handler.ShortenURL)
		api.GET("/stats/:shortCode", handler.GetStats)
		api.GET("/urls", handler.ListURLs)
		api.DELETE("/urls/:shortCode", handler.DeleteURL)
	}

	// Redirect route (must be last to avoid conflicts)
	router.GET("/:shortCode", handler.RedirectURL)

	return router
}
