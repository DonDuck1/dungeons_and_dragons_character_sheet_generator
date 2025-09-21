package domain

type AbilityScoreImprovement struct {
	abilityScoreName AbilityScoreName
	value            int
}

func NewAbilityScoreImprovement(abilityScoreName AbilityScoreName, value int) AbilityScoreImprovement {
	return AbilityScoreImprovement{abilityScoreName: abilityScoreName, value: value}
}
