package simple

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForest(t *testing.T) {
	f := newForest(concreteForestFactory{})
	assert.Equal(t, "Bhow Bhow!!", f.makeSound("dog"))
	assert.Equal(t, "Meow Meow!!", f.makeSound("cat"))
	assert.Panics(t, func() { f.makeSound("tiger") })
}
