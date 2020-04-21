package main

import (
	"net/http"
	"webserver/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

}
