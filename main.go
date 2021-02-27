package main

import (
	"go-bots/bot"
	. "go-bots/config"
	. "go-bots/guard"
	"go-bots/proxies"

	"time"
)

var (
	guard  = GetGuard()
	config = GetConfig()
)

func proxyBot(proxy string, address string) {
	conn, dialer, err := proxies.Dial(proxy, address)

	if err != nil {
		config.Cooldown = config.FastCooldown
		guard.Decrement()
		return
	}

	config.Cooldown = config.SlowCooldown
	bot.Basic(conn, bot.Data{Dialer: dialer})
}

func main() {
	for range time.Tick(config.Cooldown * time.Millisecond) {
		guard.Increment()
		go proxyBot(proxies.GetProxy(), config.Address)
	}
}
