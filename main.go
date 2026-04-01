package main

import (
	"time"

	"github.com/alexlangev/pokedex-cli/internal/pokeapi"
)

type Config struct {
	baseUrl    string
	pokeClient pokeapi.Client
	nextUrl    *string
	prevUrl    *string
}

func main() {
	const pokeAPIBaseURL = "https://pokeapi.co/api/v2/"
	const httpTimeout = 5 * time.Second

	pokeClient := pokeapi.NewClient(httpTimeout)

	// using pointer since we would be passing a value otherwise...
	cfg := &Config{
		baseUrl:    pokeAPIBaseURL,
		pokeClient: pokeClient,
	}

	repl(cfg)
}
