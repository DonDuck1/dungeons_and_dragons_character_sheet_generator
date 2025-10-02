package domain

import (
	"math"
)

type AbilityScore struct {
	Name        AbilityScoreName
	Base_value  int
	Final_value int
	Modifier    int
}

func NewAbilityScore(name AbilityScoreName, value int) AbilityScore {
	if value < 1 {
		value = 1
	} else if value > 20 {
		value = 20
	}

	return AbilityScore{
		Name:        name,
		Base_value:  value,
		Final_value: value,
		Modifier:    int(math.Floor((float64(value) - 10) / 2)),
	}
}

func (abilityScore *AbilityScore) CalculateFinalvalue(abilityScoreImprovements []AbilityScoreImprovement) {
	abilityScore.Final_value = abilityScore.Base_value

	for _, abilityScoreImprovement := range abilityScoreImprovements {
		abilityScore.Final_value += abilityScoreImprovement.Value
	}

	if abilityScore.Final_value < 1 {
		abilityScore.Final_value = 1
	} else if abilityScore.Final_value > 20 {
		abilityScore.Final_value = 20
	}

	abilityScore.Modifier = int(math.Floor((float64(abilityScore.Final_value) - 10) / 2))
}
