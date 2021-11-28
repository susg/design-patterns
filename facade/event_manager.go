package facade

// eventManager is our facade.
type eventManager struct {
	h hotelier
	f florist
	c caterer
}

func newEventManager() eventManager {
	return eventManager{
		h: hotelier{},
		f: florist{},
		c: caterer{},
	}
}

// The arrange() method encapsulates all the subsystems needed to arrange a
// marriage.
func (em eventManager) arrange() (arrangements []string) {
	arrangements = append(arrangements, em.h.bookHotel())
	arrangements = append(arrangements, em.f.setFlowerRequirements())
	arrangements = append(arrangements, em.c.setCuisine())
	return
}
