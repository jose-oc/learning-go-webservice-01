package models

import (
	"errors"
	"fmt"
)

// User is the definition of the users of the webservice
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	/*
		Slice of pointers to User.
		We user Pointers because this way we're goint to be able to manipulate each one of these user objects from any point in our application
		without having to copy the user around, so it's going to be more efficient.
		And, important, I don't care for sharing this throughout the application (security concerns).
	*/
	users []*User

	// nextID is the sequence of this simulated database
	nextID = 1
)

// GetUsers returns the list of users in the system
func GetUsers() []*User {
	return users
}

// AddUser appends the user to the DB. Returns an error if the user already exists
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New user must not include an ID, this will be provided in the returning user")
	}

	u.ID = nextID
	nextID++
	for _, userInDB := range users {
		if userInDB.FirstName == u.FirstName && userInDB.LastName == u.LastName {
			return u, errors.New("The user already exists")
		}
	}
	users = append(users, &u)
	return u, nil
}

// GetUserByID returns the User with the ID provided or an error
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

// RemoveUserByID OJO take a look at how to remove an element from a slice
// Create a new slice with the elements we want, Notice the ellipsis at the end (...)
func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found", id)
}

// UpdateUser updates the user replacing the pointer to the new object provided
// Can this be dangerous?
func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if u.ID == candidate.ID {
			users[i] = &u
			return User{}, nil
		}
	}
	return u, fmt.Errorf("User with ID '%v' doesn't exist. Impossible to update", u.ID)
}
