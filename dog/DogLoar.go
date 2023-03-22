package dog

type DogRoar struct{}

func NewDogRoar() *DogRoar {
	return new(DogRoar)
}

func (d *DogRoar) GetDogRoar() string {
	return "Bow Wow !"
}
