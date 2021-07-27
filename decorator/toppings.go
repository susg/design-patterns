package decorator

// Topping is a helper struct for Decorator
type Topping struct {
	pizza Pizza
}

// Cheese is Concrete Decorator
type Cheese struct {
	Topping
}

func NewCheeseTopping(p Pizza) Pizza {
	return Cheese{Topping{p}}
}

func (cb Cheese) Cost() int {
	return cb.pizza.Cost() + 50
}

func (cb Cheese) Description() string {
	return "cheese + " + cb.pizza.Description()
}

// Tomato is Concrete Decorator
type Tomato struct {
	Topping
}

func NewTomatoTopping(p Pizza) Pizza {
	return Tomato{Topping{p}}
}

func (t Tomato) Cost() int {
	return t.pizza.Cost() + 30
}

func (t Tomato) Description() string {
	return "tomato + " + t.pizza.Description()
}
