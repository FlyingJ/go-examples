package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	protocol := "http"
	domain := "www.google.com"
	path := "/robots.txt"

	url := fmt.Sprintf(
		"%s://%s%s",
		protocol,
		domain,
		path,
	)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf(
			"Response failed with status code: %d and\nbody: %s\n",
			res.StatusCode,
			body,
		)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}