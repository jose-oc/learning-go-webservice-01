package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// RegisterControllers is responsible for
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	jsonEncoder := json.NewEncoder(w)
	e := jsonEncoder.Encode(data)
	if e != nil {
		fmt.Println("ERROR encoding the response in JSON format")
	}
}
