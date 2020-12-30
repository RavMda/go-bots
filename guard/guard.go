package guard

import "go-pen/config"

type Guard chan struct{}

var (
	isCreated bool = false
	guard     chan struct{}
)

func GetGuard() Guard {
	if !isCreated {
		createGuard()
	}

	return guard
}

func createGuard() {
	config := config.GetConfig()
	guard = make(chan struct{}, config.Connections)

	isCreated = true
}

func (guard Guard) Increment() {
	guard <- struct{}{}
}

func (guard Guard) Decrement() {
	<-guard
}
