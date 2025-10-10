package domain

type Race struct {
	Name                     string
	AbilityScoreImprovements []AbilityScoreImprovement
	SubRace                  *SubRace
}

func NewRace(name string, abilityScoreImprovements []AbilityScoreImprovement, SubRace *SubRace) Race {
	return Race{Name: name, AbilityScoreImprovements: abilityScoreImprovements, SubRace: SubRace}
}

func (race Race) GetChosenAbilityScoreImprovements() []AbilityScoreImprovement {
	abilityScoreImprovements := []AbilityScoreImprovement{}
	abilityScoreImprovements = append(abilityScoreImprovements, race.AbilityScoreImprovements...)
	if race.SubRace != nil {
		abilityScoreImprovements = append(abilityScoreImprovements, race.SubRace.AbilityScoreImprovements...)
	}

	return abilityScoreImprovements
}
