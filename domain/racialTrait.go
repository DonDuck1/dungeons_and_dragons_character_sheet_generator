package domain

type RacialTrait struct {
	Name string
}

func NewRacialTrait(name string) RacialTrait {
	return RacialTrait{Name: name}
}
