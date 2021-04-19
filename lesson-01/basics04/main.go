/**
BASICS 04 - Loops, conditions and switches
==========================================

$ go run .
$ go run main.go
$ go build -o my_program .
$ go build -o my_program main.go
*/

package main

import (
	"fmt"
	"runtime"
)

func loop1() {
	fmt.Println("Loop example 1:")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func loop2() {
	fmt.Println("Loop example 2:")

	i := 0
	for i < 10 {
		fmt.Println(i)
		i += 1
	}

	// i := 0
	// for ; i < 10; {
	//     fmt.Println(i)
	//     i += 1
	// }
}

func loop3() {
	fmt.Println("Loop example 3:")
	i := 0
	for {
		// infinite loop
		fmt.Println(i)
		if i > 5 {
			break
		}
		i++
	}
}

func isLess(x int) bool {
	return x < 0
}

func conditions() {
	x := 0

	// Will print nothing (x = 0)
	if x < 0 {
		fmt.Println("x < 0")
	}

	// Will print nothing (x = 0)
	if res := isLess(x); res {
		fmt.Println("x < 0")
	}

	// Will print x >= 0
	if res := isLess(x); res {
		fmt.Println("x < 0")
	} else {
		fmt.Println("x >= 0")
	}

	// Will print x = 0
	if res := isLess(x); res {
		fmt.Println("x < 0")
	} else if x == 0 {
		fmt.Println("x = 0")
	} else {
		fmt.Println("x > 0")
	}
}

func main() {
	// Loop examples
	loop1()
	loop2()
	loop3()

	// If statement examples
	conditions()

	// Switch statement example
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os { // same as IF statement
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

}
