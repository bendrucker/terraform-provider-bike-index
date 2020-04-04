package bikeindex

import "github.com/dghubble/sling"

// BaseURL is the base URL for the Bike Index API
const BaseURL = "https://bikeindex.org/api/v3/"

// New creates a new bikeindex.org API client
func New(baseURL string) *Client {
	s := sling.New().Base(baseURL)

	return &Client{
		Manufacturers: newManufacturerService(s.New()),
	}
}

// Client is a bikeindex.org API client
type Client struct {
	Manufacturers *manufacturersService
}
