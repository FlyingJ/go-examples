package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	type LocationAreaPointer struct {
		Name string `json:"name"`
		Url string `json:"url"`
	}
	type LocationAreaPointers struct {
		Count int `json:"count"`
		Next string `json:"next"`
		Previous string `json:"previous"`
		Results []LocationAreaPointer `json:"results"`
	}

	url := "https://pokeapi.co/api/v2/location-area/"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	var locationAreaPointers LocationAreaPointers
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locationAreaPointers)
	if err != nil {
		log.Fatalf("failed to decode response body: %s", err)
	}

	if res.StatusCode > 299 {	
		log.Fatalf("response failed with status code: %d", res.StatusCode)
	}
	
	fmt.Println("---- Things ----")
	fmt.Printf("%s -> %d\n", "count", locationAreaPointers.Count)
	fmt.Printf("%s -> %s\n", "next page", locationAreaPointers.Next)
	fmt.Printf("%s -> %s\n", "previous page", locationAreaPointers.Previous)
	for _, result := range locationAreaPointers.Results {
		fmt.Printf("%s -> %s\n", "name", result.Name)
		fmt.Printf("%s -> %s\n", "url", result.Url)
	}
}