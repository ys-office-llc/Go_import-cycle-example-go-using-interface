package dog

type DogRoar struct{}

func NewDogRoar() *DogRoar {
	return new(DogRoar)
}

func (d *DogRoar) GetDogRoar() string {
	return "Bow Wow !"
}

func (d *DogRoar) GetDogRoar2() string {
	return "Bow Wow 2"
}
