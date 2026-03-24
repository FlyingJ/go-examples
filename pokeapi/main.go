package main

import (
	"context"
	"fmt"
	"log"

	"pokeapi/internal/api"
	"pokeapi/internal/config"
)

func main() {
	var cfg config.Config
	cfg.Init()

	client := api.NewClient(cfg)

	locationAreas, err := client.ListLocationAreas(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}
}
