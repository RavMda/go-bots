package proxies

import (
	"bufio"
	"log"
	"os"
)

var (
	scanner *bufio.Scanner
)

func Prepare() {
	file, err := os.Open("proxies.txt")

	if err != nil {
		log.Fatal("Something is wrong with proxies.txt, ", err)
	}

	scanner = bufio.NewScanner(file)
}

func GetProxy() string {
	if !scanner.Scan() {
		Prepare()
	}

	return scanner.Text()
}
