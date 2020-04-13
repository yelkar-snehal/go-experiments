package controllers

import "net/http"

// RegisterControllers function to handle routing
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
