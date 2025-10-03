package domain

import (
	"math/rand"
)

type OptionalAbilityScoreImprovementList struct {
	OptionalAbilityScoreImprovements []AbilityScoreImprovement
	AmountOfChoices                  int
}

func NewOptionalAbilityScoreImprovementList(optionalAbilityScoreImprovements []AbilityScoreImprovement, amountOfChoices int) OptionalAbilityScoreImprovementList {
	return OptionalAbilityScoreImprovementList{OptionalAbilityScoreImprovements: optionalAbilityScoreImprovements, AmountOfChoices: amountOfChoices}
}

func (optionalAbilityScoreImprovementList OptionalAbilityScoreImprovementList) ChooseRandomAbilityScoreImprovements() []AbilityScoreImprovement {
	n := optionalAbilityScoreImprovementList.AmountOfChoices
	amountOfOptionalAbilityScoreImprovements := len(optionalAbilityScoreImprovementList.OptionalAbilityScoreImprovements)

	if n > amountOfOptionalAbilityScoreImprovements {
		n = amountOfOptionalAbilityScoreImprovements
	}

	shuffledList := make([]AbilityScoreImprovement, amountOfOptionalAbilityScoreImprovements)
	copy(shuffledList, optionalAbilityScoreImprovementList.OptionalAbilityScoreImprovements)

	rand.Shuffle(amountOfOptionalAbilityScoreImprovements, func(i, j int) {
		shuffledList[i], shuffledList[j] = shuffledList[j], shuffledList[i]
	})

	return shuffledList[:n]
}
