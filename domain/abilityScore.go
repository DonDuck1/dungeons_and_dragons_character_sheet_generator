package domain

type AbilityScore struct {
	name        AbilityScoreName
	base_value  int
	final_value int
}

func NewAbilityScore(name AbilityScoreName, value int) AbilityScore {
	return AbilityScore{name, value, value}
}

func (abilityScore *AbilityScore) CalculateFinalvalue(abilityScoreImprovements []AbilityScoreImprovement) {
	abilityScore.final_value = abilityScore.base_value

	for _, abilityScoreImprovement := range abilityScoreImprovements {
		abilityScore.final_value += abilityScoreImprovement.GetValue()
	}
}
