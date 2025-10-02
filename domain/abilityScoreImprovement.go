package domain

type AbilityScoreImprovement struct {
	AbilityScoreName AbilityScoreName
	Value            int
}

func NewAbilityScoreImprovement(abilityScoreName AbilityScoreName, value int) AbilityScoreImprovement {
	return AbilityScoreImprovement{AbilityScoreName: abilityScoreName, Value: value}
}
