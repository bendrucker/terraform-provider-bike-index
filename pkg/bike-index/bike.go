package bikeindex

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/dghubble/sling"
)

// Bike is a bike! On Bike Index.
type Bike struct {
	Description            string       `json:"description"`
	FrameColors            []string     `json:"frame_colors"`
	FrameModel             string       `json:"frame_model"`
	ID                     int          `json:"id"`
	Serial                 string       `json:"serial"`
	URL                    string       `json:"url"`
	Year                   string       `json:"year"`
	ManufacturerID         int          `json:"manufacturer_id"`
	Name                   string       `json:"name"`
	FrameSize              string       `json:"frame_size"`
	RearTireNarrow         bool         `json:"rear_tire_narrow"`
	FrontTireNarrow        bool         `json:"front_tire_narrow"`
	TestBike               bool         `json:"test_bike"`
	RearWheelSizeIsoBsd    string       `json:"rear_wheel_size_iso_bsd"`
	FrontWheelSizeIsoBsd   string       `json:"front_wheel_size_iso_bsd"`
	HandlebarTypeSlug      string       `json:"handlebar_type_slug"`
	FrameMaterialSlug      string       `json:"frame_material_slug"`
	FrontGearTypeSlug      string       `json:"front_gear_type_slug"`
	RearGearTypeSlug       string       `json:"rear_gear_type_slug"`
	AdditionalRegistration string       `json:"additional_registration"`
	Components             []*Component `json:"components"`
}

// Component is a bike component, anything that's not the frame
type Component struct {
	ID               int    `json:"id" url:"id,omitempty"`
	Description      string `json:"description" url:"description"`
	SerialNumber     string `json:"serial_number" url:"serial"`
	ComponentType    string `json:"component_type" url:"component_type"`
	ManufacturerName string `json:"manufacturer_name" url:"manufacturer"`
	ModelName        string `json:"model_name" url:"model"`
	Year             string `json:"year" url:"year"`
	Destroy          bool   `url:"destroy,omitempty"`
}

type bikeService struct {
	sling *sling.Sling
}

func newBikeService(sling *sling.Sling) *bikeService {
	return &bikeService{
		sling: sling.Path("bikes/"),
	}
}

type bikeResponse struct {
	Bike *Bike `json:"bike"`
}

// Get gets a bike by ID
func (s *bikeService) Get(id string) (*Bike, error) {
	result := new(bikeResponse)
	errResponse := new(ErrorResponse)

	resp, err := s.sling.New().Get(id).Receive(result, errResponse)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errResponse
	}

	return result.Bike, nil
}

// CreateUpdateBikeRequest creates or updates a bike
type CreateUpdateBikeRequest struct {
	OwnerEmail             string       `url:"owner_email,omitempty"`
	Description            string       `url:"description,omitempty"`
	Color                  string       `url:"color,omitempty"`
	PrimaryFrameColor      string       `url:"primary_frame_color,omitempty"`
	SecondaryFrameColor    string       `url:"secondary_frame_color,omitempty"`
	TertiaryFrameColor     string       `url:"tertiary_frame_color,omitempty"`
	FrameModel             string       `url:"frame_model,omitempty"`
	ID                     int          `url:"id,omitempty"`
	Serial                 string       `url:"serial,omitempty"`
	URL                    string       `url:"url,omitempty"`
	Year                   string       `url:"year,omitempty"`
	Manufacturer           string       `url:"manufacturer,omitempty"`
	Name                   string       `url:"name,omitempty"`
	FrameSize              string       `url:"frame_size,omitempty"`
	RearTireNarrow         bool         `url:"rear_tire_narrow,omitempty"`
	FrontTireNarrow        bool         `url:"front_tire_narrow,omitempty"`
	TestBike               bool         `url:"test_bike,omitempty"`
	RearWheelSize          string       `url:"rear_wheel_bsd,omitempty"`
	FrontWheelSize         string       `url:"front_wheel_size_bsd,omitempty"`
	HandlebarTypeSlug      string       `url:"handlebar_type_slug,omitempty"`
	FrameMaterialSlug      string       `url:"frame_material_slug,omitempty"`
	FrontGearTypeSlug      string       `url:"front_gear_type_slug,omitempty"`
	RearGearTypeSlug       string       `url:"rear_gear_type_slug,omitempty"`
	AdditionalRegistration string       `url:"additional_registration,omitempty"`
	Components             []*Component `url:"components,omitempty"`
}

// Create creates a new bike
func (s *bikeService) Create(bike *CreateUpdateBikeRequest) (*Bike, error) {
	if bike.ID != 0 {
		return nil, fmt.Errorf("cannot create bike with existing ID: %d", bike.ID)
	}

	result := new(bikeResponse)
	errResponse := new(ErrorResponse)

	resp, err := s.sling.New().Post("").BodyForm(bike).Receive(result, errResponse)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, errResponse
	}

	return result.Bike, nil
}

// Update updates a existing bike
func (s *bikeService) Update(bike *CreateUpdateBikeRequest) (*Bike, error) {
	if bike.ID == 0 {
		return nil, errors.New("cannot update bike without ID")
	}
	id := strconv.Itoa(bike.ID)

	result := new(bikeResponse)
	errResponse := new(ErrorResponse)

	resp, err := s.sling.New().Put(id).BodyForm(bike).Receive(result, errResponse)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, errResponse
	}

	return result.Bike, nil
}
