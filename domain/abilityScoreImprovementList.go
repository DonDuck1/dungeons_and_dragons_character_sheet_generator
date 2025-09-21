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
		if abilityScoreImprovement.abilityScoreName == abilityScoreName {
			foundAbilityScoreImprovements = append(foundAbilityScoreImprovements, abilityScoreImprovement)
		}
	}

	return foundAbilityScoreImprovements
}

func NewAbilityScoreImprovementList(abilityScoreImprovements []AbilityScoreImprovement) AbilityScoreImprovementList {
	return AbilityScoreImprovementList{
		strengthImprovements:     getAbilityScoreImprovementsByAbilityScoreName(abilityScoreImprovements, Strength),
		dexterityImprovements:    getAbilityScoreImprovementsByAbilityScoreName(abilityScoreImprovements, Dexterity),
		constitutionImprovements: getAbilityScoreImprovementsByAbilityScoreName(abilityScoreImprovements, Constitution),
		intelligenceImprovements: getAbilityScoreImprovementsByAbilityScoreName(abilityScoreImprovements, Intelligence),
		wisdomImprovements:       getAbilityScoreImprovementsByAbilityScoreName(abilityScoreImprovements, Wisdom),
		charismaImprovements:     getAbilityScoreImprovementsByAbilityScoreName(abilityScoreImprovements, Charisma),
	}
}
