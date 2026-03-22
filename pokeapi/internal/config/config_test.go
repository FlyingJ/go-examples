package config_test

import (
	"pokeapi/internal/config"
	"testing"
)

func TestConfig(t *testing.T) {
	var c config.Config
	c.Init()

	expectedUrl := "https://pokeapi.co/api/v2/location-area/"
	if c.Next != expectedUrl {
		t.Errorf("c.Next not correctly set: %s should be %s", c.Next, expectedUrl)
	}
	if c.Previous != "" {
	}
}