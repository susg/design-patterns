package decorator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Pizza(t *testing.T) {
	p := NewVegMania()
	p = NewCheeseTopping(p)
	assert.Equal(t, 350, p.Cost())
	assert.Equal(t, "cheese + veg mania pizza", p.Description())

	p = NewTomatoTopping(p)
	assert.Equal(t, 380, p.Cost())
	assert.Equal(t, "tomato + cheese + veg mania pizza", p.Description())
}
