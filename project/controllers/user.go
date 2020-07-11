package controllers

import (
	"net/http"
	"regexp"

	"github.com/pluralsight/webservice/models"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

// print static result only
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the User Controller!"))
}

// retrieve all users from models layer
// and return it
func (uc userController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetUsers(), w)
}

// get specific user by id
func (uc userController) get(id int, w http.ResponseWriter) {
	u, err := models.GetUserByID(id)
	if nil != err {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(u, w)
}

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
