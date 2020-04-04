package bikeindex

import "github.com/dghubble/sling"

// BaseURL is the base URL for the Bike Index API
const BaseURL = "https://bikeindex.org/api/v3/"

// Config configures a Client
type Config struct {
	URL   string
	Token string
}

// Client is a bikeindex.org API client
type Client struct {
	Manufacturers *manufacturersService
	Bikes         *bikeService
}

// New creates a new bikeindex.org API client
func New(config Config) *Client {
	s := sling.New().Base(config.URL).Set("Authorization", "Bearer "+config.Token)

	return &Client{
		Manufacturers: newManufacturerService(s.New()),
		Bikes:         newBikeService(s.New()),
	}
}
