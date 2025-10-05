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

func (race Race) GetChosenAbilityScoreImprovements() []AbilityScoreImprovement {
	abilityScoreImprovements := []AbilityScoreImprovement{}
	abilityScoreImprovements = append(abilityScoreImprovements, race.AbilityScoreImprovements...)
	abilityScoreImprovements = append(abilityScoreImprovements, race.OptionalAbilityScoreImprovementList.ChooseRandomAbilityScoreImprovements()...)
	return abilityScoreImprovements
}
