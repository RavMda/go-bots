package proxies

import (
	"bufio"
	"log"
	"os"
)

var (
	isLoaded bool
	scanner  *bufio.Scanner
)

func GetProxy() string {
	if !isLoaded {
		parseProxies()
	}

	if scanner.Scan() {
		return scanner.Text()
	}

	isLoaded = false
	return GetProxy()
}

func parseProxies() {
	file, err := os.Open("proxies.txt")

	if err != nil {
		log.Fatal("Something is wrong with proxies.txt, ", err)
	}

	scanner = bufio.NewScanner(file)
	isLoaded = true
}
