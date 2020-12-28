package config

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	url           = "https://raw.githubusercontent.com/lik78/Gravy-Reloaded/master/nicks.txt"
	downloadError = "Error encountered while downloading names.txt!"
)

var (
	isLoaded = false
	file     *os.File
	scanner  *bufio.Scanner
)

func GetName() string {
	if !isLoaded {
		parseNames()
	}

	if scanner.Scan() {
		return scanner.Text()
	}

	isLoaded = false
	return GetName()
}

func parseNames() {
	if checkFile() {
		if err := downloadNames(); err != nil {
			log.Fatal(downloadError, err)
		}
	}

	if err := openFile(); err != nil {
		log.Fatal(err)
	}

	isLoaded = true
}

func openFile() error {
	file, err := os.Open("names.txt")
	if err != nil {
		return err
	}

	scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	return nil
}

func checkFile() bool {
	_, err := os.Stat("names.txt")
	return os.IsNotExist(err)
}

func downloadNames() error {
	fmt.Println("Parsing names.txt from remote resource..")

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	out, err := os.Create("names.txt")
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, response.Body)
	if err != nil {
		return err
	}

	return nil
}
