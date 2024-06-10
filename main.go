package main

import (
	"fmt"
	"log"
	"net/http"

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

	http.HandleFunc("GET /", hello)
	http.ListenAndServe(fmt.Sprintf(":%s", cfg.AppPort), nil)
}
