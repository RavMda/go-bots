package main

import (
	"fmt"
	"go-pen/methods"
	"net"
	"os"
	"strings"
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

func main() {
	var serverAddress = checkArguments(os.Args[1:])
	var conn, err = net.Dial("tcp", serverAddress)

	if err != nil {
		fmt.Println("Something went wrong!", err)
	} else {
		fmt.Println("Connected!")

		data := methods.MethodData{Out: conn, Address: serverAddress, Loop: 5}
		methods.Extreme1(data)
	}
}
