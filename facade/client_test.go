package facade

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrangeMarriage(t *testing.T) {
	a := arrangeMarriage()
	assert.Equal(t, 3, len(a))
	assert.Contains(t, a, "Hotel is booked")
	assert.Contains(t, a, "Flowers are set")
	assert.Contains(t, a, "Cuisine is set")
}
