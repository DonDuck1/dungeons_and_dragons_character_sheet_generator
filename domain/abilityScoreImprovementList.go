package domain

type AbilityScoreImprovementList struct {
	strengthImprovements     []AbilityScoreImprovement
	dexterityImprovements    []AbilityScoreImprovement
	constitutionImprovements []AbilityScoreImprovement
	intelligenceImprovements []AbilityScoreImprovement
	wisdomImprovements       []AbilityScoreImprovement
	charismaImprovements     []AbilityScoreImprovement
}

func getAbilityScoreImprovementsByAbilityScoreName(abilityScoreImprovements []AbilityScoreImprovement, abilityScoreName AbilityScoreName) []AbilityScoreImprovement {
	foundAbilityScoreImprovements := []AbilityScoreImprovement{}

	for _, abilityScoreImprovement := range abilityScoreImprovements {
		if abilityScoreImprovement.GetAbilityScoreName() == abilityScoreName {
			foundAbilityScoreImprovements = append(foundAbilityScoreImprovements, abilityScoreImprovement)
		}
	}

	return foundAbilityScoreImprovements
}

func NewAbilityScoreImprovementList(abilityScoreImprovements []AbilityScoreImprovement) AbilityScoreImprovementList {
	return AbilityScoreImprovementList{
		getAbilityScoreImprovementsByAbilityScoreName(abilityScoreImprovements, Strength),
		getAbilityScoreImprovementsByAbilityScoreName(abilityScoreImprovements, Dexterity),
		getAbilityScoreImprovementsByAbilityScoreName(abilityScoreImprovements, Constitution),
		getAbilityScoreImprovementsByAbilityScoreName(abilityScoreImprovements, Intelligence),
		getAbilityScoreImprovementsByAbilityScoreName(abilityScoreImprovements, Wisdom),
		getAbilityScoreImprovementsByAbilityScoreName(abilityScoreImprovements, Charisma),
	}
}
