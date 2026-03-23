package main

import (
	"pokeapi/internal/api"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Config struct {
	BaseURL string
	APIPath string
}

type Client struct {
	cfg        Config
	httpClient *http.Client
}

type Endpoint struct{
	Named bool
	Path  string
}

func GetJSON[T any](ctx context.Context, c *Client, ep Endpoint) (*T, error) {
	// build URL
	// make request
	// check status
	// decode into T
	// return &value
	return nil, fmt.Errorf("not implemented error")
}

// the Pokeapi is pleasant enough to provide a few simple response
// structures
type APIResource struct{
	Url      string `json:"url"`
}
type NamedAPIResource struct{
	Name     string `json:"name"`
	Url      string `json:"url"`
}
type APIResourceList[T any] struct{
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []T    `json:"results"`
}

var endpoints = map[string]Endpoint{
	"location": {Path: "/location", Named: true}
	"location-area":  {Path: "/location-area", Named: true},
	"pokemon": {Path: "/pokemon", Named: true},
	"region": {Path: "/region", Named: true},
}

func main() {
	err := api.listLocationAreas()
	if err != nil {
		log.Fatal(err)
	}

	var locationAreas LocationAreaList
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locationAreas)
	if err != nil {
		log.Fatalf("failed to decode response body: %s", err)
	}

	if res.StatusCode > 299 {	
		log.Fatalf("response failed with status code: %d", res.StatusCode)
	}
	
	fmt.Printf("Response from %s\n", url)
	fmt.Printf("count    -> %d\n", locationAreas.Count)
	fmt.Printf("next     -> %s\n", locationAreas.Next)
	fmt.Printf("previous -> %s\n", locationAreas.Previous)
	for _, locationArea := range locationAreas.Results {
		fmt.Printf("name: %s -> url: %s\n", locationArea.Name, locationArea.Url)
	}
}