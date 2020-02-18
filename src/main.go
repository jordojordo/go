package main

import (
	"fmt"

	"github.com/jordojordo/go/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Jordo",
		LastName:  "leach",
	}
	fmt.Println(u)
}
