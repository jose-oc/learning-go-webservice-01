package controllers

import (
	"net/http"
)

// RegisterControllers is responsible for
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
