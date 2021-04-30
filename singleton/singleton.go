package singleton

import "sync"

type Singleton struct{}

var (
	instance *Singleton
	once     sync.Once
	wg       sync.WaitGroup
)

func GetSingletonInstance(instanceInfo chan string) *Singleton {
	defer wg.Done()
	once.Do(func() {
		instance = &Singleton{}
		instanceInfo <- "creating instance"
	})

	return instance
}
