package main

import (
	"fmt"

	"football_elimination/models"
)

func main() {
	dog := models.Dog{
		Name:  "Fido",
		Breed: "Collie",
		Age:   4,
	}
	fmt.Printf("Dog:%v\n", dog.Name)

}
