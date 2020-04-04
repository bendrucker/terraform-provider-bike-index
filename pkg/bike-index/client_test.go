package bikeindex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	client := New(BaseURL)
	assert.NotNil(t, client)
}
