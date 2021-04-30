package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSingletonInstance(t *testing.T) {
	instanceInfo := make(chan string)
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go GetSingletonInstance(instanceInfo)
	}

	go func() {
		wg.Wait()
		close(instanceInfo)
	}()

	instanceCreated := 0
	for range instanceInfo {
		instanceCreated += 1
	}

	assert.Equal(t, 1, instanceCreated)
}
