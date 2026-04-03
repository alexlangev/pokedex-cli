package main

import (
	"fmt"
)

func commandExplore(cfg *Config, arg string) error {
	fullURL := cfg.baseUrl + "location-area/" + arg

	fmt.Printf("Exploring %s...\n", arg)

	locDetails, err := cfg.pokeClient.ExploreLocationArea(fullURL, cfg.pokeCache)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")

	for _, pe := range locDetails.PokemonEncounters {
		fmt.Printf(" - %s\n", pe.Pokemon.Name)
	}

	return nil
}
