package config_test

import (
	"pokeapi/internal/config"
	"testing"
)

func TestConfig(t *testing.T) {
	var c config.Config
	c.Init()

	var expectation string

	expectation = "https://pokeapi.co"
	if c.BaseURL != expectation {
		t.Errorf("c.BaseURL not correctly set: %s should be %s", c.BaseURL, expectation)
	}

	expectation = "/api/v2"
	if c.APIPath != expectation {
		t.Errorf("c.APIPath not correctly set: %s should be %s", c.APIPath, expectation)
	}

	expectation = "https://pokeapi.co/api/v2/location-area/"
	if c.Next != expectation {
		t.Errorf("c.Next not correctly set: %s should be %s", c.Next, expectation)
	}

	expectation = ""
	if c.Previous != expectation {
		t.Errorf("c.Previous not correctly set: %s supposed to be %s", c.Previous, expectation)
	}
}