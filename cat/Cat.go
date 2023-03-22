package cat

type Cat struct {
	DogIF DogInterface
}

func (c *Cat) GetCat() string {
	return c.DogIF.GetDogRoar()
}

func (c Cat) GetCat2() string {
	return c.DogIF.GetDogRoar2()
}
