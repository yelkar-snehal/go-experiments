package models

import (
	"errors"
	"fmt"
)

// User structure
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User // slice of user objects
	nextID = 1     // : operator not needed on package level
)

// GetUsers function returns slice of pointers
// pointing ar user objects
func GetUsers() []*User {
	return users
}

// AddUser function appends the newly created object's address
// to the existing users slice
func AddUser(u User) (User, error) {
	if 0 != u.ID {
		return User{}, errors.New("New user must not  include an ID or it must be set to zero")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

// GetUserByID function returns an user if id matches
// otherwise return an empty object with an error message
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if id == u.ID {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User with id %v not found", id)
}

// UpdateUser function updates an user if given ID exists in the slice
// previous addr is replaced with new
// otherwise return an empty object with an error message
func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with id %v not found", u.ID)
}

// RemoveUserByID function removes an user if id matches
// otherwise return an empty object with an error message
func RemoveUserByID(id int) error {
	for i, u := range users {
		if id == u.ID {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User with id %v not found", id)
}
