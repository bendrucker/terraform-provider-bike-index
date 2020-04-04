package bikeindex

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestManufacturersGet(t *testing.T) {
	manufacturers := newManufacturerService(sling.New().Base(BaseURL))
	manufacturer, err := manufacturers.Get("cannondale")

	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "Cannondale", manufacturer.Name)
}
