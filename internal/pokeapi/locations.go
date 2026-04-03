package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/alexlangev/pokedex-cli/internal/pokecache"
)

type LocationAreas struct {
	Count     int        `json:"count"`
	Next      *string    `json:"next"`
	Previous  *string    `json:"previous"`
	Locations []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationDetails struct {
	ID       int `json:"id"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon Pokemon `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetLocationAreas(url string, cache *pokecache.Cache) (LocationAreas, error) {
	cachedData, ok := cache.Get(url)
	if ok {
		locationsRes := LocationAreas{}
		err := json.Unmarshal(cachedData, &locationsRes)
		if err != nil {
			return LocationAreas{}, err
		}

		return locationsRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreas{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreas{}, nil
	}

	cache.Add(url, data)

	locationsRes := LocationAreas{}
	err = json.Unmarshal(data, &locationsRes)
	if err != nil {
		return LocationAreas{}, err
	}

	return locationsRes, nil
}

func (c *Client) ExploreLocationArea(url string, cache *pokecache.Cache) (LocationDetails, error) {
	// check cache
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationDetails{}, nil
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationDetails{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationDetails{}, err
	}

	var locDetails LocationDetails
	err = json.Unmarshal(data, &locDetails)
	if err != nil {
		return LocationDetails{}, err
	}

	return locDetails, nil
}
