package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	// 0 is success
	os.Exit(0)
	return nil
}
