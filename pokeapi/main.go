package main

import (
	"context"
	"fmt"
	"log"

	"pokeapi/internal/api"
)

func main() {
	client := api.NewClient(api.NewConfig())

	locationAreas, err := client.ListLocationAreas(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}
}
