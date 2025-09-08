package client

import (
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2/"

type Client struct {
	HttpClient *http.Client
	BaseURL    string
}

func NewClient() *Client {
	return &Client{
		HttpClient: &http.Client{Timeout: time.Second * 5}, // should be plenty
		BaseURL:    baseURL,
	}
}
