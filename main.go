package main

import (
	"fmt"
	"go-bots/bot/maidan"
	. "go-bots/config"
	. "go-bots/guard"
	"go-bots/proxies"
	"net"

	"time"
)

var (
	guard  Guard
	config *Config
)

func proxyBot(proxy string, address string) {
	conn, err := proxies.Dial(proxy, address)

	if err != nil {
		config.Cooldown = config.FastCooldown
		guard.Decrement()
		return
	}

	config.Cooldown = config.SlowCooldown
	maidan.CreateBot(conn)
}

func makeBot(address string) {
	conn, err := net.Dial("tcp", address)

	if err != nil {
		fmt.Println(err)
		return
	}

	maidan.CreateBot(conn)
}

func main() {
	CreateConfig()
	CreateGuard()

	guard = GetGuard()
	config = GetConfig()

	if config.UseProxies {
		proxies.Prepare()
	}

	for range time.Tick(config.Cooldown * time.Millisecond) {
		guard.Increment()

		if config.UseProxies {
			go proxyBot(proxies.GetProxy(), config.Address)
		} else {
			go makeBot(config.Address)
		}
	}
}
