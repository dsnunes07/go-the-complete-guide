package user

import (
	"errors"
	"fmt"
	"time"
)

// define a struct (lower case symbols are visible only inside package)
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// embed struct - like Admin inherits from User
type Admin struct {
	email    string
	password string
	// anonymous embedding
	User
}

func (user User) OutputUserDetails() string {
	return fmt.Sprintf("firstName: %s, lastName: %s, birthDate: %s\n",
		user.firstName,
		user.lastName,
		user.birthDate)
}

// receivers are not references, but copies. So to modify values, receiver must be a pointer!
func (user *User) ClearUsername() {
	user.firstName = ""
	user.lastName = ""
}

// constructor
func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name and birth date are required")
	}
	return &User{
		firstName,
		lastName,
		birthDate,
		time.Now(),
	}, nil
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "----",
			createdAt: time.Now(),
		},
	}
}
