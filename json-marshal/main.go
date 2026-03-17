package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type listEntry struct {
	name string
	url string
}

type listResponse struct{
	next string
	previous string
	results []listEntry
}

func main() {
	url := "https://pokeapi.co/api/v2/location-area/"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("error while reading body: %s", err)
	}
	
	decoder := json.NewDecoder(res.Body)
	var results listResponse
	
	err = decoder.Decode(&results)
	if err != nil {
		log.Fatal("error while decoding response body: %v", err)
	}

	if res.StatusCode > 299 {
		log.Fatalf(
			"response failed with status code: %d and\nbody: %s\n",
			res.StatusCode,
			body,
		)
	}
	fmt.Printf("%v", results)
}
