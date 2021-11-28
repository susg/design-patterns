package facade

func arrangeMarriage() []string {
	em := newEventManager()
	return em.arrange()
}
