package command

// Order is the command interface
type Order interface {
	Execute() string
	Undo() string
}

// BuyOrder is a command
type BuyOrder struct {
	stock   Stock
	prevQty int
	qty     int
}

func NewBuyOrder(s Stock, q int) *BuyOrder {
	return &BuyOrder{stock: s, qty: q}
}

func (bo *BuyOrder) Execute() string {
	bo.prevQty = bo.qty
	return bo.stock.Buy(bo.qty)
}

func (bo *BuyOrder) Undo() string {
	return bo.stock.Sell(bo.prevQty)
}

// SellOrder is a command
type SellOrder struct {
	stock   Stock
	prevQty int
	qty     int
}

func NewSellOrder(s Stock, q int) *SellOrder {
	return &SellOrder{stock: s, qty: q}
}

func (so *SellOrder) Execute() string {
	so.prevQty = so.qty
	return so.stock.Sell(so.qty)
}

func (so *SellOrder) Undo() string {
	return so.stock.Buy(so.prevQty)
}
