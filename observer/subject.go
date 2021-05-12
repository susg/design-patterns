package observer

// Subject Interface
type Publisher interface {
	RegisterSubscriber(Subscriber)
	RemoveSubscriber(Subscriber)
	NotifySubscribers()
}

// Concrete Subject
type NewsPublisher struct {
	Subscribers []Subscriber
	News        string
}

func (np *NewsPublisher) RegisterSubscriber(s Subscriber) {
	np.RemoveSubscriber(s) // removing first so that only one instance of subscriber is present in slice
	np.Subscribers = append(np.Subscribers, s)
}

func (np *NewsPublisher) RemoveSubscriber(s Subscriber) {
	found := false
	i := 0
	for ; i < len(np.Subscribers); i++ {
		if np.Subscribers[i] == s {
			found = true
			break
		}
	}

	if found {
		np.Subscribers = append(np.Subscribers[:i], np.Subscribers[i+1:]...)
	}
}

func (np NewsPublisher) NotifySubscribers() {
	for _, subscriber := range np.Subscribers {
		subscriber.Update(np.News)
	}
}

func (np NewsPublisher) UpdateNews(n string) {
	np.News = n
	np.NotifySubscribers()
}
