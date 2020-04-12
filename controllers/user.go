package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

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
	w.Header().Add("Content-Type", "application/json")
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w, r)
		case http.MethodPost:
			uc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			delete(w.Header(), "Content-Type")
			w.Write([]byte("An ID has to be passed in the URL (http://localhost:3000/users/{id} where {id} is an integer)"))
			return
		}

		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			delete(w.Header(), "Content-Type")
			w.Write([]byte("The ID is not valid, it must be an integer"))
			return
		}

		switch r.Method {
		case http.MethodGet:
			uc.get(id, w)
		case http.MethodPut:
			uc.put(id, w, r)
		case http.MethodDelete:
			uc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (uc *userController) parseRequest(r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body)
	var u models.User
	err := dec.Decode(&u)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

func (uc *userController) getAll(w http.ResponseWriter, r *http.Request) {
	allUsers := models.GetUsers()
	encodeResponseAsJSON(allUsers, w)
}

func (uc *userController) get(id int, w http.ResponseWriter) {
	theUser, err := models.GetUserByID(id)
	if err != nil {
		fmt.Printf("ERROR getting the user '%v':%v\n", id, err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	encodeResponseAsJSON(theUser, w)
}

func (uc *userController) post(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		fmt.Println("ERROR parsing the request", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error parsing the data"))
		return
	}
	u, err = models.AddUser(u)
	if err != nil {
		fmt.Println("ERROR adding the new user", err)
		w.WriteHeader(http.StatusInternalServerError)
		encodeResponseAsJSON(map[string]string{"error": err.Error()}, w)
	} else {
		fmt.Printf("User added: %+v\n", u)
		encodeResponseAsJSON(u, w)
	}
}

func (uc *userController) put(id int, w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}
	if id != u.ID {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ID of submitted user must match ID in URL"))
		return
	}
	u, err = models.UpdateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encodeResponseAsJSON(map[string]string{"error": err.Error()}, w)
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *userController) delete(id int, w http.ResponseWriter) {
	err := models.RemoveUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
