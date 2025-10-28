package domain

import "strings"

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

func (race Race) GetMaxHitPointsFromRace(characterLevel int) int {
	maxHitPoints := 0

	if race.SubRace != nil {
		for _, SubRacialTrait := range race.SubRace.RacialTraits {
			if strings.EqualFold(SubRacialTrait.Name, "Dwarven Toughness") {
				maxHitPoints += characterLevel
			}
		}
	}

	return maxHitPoints
}
