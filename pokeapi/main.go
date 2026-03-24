package main

import (
	"pokeapi/internal/api"
	"pokeapi/internal/config"
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

	var locationAreas LocationAreaList
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locationAreas)
	if err != nil {
		log.Fatalf("failed to decode response body: %s", err)
	}

	if res.StatusCode > 299 {	
		log.Fatalf("response failed with status code: %d", res.StatusCode)
	}
}