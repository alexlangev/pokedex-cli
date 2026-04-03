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
