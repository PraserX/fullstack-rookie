/**
BASICS 07 - Arrays, Slices, Maps
================================

(!) Slices vs Arrays --> https://blog.golang.org/slices-intro

$ go run .
$ go run main.go
$ go build -o my_program .
$ go build -o my_program main.go
*/

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func main() {
	fmt.Printf("%s\n", "Arrays")
	var arr [2]string
	arr[0] = "Hello"
	arr[1] = "World"
	fmt.Println(arr[0], arr[1])
	fmt.Println(arr)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	fmt.Printf("\n\n%s\n", "Slices")
	arrayPrimes := [6]int{2, 3, 5, 7, 11, 13}

	var sliceFromArray []int = arrayPrimes[1:4]
	fmt.Println(sliceFromArray)

	slice := []int{2, 3, 5, 7, 11, 13}
	printSlice(slice)
	slice = slice[:0] // Slice the slice to give it zero length.
	printSlice(slice)
	slice = slice[:4] // Extend its length.
	printSlice(slice)
	slice = slice[2:] // Drop its first two values.
	printSlice(slice)

	fmt.Printf("\n\n%s\n", "Slices (append, pop, push)")
	var x int
	sliceOp := []int{2, 3, 5, 7, 11, 13}
	sliceOp = append(sliceOp, 15, 16) // append
	fmt.Println(sliceOp)
	x, sliceOp = sliceOp[0], sliceOp[1:] // pop
	fmt.Println(x, sliceOp)
	sliceOp = append([]int{10}, sliceOp...) // push
	fmt.Println(sliceOp)
	// s = append([]{10}, s)    // syntax error
	// fmt.Println(s)

	fmt.Printf("\n\n%s\n", "Slices (make)")
	a := make([]int, 5)
	printSlice2("a", a)

	b := make([]int, 0, 5)
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)

	fmt.Printf("\n\n%s\n", "Slices (iterate through by range)")
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	// i --> index
	// v --> value
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// same result without range
	for i := 0; i < len(pow); i++ {
		fmt.Printf("2**%d = %d\n", i, pow[i])
	}

	fmt.Printf("\n\n%s\n", "Maps")
	var m map[string]Vertex     // declaration of 'm'
	m = make(map[string]Vertex) // definition of 'm'
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
