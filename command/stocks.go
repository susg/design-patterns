package command

import "fmt"

// Stock is the interface for receiver
type Stock interface {
	Buy(int) string
	Sell(int) string
}

// Stock1 is a receiver
type Stock1 struct{}

func (s Stock1) Buy(q int) string {
	return fmt.Sprintf("bought %v units of stock1", q)
}

func (s Stock1) Sell(q int) string {
	return fmt.Sprintf("sold %v units of stock1", q)
}

// Stock2 is a receiver
type Stock2 struct{}

func (s Stock2) Buy(q int) string {
	return fmt.Sprintf("bought %v units of stock2", q)
}

func (s Stock2) Sell(q int) string {
	return fmt.Sprintf("sold %v units of stock2", q)
}
