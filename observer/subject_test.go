package observer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisterSubscriber(t *testing.T) {
	np := &NewsPublisher{}
	_ = NewSmsSubscriber(np)
	_ = NewEmailSubscriber(np)
	assert.Equal(t, 2, len(np.Subscribers))
}

func TestRemoveSubscriber(t *testing.T) {
	np := &NewsPublisher{}
	ss := NewSmsSubscriber(np)
	assert.Equal(t, 1, len(np.Subscribers))

	np.RemoveSubscriber(ss)
	assert.Equal(t, 0, len(np.Subscribers))
}

func TestNotifySubscribers(t *testing.T) {
	np := &NewsPublisher{}
	_ = NewSmsSubscriber(np)
	_ = NewEmailSubscriber(np)
	news := "news update"
	np.UpdateNews(news)

	assert.Equal(t, 2, updated)
}
