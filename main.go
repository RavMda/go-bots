package main

import (
	"bufio"
	"fmt"
	"go-pen/methods"
	"log"
	"net"
	"os"
	"strings"

	"h12.io/socks"
)

func checkArguments(cleanArgs []string) string {
	if len(cleanArgs) == 0 {
		fmt.Println("Arguments: ip:port")
		os.Exit(1) // is it a good way to do that....
	}

	var serverAddress = cleanArgs[0]

	if !strings.Contains(serverAddress, ":") {
		fmt.Println("Arguments: ip:port")
		os.Exit(1)
	}

	return serverAddress
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

func makeConnection(proxyAddress string, serverAddress string) (net.Conn, error) {
	address := fmt.Sprintf("socks4://%s?timeout=4s", proxyAddress)

	dialSocks := socks.Dial(address)
	conn, err := dialSocks("tcp", serverAddress)

	return conn, err
}

func main() {
	var serverAddress = checkArguments(os.Args[1:])

	proxies := parseProxies()
	data := methods.MethodData{Address: serverAddress, Loop: 2}

	// Guard is needed to limit goroutines
	guard := make(chan struct{}, 10)

	for _, proxy := range proxies {
		guard <- struct{}{}
		go func(proxy string, serverAddress string, data *methods.MethodData) {
			var conn, err = makeConnection(proxy, serverAddress)

			if err != nil {
				fmt.Println(err)
				// remove proxy from file (maybe) and slice
			} else {
				methods.Extreme1(data, conn)
			}

			<-guard
		}(proxy, serverAddress, &data)
	}
}
