package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"pokeapi/internal/config"
)

// PokeAPI returns a few reusable response shapes across many endpoints.
type APIResource struct {
	URL string `json:"url"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type APIResourceList[T any] struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []T    `json:"results"`
}

type Client struct {
	cfg        config.Config
	httpClient *http.Client
}

type Endpoint struct {
	Named bool
	Path  string
}

var endpoints = map[string]Endpoint{
	"location":      {Path: "/location", Named: true},
	"location-area": {Path: "/location-area", Named: true},
	"pokemon":       {Path: "/pokemon", Named: true},
	"region":        {Path: "/region", Named: true},
}

func NewClient(cfg config.Config) *Client {
	return &Client{
		cfg:        cfg,
		httpClient: http.DefaultClient,
	}
}

func GetJSON[T any](ctx context.Context, c *Client, url string) (*T, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("perform request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	var value T
	if err := json.NewDecoder(res.Body).Decode(&value); err != nil {
		return nil, fmt.Errorf("decode response body: %w", err)
	}

	return &value, nil
}

func (c *Client) ListLocationAreas(ctx context.Context) (*APIResourceList[NamedAPIResource], error) {
	return GetJSON[APIResourceList[NamedAPIResource]](ctx, c, c.cfg.Next)
}
