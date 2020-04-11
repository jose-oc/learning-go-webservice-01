package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jose-oc/learning-go-webservice-01/models"

	"github.com/jose-oc/learning-go-webservice-01/controllers"
)

func main() {
	fmt.Println("Initializing my webservice")
	for i := 0; i < 5; i++ {
		models.AddUser(models.User{
			FirstName: "Jose " + strconv.Itoa(i),
			LastName:  "Ortiz",
		})
	}

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
