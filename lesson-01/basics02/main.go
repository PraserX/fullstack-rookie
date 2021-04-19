/**
BASICS 02 - Packages
====================

$ go run .
$ go run main.go
$ go build -o my_program .
$ go build -o my_program main.go
*/

package main

import (
	"fmt"       // https://golang.org/pkg/fmt/
	"math/rand" // https://golang.org/pkg/math/rand/
	"time"      // https://golang.org/pkg/time/
)

func Exported() {
	fmt.Println("Exported function")
}

func unexported() {
	fmt.Println("Unexported function")
}

func main() {
	fmt.Println("Hello world", rand.Intn(10))

	// If we working with rand package, we do not have to forget to seed
	rand.Seed(time.Now().UnixNano())
	// Now we get random number from 0 to 9
	fmt.Println("Hello world with seed", rand.Intn(10))
}
