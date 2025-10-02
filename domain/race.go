package domain

type Race struct {
	Name                     string
	AbilityScoreImprovements []AbilityScoreImprovement
	NumberOfHandSlots        int
}

func NewRace(name string, abilityScoreImprovements []AbilityScoreImprovement, numberOfHandSlots int) Race {
	return Race{Name: name, AbilityScoreImprovements: abilityScoreImprovements, NumberOfHandSlots: numberOfHandSlots}
}
