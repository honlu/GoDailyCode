package singleton

import "sync"

type Singleton struct{}

var eagerInstance = &Singleton{}

func GetInstance() *Singleton {
	return eagerInstance
}

var (
	lazyInstance *Singleton
	once         sync.Once
)

func GetLazyInstance() *Singleton {
	once.Do(func() {
		lazyInstance = &Singleton{}
	})
	return lazyInstance
}
