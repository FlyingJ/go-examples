package main

import (
	"context"
	"log"

	"pokeapi/internal/api"
)

func main() {
	client := api.NewClient(api.NewConfig())

	err := client.ListLocationAreas(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
