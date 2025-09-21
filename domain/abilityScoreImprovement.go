package domain

type AbilityScoreImprovement struct {
	abilityScoreName AbilityScoreName
	value            int
}

func NewAbilityScoreImprovement(name AbilityScoreName, value int) AbilityScoreImprovement {
	return AbilityScoreImprovement{name, value}
}

func (abilityScoreImprovement AbilityScoreImprovement) GetAbilityScoreName() AbilityScoreName {
	return abilityScoreImprovement.abilityScoreName
}

func (abilityScoreImprovement AbilityScoreImprovement) GetValue() int {
	return abilityScoreImprovement.value
}
