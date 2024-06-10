package main

import (
	"log"
	"log/slog"
	"os"

	"gotu/bookstore/internal/config"
	"gotu/bookstore/internal/db"
	"gotu/bookstore/internal/handler"
	"gotu/bookstore/internal/repository"
	"gotu/bookstore/internal/server"
	"gotu/bookstore/internal/service"

	"github.com/joho/godotenv"
)

func main() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load configuration
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Set logging
	if cfg.LogLevel == "debug" {
		slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: false,
		})))
	} else {
		slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
			Level:     slog.LevelInfo,
			AddSource: false,
		})))
	}
	slog.Info("env loaded")
	slog.Info("config loaded")
	slog.Info("logging started")

	// Initialize database
	database, err := db.New(cfg.DbUrl)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Check database connection
	if err := database.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	slog.Info("database connected")

	// Init dependencies
	bookRepository := repository.NewBookRepository(database)
	bookService := service.NewBookService(bookRepository)

	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepository)
	apiPublicHandler := handler.NewApiHandler(bookService, userService)

	// Start server
	srv := server.NewServer(apiPublicHandler)
	if err := srv.Start(cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}
