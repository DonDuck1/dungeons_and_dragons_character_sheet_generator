package domain

type Race struct {
	name                     string
	abilityScoreImprovements []AbilityScoreImprovement
	numberOfHandSlots        int
}

func NewRace(name string, abilityScoreImprovements []AbilityScoreImprovement, numberOfHandSlots int) Race {
	return Race{name: name, abilityScoreImprovements: abilityScoreImprovements, numberOfHandSlots: numberOfHandSlots}
}
