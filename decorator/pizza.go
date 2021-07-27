package decorator

// Pizza is the Component
type Pizza interface {
	Cost() int
	Description() string
}

// VegMania is the Concrete Component
type VegMania struct {
	description string
}

func NewVegMania() Pizza {
	return VegMania{description: "veg mania pizza"}
}

func (vm VegMania) Cost() int {
	return 300
}

func (vm VegMania) Description() string {
	return vm.description
}
