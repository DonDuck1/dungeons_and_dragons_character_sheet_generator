package domain

type SubRace struct {
	Name                     string
	AbilityScoreImprovements []AbilityScoreImprovement
	RacialTraits             []RacialTrait
}

func NewSubRace(name string, abilityScoreImprovements []AbilityScoreImprovement, racialTraits []RacialTrait) SubRace {
	return SubRace{Name: name, AbilityScoreImprovements: abilityScoreImprovements, RacialTraits: racialTraits}
}
