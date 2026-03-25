package api_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"pokeapi/internal/api"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func TestGetJSON(t *testing.T) {
	cfg := api.Config{
		Next: "https://example.test/api/v2/location-area/",
	}

	client := api.NewClient(cfg)
	client.HttpClient = &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			if r.Method != http.MethodGet {
				t.Fatalf("unexpected method: %s", r.Method)
			}

			if r.URL.String() != cfg.Next {
				t.Fatalf("unexpected URL: %s", r.URL.String())
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body: io.NopCloser(strings.NewReader(
					`{"count":1,"next":"","previous":"","results":[{"name":"canalave-city-area","url":"https://pokeapi.co/api/v2/location-area/1/"}]}`,
				)),
				Header: make(http.Header),
			}, nil
		}),
	}

	got, err := client.ListLocationAreas(context.Background())
	if err != nil {
		t.Fatalf("ListLocationAreas returned error: %v", err)
	}

	if got.Count != 1 {
		t.Fatalf("got count %d, want 1", got.Count)
	}

	if len(got.Results) != 1 {
		t.Fatalf("got %d results, want 1", len(got.Results))
	}

	if got.Results[0].Name != "canalave-city-area" {
		t.Fatalf("got result name %q, want %q", got.Results[0].Name, "canalave-city-area")
	}
}
