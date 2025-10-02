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
		if abilityScoreImprovement.AbilityScoreName == abilityScoreName {
			foundAbilityScoreImprovements = append(foundAbilityScoreImprovements, abilityScoreImprovement)
		}
	}

	return foundAbilityScoreImprovements
}

func NewAbilityScoreImprovementList(abilityScoreImprovements []AbilityScoreImprovement) AbilityScoreImprovementList {
	return AbilityScoreImprovementList{
		strengthImprovements:     getAbilityScoreImprovementsByAbilityScoreName(abilityScoreImprovements, STRENGTH),
		dexterityImprovements:    getAbilityScoreImprovementsByAbilityScoreName(abilityScoreImprovements, DEXTERITY),
		constitutionImprovements: getAbilityScoreImprovementsByAbilityScoreName(abilityScoreImprovements, CONSTITUTION),
		intelligenceImprovements: getAbilityScoreImprovementsByAbilityScoreName(abilityScoreImprovements, INTELLIGENCE),
		wisdomImprovements:       getAbilityScoreImprovementsByAbilityScoreName(abilityScoreImprovements, WISDOM),
		charismaImprovements:     getAbilityScoreImprovementsByAbilityScoreName(abilityScoreImprovements, CHARISMA),
	}
}
