package adapter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_WindowsAdapter(t *testing.T) {
	p1 := LightningPort{}
	assert.Equal(t, "inserted into lightning port", InsertDevice(p1))

	p2 := USBPort{}
	adp := WindowsAdapter{port: p2}
	assert.Equal(t, "inserted into USB port", InsertDevice(adp))
}

func InsertDevice(mp MacPort) string {
	return mp.InsertIntoLightningPort()
}
