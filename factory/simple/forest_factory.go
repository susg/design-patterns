package simple

type forestFactory interface {
	createAnimal(string) animal
}

type concreteForestFactory struct{}

func (f concreteForestFactory) createAnimal(name string) animal {
	if name == "dog" {
		return dog{}
	} else if name == "cat" {
		return cat{}
	} else {
		panic("invalid animal type")
	}
}
