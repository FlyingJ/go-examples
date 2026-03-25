package api

type Config struct {
	BaseURL  string
	APIPath  string
	Next     string
	Previous string
}

func NewConfig() Config {
	domainname := "https://pokeapi.co"
	apiPath := "/api/v2"
	return Config{
		BaseURL:  domainname,
		APIPath:  apiPath,
		Next:     domainname + apiPath + "/location-area/",
		Previous: "",
	}
}
