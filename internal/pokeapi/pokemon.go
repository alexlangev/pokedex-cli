package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/alexlangev/pokedex-cli/internal/pokecache"
)

type Pokemon struct {
	BaseExperience         int    `json:"base_experience,omitempty"`
	Height                 int    `json:"height,omitempty"`
	ID                     int    `json:"id,omitempty"`
	LocationAreaEncounters string `json:"location_area_encounters,omitempty"`
	Name                   string `json:"name,omitempty"`
	Species                struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"species"`
	Types []struct {
		Slot int `json:"slot,omitempty"`
		Type struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"type"`
	} `json:"types,omitempty"`
	Weight int `json:"weight,omitempty"`
}

func (c *Client) PokemonInfo(url string, cache *pokecache.Cache) (Pokemon, error) {
	cachedData, ok := cache.Get(url)
	if ok {
		var pokeInfo Pokemon
		err := json.Unmarshal(cachedData, &pokeInfo)
		if err != nil {
			return Pokemon{}, nil
		}
		return pokeInfo, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, nil
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, nil
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return Pokemon{}, nil
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, nil
	}

	cache.Add(url, data)

	var pokeInfo Pokemon
	err = json.Unmarshal(data, &pokeInfo)
	if err != nil {
		return Pokemon{}, nil
	}

	return pokeInfo, nil
}
