package domain

import "math"

type AbilityScore struct {
	name        AbilityScoreName
	base_value  int
	final_value int
	modifier    int
}

func NewAbilityScore(name AbilityScoreName, value int) AbilityScore {
	return AbilityScore{
		name:        name,
		base_value:  value,
		final_value: value,
		modifier:    int(math.Floor((float64(value) - 10) / 2)),
	}
}

func (abilityScore *AbilityScore) CalculateFinalvalue(abilityScoreImprovements []AbilityScoreImprovement) {
	abilityScore.final_value = abilityScore.base_value

	for _, abilityScoreImprovement := range abilityScoreImprovements {
		abilityScore.final_value += abilityScoreImprovement.value
	}

	abilityScore.modifier = int(math.Floor((float64(abilityScore.final_value) - 10) / 2))
}
