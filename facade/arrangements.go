package facade

type hotelier struct{}

func (h hotelier) bookHotel() string {
	return "Hotel is booked"
}

type florist struct{}

func (f florist) setFlowerRequirements() string {
	return "Flowers are set"
}

type caterer struct{}

func (c caterer) setCuisine() string {
	return "Cuisine is set"
}
