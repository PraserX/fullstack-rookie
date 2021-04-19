/**
BASICS 10 - Command line arguments
==================================

$ go run . test.txt
$ go run main.go test.txt
$ go build -o my_program .
$ go build -o my_program main.go
$ ./my_program test.txt
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var err error
	var fd *os.File
	args := os.Args
	if fd, err = os.Open(args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
