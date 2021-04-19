/**
BASICS 05 - Defers and pointers
===============================

$ go run .
$ go run main.go
$ go build -o my_program .
$ go build -o my_program main.go
*/

package main

import "fmt"

func deferExample() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

func main() {
	// Defer example
	deferExample()

	// Pointers example
	i := 42

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	fmt.Println("Address", "Value")
	fmt.Println(&i, i)
	fmt.Println(p, *p)
}
