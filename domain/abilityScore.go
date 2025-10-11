package domain

import (
	"math"
)

type AbilityScore struct {
	Name       AbilityScoreName
	BaseValue  int
	FinalValue int
	Modifier   int
}

func NewAbilityScore(name AbilityScoreName, value int) AbilityScore {
	if value < 1 {
		value = 1
	} else if value > 20 {
		value = 20
	}

	return AbilityScore{
		Name:       name,
		BaseValue:  value,
		FinalValue: value,
		Modifier:   int(math.Floor((float64(value) - 10) / 2)),
	}
}

func (abilityScore *AbilityScore) CalculateFinalvalue(abilityScoreImprovements []AbilityScoreImprovement) {
	abilityScore.FinalValue = abilityScore.BaseValue

	for _, abilityScoreImprovement := range abilityScoreImprovements {
		abilityScore.FinalValue += abilityScoreImprovement.Value
	}

	if abilityScore.FinalValue < 1 {
		abilityScore.FinalValue = 1
	} else if abilityScore.FinalValue > 20 {
		abilityScore.FinalValue = 20
	}

	abilityScore.Modifier = int(math.Floor((float64(abilityScore.FinalValue) - 10) / 2))
}
