package main

import (
	"fmt"

	"github.com/PraserX/fullstack-rookie/pkg/list"
	"github.com/PraserX/fullstack-rookie/pkg/user"
	"github.com/kr/pretty"
)

func main() {
	var userParams []user.Option
	userParams = append(userParams, user.OptionUsername("bonj"))
	userParams = append(userParams, user.OptionFirstName("James"))
	userParams = append(userParams, user.OptionLastName("Bond"))

	myUser := user.New(userParams...)
	pretty.Println(myUser)

	passwordHash := myUser.CreateRandomPassword(17)
	pretty.Println(myUser)
	fmt.Println(passwordHash)
	fmt.Println(myUser.GetPasswordHash())
	fmt.Println()
	fmt.Println(myUser.GetBadPasswordHash()) // GetBadPasswordHash works with copy (non pointer receiver)
	fmt.Println(myUser.GetPasswordHash())

	fmt.Println()

	users := list.New()
	users.AddUser(myUser)
	for _, userFullname := range users.GetUsers() {
		fmt.Println(userFullname)
	}
}
