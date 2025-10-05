package domain

type SubRace struct {
	Name                     string
	AbilityScoreImprovements []AbilityScoreImprovement
}

func NewSubRace(name string, abilityScoreImprovements []AbilityScoreImprovement) SubRace {
	return SubRace{Name: name, AbilityScoreImprovements: abilityScoreImprovements}
}
