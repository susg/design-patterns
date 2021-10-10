package template

type iTrip interface {
	day1Activities() string
	day2Activities() string
	setVehicleType() string
	isRestNeeded() bool //hook
	returnHome() string
}

type trip struct {
	iTrip iTrip
}

// makeItinerary is the template method.
// Structs which embed `trip` need to implement `iTrip`
func (t trip) makeItinerary() (itinerary []string) {
	itinerary = append(itinerary, t.iTrip.setVehicleType())
	itinerary = append(itinerary, t.iTrip.day1Activities())
	if t.iTrip.isRestNeeded() {
		itinerary = append(itinerary, "rest needed")
	} else {
		itinerary = append(itinerary, t.iTrip.day2Activities())
	}
	itinerary = append(itinerary, t.iTrip.returnHome())
	return
}
