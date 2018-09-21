package main

import (
	"log"
	"net/http"
	"os"

	"gitlab.com/koffee/little-bird/backend/api"
)

func main() {
	api.RegisterHandlers()

	port := "8080"
	if s := os.Getenv("PORT"); s != "" {
		port = s
	}
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("http.ListenAndServe: %v", err)
	}
}
