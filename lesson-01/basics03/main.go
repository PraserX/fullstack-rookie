/**
BASICS 03 - Functions and variables
===================================

$ go run .
$ go run main.go
$ go build -o my_program .
$ go build -o my_program main.go
*/

package main

import (
	"fmt"
)

var x = 10          // unexported global variable
var Z = 15          // exported global variable
const Constant = 15 // exported constant

func add(a, b int) int {
	return a + b
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

func div2(a, b int) (x int, y int) {
	x = a / b
	y = a % b
	return x, y
}

func main() {

	fmt.Println("Add result:", add(5, 10))

	x, y := div(10, 7)
	fmt.Println("Div result:", x, y)

	x, y = div2(10, 7)
	fmt.Println("Div result:", x, y)

	// Declarations and definitions
	var a int  // declaration of local variable
	var b = 10 // definition of local variable
	c := 10    // definition of local variable
	fmt.Println(a, b, c)

	var d, e, f int
	var g, h int = 10, 12
	var i, j = 10, 15
	k, l := 10, 15
	m, n, o := true, false, "damn"
	fmt.Println(d, e, f, g, h, i, j, k, l, m, n, o)
}
