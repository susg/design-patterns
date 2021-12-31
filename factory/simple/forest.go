package simple

type forest struct {
	factory forestFactory
}

func newForest(f forestFactory) forest {
	return forest{factory: f}
}

func (f forest) makeSound(a string) string {
	animal := f.factory.createAnimal(a)
	return animal.speak()
}
