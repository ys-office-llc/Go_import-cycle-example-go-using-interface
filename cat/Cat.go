package cat

type Cat struct {
	CatObj DogInterface
}

func (c *Cat) GetCat() string {
	return c.CatObj.GetDogRoar()
}
