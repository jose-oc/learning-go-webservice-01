package main

import (
	"fmt"
	"net/http"

	"github.com/jose-oc/learning-go-webservice-01/controllers"
)

func main() {
	fmt.Println("Initializing my webservice")

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
