package observer

// TODO: Think of a better way to check if news is updated
var updated = 0

// Observer Interface
type Subscriber interface {
	Update(string) string
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

func (s SmsSubscriber) Update(news string) string {
	updated += 1
	return news
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

func (e EmailSubscriber) Update(news string) string {
	updated += 1
	return news
}
