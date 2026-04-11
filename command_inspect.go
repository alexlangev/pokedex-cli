package main

import (
	"fmt"

	"github.com/alexlangev/pokedex-cli/internal/pokedex"
)

func commandInspect(cfg *Config, arg string) error {
	info, ok := pokedex.Pokedex[arg]
	if !ok {
		fmt.Printf("You have not caught %s\n", info.Name)
		return nil
	}

	fmt.Println("Name: ", info.Name)
	fmt.Println("Height: ", info.Height)
	fmt.Println("Weight: ", info.Weight)

	return nil
}
