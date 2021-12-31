package simple

type animal interface {
	speak() string
}

type dog struct{}

func (d dog) speak() string {
	return "Bhow Bhow!!"
}

type cat struct{}

func (c cat) speak() string {
	return "Meow Meow!!"
}
