package guard

import . "go-bots/config"

type Guard chan struct{}

var (
	guard     chan struct{}
)

func GetGuard() Guard {
	return guard
}

func CreateGuard() {
	guard = make(chan struct{}, GetConfig().Connections)
}

func (guard Guard) Increment() {
	guard <- struct{}{}
}

func (guard Guard) Decrement() {
	<-guard
}
