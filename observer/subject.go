package observer

// Subject Interface
type Publisher interface {
	RegisterSubscriber(Subscriber)
	RemoveSubscriber(Subscriber)
	NotifySubscribers() int
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

func (np NewsPublisher) NotifySubscribers() (notifiedCount int) {
	for _, subscriber := range np.Subscribers {
		notifiedCount += subscriber.Update(np.News)
	}
	return
}

func (np *NewsPublisher) UpdateNews(n string) int {
	np.News = n
	return np.NotifySubscribers()
}
