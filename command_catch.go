package main

import (
	"fmt"
	"math/rand"

	"github.com/alexlangev/pokedex-cli/internal/pokedex"
)

func commandCatch(cfg *Config, arg string) error {
	fullURL := "https://pokeapi.co/api/v2/pokemon/" + arg

	info, err := cfg.pokeClient.PokemonInfo(fullURL, cfg.pokeCache)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", arg)

	score := rand.Intn(700)
	baseProb := (680 - info.BaseExperience)

	if score <= baseProb {
		pokedex.Pokedex[info.Name] = info
		fmt.Printf("%s was caught!\n", arg)
	} else {
		fmt.Printf("%s escaped!\n", arg)
	}

	return nil
}
