package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// type resultsElement struct {
	// 	name string "json:`name`"
	// 	url string "json: `url`"
	// }
	type responseStruct struct {
		count int "json:`count`"
		// next string "json:`next`"
		// previous string "json:`previous`"
		// results []resultsElement "json:`results`"
	}

	url := "https://pokeapi.co/api/v2/location-area/"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}	
	var response responseStruct
	err = decoder.Decode(&response)
	if err != nil {
		log.Fatalf("error while decoding response body: %s", err)
	}

	res.Body.Close()

	if res.StatusCode > 299 {	
		log.Fatalf(
			"response failed with status code: %d and\nbody: %s\n",
			res.StatusCode,
			body,
		)
	}
	
	fmt.Println("---- Stuff ----")
	fmt.Println(body)
	fmt.Println("---- Things ----")
	fmt.Printf("count: %d\n", response["count"])
}