package bikeindex

import (
	"github.com/dghubble/sling"
)

type Manufacturer struct {
	Name        string `json:"name"`
	CompanyURL  string `json:"company_url"`
	ID          int    `json:"id"`
	FrameMarker bool   `json:"frame_marker"`
	Image       string `json:"image"`
	Description string `json:"description"`
	ShortName   string `json:"short_name"`
	Slug        string `json:"slug"`
}

type manufacturersService struct {
	sling *sling.Sling
}

func newManufacturerService(sling *sling.Sling) *manufacturersService {
	return &manufacturersService{
		sling: sling.Path("manufacturers/"),
	}
}

type manufacturerResponse struct {
	Manufacturer *Manufacturer `json:"manufacturer"`
}

// Get gets a manufacturer, querying by ID or slug
func (s *manufacturersService) Get(query string) (*Manufacturer, error) {
	result := new(manufacturerResponse)
	errResponse := new(ErrorResponse)

	resp, err := s.sling.New().Get(query).Receive(result, errResponse)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errResponse
	}

	return result.Manufacturer, nil
}
