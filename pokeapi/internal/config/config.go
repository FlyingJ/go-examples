package config

type Config struct{
	Next string
	Previous string
}

func (c *Config) Init() {
	c.Next = "https://pokeapi.co/api/v2/location-area/"
	c.Previous = ""
}