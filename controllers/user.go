package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/jose-oc/learning-go-webservice-01/models"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d)/?`),
	}
}

// ServeHTTP is my mocked HTTP server
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is Jose's response from the User's controller\n"))
	switch r.Method {
	case http.MethodGet:
		fmt.Println("JOSE GET")
		if r.URL.Path == "/users/" {
			fmt.Println("JOSE GET all users")
			uc.getAll(w, r)
		} else {
			fmt.Println("JOSE GET specific user with id:")
			//uc.get(w, r)
		}
	case http.MethodPost:
		fmt.Println("JOSE POST")
		//uc.post(w, r)
	case http.MethodPut:
		fmt.Println("JOSE PUT")
		//uc.put(w, r)
	default:
		fmt.Println("non supported method", r.Method)
	}
}

func encodeResponseAsJSON(users []*models.User, w http.ResponseWriter) {
	jsonEncoder := json.NewEncoder(w)
	e := jsonEncoder.Encode(users)
	if e != nil {
		fmt.Println("ERROR encoding the response in JSON format")
	}
}

func (uc userController) getAll(w http.ResponseWriter, r *http.Request) {
	allUsers := models.GetUsers()
	encodeResponseAsJSON(allUsers, w)
}
