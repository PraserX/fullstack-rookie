/**
BASICS 09 - Read file example
=============================

$ go run .
$ go run main.go
$ go build -o my_program .
$ go build -o my_program main.go
*/

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var err error

	// Read file by ioutil
	var data []byte
	if data, err = ioutil.ReadFile("./test.txt"); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
	}

	fmt.Println(string(data))

	// Read file by bufio scanner
	var fd *os.File

	if fd, err = os.Open("./test.txt"); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
