package main

import (
	"time"

	"github.com/alexlangev/pokedex-cli/internal/pokeapi"
	"github.com/alexlangev/pokedex-cli/internal/pokecache"
)

type Config struct {
	baseUrl    string
	pokeClient pokeapi.Client
	pokeCache  *pokecache.Cache
	nextUrl    *string
	prevUrl    *string
}

func main() {
	const pokeAPIBaseURL = "https://pokeapi.co/api/v2/"
	const httpTimeout = 5 * time.Second
	const cacheDuration = 60 * time.Second

	pokeClient := pokeapi.NewClient(httpTimeout)
	pokeCache := pokecache.NewCache(cacheDuration)

	// using pointer since we would be passing a value otherwise...
	cfg := &Config{
		baseUrl:    pokeAPIBaseURL,
		pokeClient: pokeClient,
		pokeCache:  pokeCache,
	}

	repl(cfg)
}
