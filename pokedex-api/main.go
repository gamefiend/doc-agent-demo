package main

import (
	"log"
	"net/http"

	"github.com/yourorg/doc-agent-demo/pokedex-api/pokedex"
)

func main() {
	// Initialize the in-memory store
	store := pokedex.NewStore()

	// Create handler context with the store
	handlerContext := &pokedex.HandlerContext{
		Store: store,
	}

	// Build the router
	router := pokedex.NewRouter(handlerContext)

	// Start the HTTP server
	log.Println("Starting Pokédex API server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
