package main

import (
	"fmt"

	"github.com/ys-office-llc/Go_import-cycle-example-go-using-interface/dog"
)

func main() {
	dog := dog.NewDog()
	// dog.GetDog() -> cat.Cat -> DogIF -> DogInterface -> GetDogRoar() -> dog.GetDogRoar() -> "Bow Wow !"
	fmt.Println(dog.GetDog()) // Bow Wow !
}
