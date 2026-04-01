package main

import "fmt"

func commandMapb(cfg *Config) error {
	if cfg.prevUrl == nil {
		fmt.Println("No previous locations")
		return nil
	}

	locs, err := cfg.pokeClient.GetLocationAreas(*cfg.prevUrl)
	if err != nil {
		return err
	}

	// update config with new urls
	cfg.nextUrl = locs.Next
	cfg.prevUrl = locs.Previous

	for _, area := range locs.Locations {
		fmt.Println(area.Name)
	}

	return nil
}
