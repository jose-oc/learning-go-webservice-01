package models

import (
	"errors"
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
