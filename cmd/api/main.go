package main

import (
	"gothwind/internal/server"
	"log"
)

func main() {
	server, err := server.New()
	if err != nil {
		log.Fatalf("Failed to initialize server: %s\n", err)
	}

	server.Connect()
}
