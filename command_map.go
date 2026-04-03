package main

import "fmt"

func commandMap(cfg *Config, arg string) error {
	const locationAreasEndpoint = "location-area/"

	fullURL := cfg.nextUrl
	if fullURL == nil {
		// reached the last page
		if cfg.prevUrl != nil {
			fmt.Println("No next locations")
		}
		url := cfg.baseUrl + locationAreasEndpoint
		fullURL = &url
	}

	locs, err := cfg.pokeClient.GetLocationAreas(*fullURL, cfg.pokeCache)
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
