package main

import (
	"log"
	"net/http"
	"os"

	"github.com/golovers/little-bird/backend/handlers"
)

func main() {
	handlers.Register()

	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("http.ListenAndServe: %v", err)
	}
}
