package main

import (
	"net/http"

	"github.com/jordojordo/go/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
