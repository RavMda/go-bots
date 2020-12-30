package guard

import "go-pen/config"

var (
	isCreated bool = false
	guard     chan struct{}
)

type Guard chan struct{}

func GetGuard() Guard {
	if !isCreated {
		createGuard()
	}

	return guard
}

func (guard Guard) Increment() {
	guard <- struct{}{}
}

func (guard Guard) Decrement() {
	<-guard
}

func createGuard() {
	config := config.GetConfig()
	guard = make(chan struct{}, config.Connections)

	isCreated = true
}
