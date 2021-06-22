package command

// StockBroker is the invoker
type StockBroker struct {
	order     Order
	prevOrder Order
}

// SetOrder sets the command to executed
func (se *StockBroker) SetOrder(o Order) {
	se.order = o
}

// ExecuteOrder executes the command. It also stores the command in prevOrder.
func (se *StockBroker) ExecuteOrder() string {
	se.prevOrder = se.order
	return se.order.Execute()
}

// UndoPrevOrder undo's the last invoked command
func (se *StockBroker) UndoPrevOrder() string {
	return se.prevOrder.Undo()
}
