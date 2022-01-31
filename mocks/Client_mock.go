package mocks

import (
	"net/http"

	taxjar "github.com/taxjar/taxjar-go"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func Client(fn RoundTripFunc) taxjar.Config {
	client := taxjar.NewClient(taxjar.Config{
		APIKey:     "mock-api-key",
		HTTPClient: NewTestClient(fn),
	})
	return client
}
