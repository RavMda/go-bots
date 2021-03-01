package main

import (
	"go-bots/bot"
	. "go-bots/config"
	. "go-bots/guard"
	"go-bots/proxies"
	"net"

	"time"
)

var (
	guard Guard
	config *Config
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

func makeBot(address string) {
	conn, err := net.Dial("tcp", address)

	if err != nil {
		panic(err)
	}

	bot.Basic(conn, bot.Data{Dialer: net.Dial})
}

func main() {
	CreateConfig()
	CreateGuard()
	
	guard = GetGuard()
	config = GetConfig()

	for range time.Tick(config.Cooldown * time.Millisecond) {
		guard.Increment()

		if config.UseProxies {
			go proxyBot(proxies.GetProxy(), config.Address)
		} else {
			go makeBot(config.Address)
		}
	}
}
