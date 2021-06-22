package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuyOrder(t *testing.T) {
	broker := &StockBroker{}
	order := NewBuyOrder(Stock1{}, 5)
	broker.SetOrder(order)
	assert.Equal(t, "bought 5 units of stock1", broker.ExecuteOrder())

	order = NewBuyOrder(Stock2{}, 3)
	broker.SetOrder(order)
	assert.Equal(t, "bought 3 units of stock2", broker.ExecuteOrder())

	assert.Equal(t, "sold 3 units of stock2", broker.UndoPrevOrder())
}

func TestSellOrder(t *testing.T) {
	broker := &StockBroker{}
	order := NewSellOrder(Stock1{}, 5)
	broker.SetOrder(order)
	assert.Equal(t, "sold 5 units of stock1", broker.ExecuteOrder())

	order = NewSellOrder(Stock2{}, 3)
	broker.SetOrder(order)
	assert.Equal(t, "sold 3 units of stock2", broker.ExecuteOrder())

	assert.Equal(t, "bought 3 units of stock2", broker.UndoPrevOrder())
}
