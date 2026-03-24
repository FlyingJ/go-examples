package config

type Config struct {
	BaseURL  string
	APIPath  string
	Next     string
	Previous string
}

func (c *Config) Init() {
	c.BaseURL = "https://pokeapi.co"
	c.APIPath = "/api/v2"
	c.Next = c.BaseURL + c.APIPath + "/location-area/"
	c.Previous = ""
}
