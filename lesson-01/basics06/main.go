/**
BASICS 06 - Structs
===================

$ go run .
$ go run main.go
$ go build -o my_program .
$ go build -o my_program main.go
*/

package main

import "github.com/kr/pretty"

type User struct {
	FirstName string
	LastName  string
	Username  string
	password  string // unexported field password!
}

func main() {
	user := User{
		FirstName: "Tom",
		LastName:  "Baker",
		Username:  "bakt",
		password:  "Passw0rd", // <-- don't forget ',' at the end!
	}

	// Useful package pretty
	pretty.Println(user)
}
