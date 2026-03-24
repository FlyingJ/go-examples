package api

// the Pokeapi is pleasant enough to provide a few simple response
// structures
type APIResource struct{
	Url      string `json:"url"`
}
type NamedAPIResource struct{
	Name     string `json:"name"`
	Url      string `json:"url"`
}
type APIResourceList[T any] struct{
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []T    `json:"results"`
}

type Client struct {
	cfg        Config
	httpClient *http.Client
}

type Endpoint struct{
	Named bool
	Path  string
}

var endpoints = map[string]Endpoint{
	"location": {Path: "/location", Named: true}
	"location-area":  {Path: "/location-area", Named: true},
	"pokemon": {Path: "/pokemon", Named: true},
	"region": {Path: "/region", Named: true},
}

func GetJSON[T any](ctx context.Context, c *Client, ep Endpoint) (*T, error) {
	// build URL
	// make request
	// check status
	// decode into T
	// return &value
	return nil, fmt.Errorf("not implemented error")
}
