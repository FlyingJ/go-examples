package main

import (
	"pokeapi/internal/api"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	err := api.listLocationAreas()
	if err != nil {
		log.Fatal(err)
	}

	type LocationArea struct {
		Name string `json:"name"`
		Url string `json:"url"`
	}
	type LocationAreaList struct {
		Count int `json:"count"`
		Next string `json:"next"`
		Previous string `json:"previous"`
		Results []LocationArea `json:"results"`
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