package template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_makeItinerary(t *testing.T) {
	mp := newMalaysiaPackage("Delhi", false)
	iti := mp.makeItinerary()
	assert.Equal(t, 4, len(iti))
	assert.Equal(t, "need SUV", iti[0])
	assert.Equal(t, "Malaysia day1", iti[1])
	assert.Equal(t, "Malaysia day2", iti[2])
	assert.Equal(t, "return Delhi", iti[3])

	dp := newDubaiPackage("Mumbai", true)
	iti = dp.makeItinerary()
	assert.Equal(t, 4, len(iti))
	assert.Equal(t, "need Sedan", iti[0])
	assert.Equal(t, "Dubai day1", iti[1])
	assert.Equal(t, "rest needed", iti[2])
	assert.Equal(t, "return Mumbai", iti[3])

	dp = newDubaiPackage("Kolkata", false)
	iti = dp.makeItinerary()
	assert.Equal(t, 4, len(iti))
	assert.Equal(t, "need Sedan", iti[0])
	assert.Equal(t, "Dubai day1", iti[1])
	assert.Equal(t, "Dubai day2", iti[2])
	assert.Equal(t, "return Kolkata", iti[3])
}
