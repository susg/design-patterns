package observer

import "fmt"

// Observer Interface
type Subscriber interface {
	Update(string) int
}

// Concrete Observer-1
type SmsSubscriber struct {
	NewsPublisher Publisher
}

func NewSmsSubscriber(np Publisher) Subscriber {
	s := SmsSubscriber{NewsPublisher: np}
	np.RegisterSubscriber(s)
	return s
}

func (s SmsSubscriber) Update(news string) int {
	fmt.Println("SmsSubscriber : Update :  ", news)
	return 1
}

// Concrete Observer-2
type EmailSubscriber struct {
	NewsPublisher Publisher
}

func NewEmailSubscriber(np Publisher) Subscriber {
	e := EmailSubscriber{NewsPublisher: np}
	np.RegisterSubscriber(e)
	return e
}

func (e EmailSubscriber) Update(news string) int {
	fmt.Println("EmailSubscriber : Update :  ", news)
	return 1
}
