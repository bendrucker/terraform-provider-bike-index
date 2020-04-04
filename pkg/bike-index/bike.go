package bikeindex

import (
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
	ID               int    `json:"id"`
	Description      string `json:"description"`
	SerialNumber     string `json:"serial_number"`
	ComponentType    string `json:"component_type"`
	ComponentGroup   string `json:"component_group"`
	ManufacturerName string `json:"manufacturer_name"`
	ModelName        string `json:"model_name"`
	Year             string `json:"year"`
}

type bikeService struct {
	sling *sling.Sling
}

func newBikeService(sling *sling.Sling) *bikeService {
	return &bikeService{
		sling: sling.Path("bike/"),
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
