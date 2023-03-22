package main

import (
	"fmt"

	"github.com/ys-office-llc/Go_import-cycle-example-go-using-interface/dog"
)

func main() {
	dog := dog.NewDog()
	fmt.Println(dog.GetDog()) // Bow Wow !
}
