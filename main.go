package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {

	//  load env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	http.ListenAndServe(":8080", nil)
}
