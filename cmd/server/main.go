package main

import (
	"fmt"
	"log"

	httpServer "github/k-tsurumaki/quilldeck/internal/interfaces/http"
)

func main() {
	server := httpServer.NewServer()
	
	port := "8080"
	fmt.Printf("Starting QuillDeck server on port %s...\n", port)
	
	if err := server.Start(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}