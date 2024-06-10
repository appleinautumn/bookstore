package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"gotu/bookstore/internal/config"

	"github.com/joho/godotenv"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

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

	log.Println(cfg)

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

	http.HandleFunc("GET /", hello)
	http.ListenAndServe(fmt.Sprintf(":%s", cfg.AppPort), nil)
}
