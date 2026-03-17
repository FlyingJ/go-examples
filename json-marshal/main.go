package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	type resultsElement struct {
		name string
		url string
	}
	type responseStruct struct {
		count int
		next string
		previous string
		results []resultsElement
	}

	url := "https://pokeapi.co/api/v2/location-area/"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	
	var response responseStruct
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&response)
	if err != nil {
		log.Fatal("error while decoding response body: %v", err)
	}

	if res.StatusCode > 299 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Fatalf(
			"response failed with status code: %d and\nbody: %s\n",
			res.StatusCode,
			body,
		)
	}
	
	fmt.Println("---- Stuff ----")
	fmt.Printf("count: %d\n", response.count)
}