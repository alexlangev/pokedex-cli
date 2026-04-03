package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *Config, arg string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	// 0 is success
	os.Exit(0)
	return nil
}
