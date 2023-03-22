package dog

import "github.com/ys-office-llc/Go_import-cycle-example-go-using-interface/cat"

type Dog struct{}

func NewDog() *Dog {
	return new(Dog)
}

func (d *Dog) GetDog() string {
	cat := cat.Cat{DogIF: NewDogRoar()}
	// GetCat -> DogInterface -> GetDogRoar
	return cat.GetCat2()
}
