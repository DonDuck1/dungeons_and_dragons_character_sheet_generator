package domain

type Race struct {
	Name                                string
	AbilityScoreImprovements            []AbilityScoreImprovement
	OptionalAbilityScoreImprovementList *OptionalAbilityScoreImprovementList
	NumberOfHandSlots                   int
}

func NewRace(name string, abilityScoreImprovements []AbilityScoreImprovement, optionalAbilityScoreImprovementList *OptionalAbilityScoreImprovementList, numberOfHandSlots int) Race {
	return Race{Name: name, AbilityScoreImprovements: abilityScoreImprovements, OptionalAbilityScoreImprovementList: optionalAbilityScoreImprovementList, NumberOfHandSlots: numberOfHandSlots}
}
