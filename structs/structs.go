package main

import (
	"fmt"
	"structs/custom_type"
	"structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	var appUser *user.User
	var err error

	appUser, err = user.New(firstName, lastName, birthDate)
	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@example.com", "abc123")
	fmt.Println(admin.OutputUserDetails())

	// ... do something awesome with that gathered data!

	fmt.Print(appUser.OutputUserDetails())

	var customString custom_type.CustomString = "Test"
	customString.Log()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
