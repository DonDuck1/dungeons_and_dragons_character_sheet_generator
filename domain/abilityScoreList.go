package domain

type AbilityScoreList struct {
	strength     AbilityScore
	dexterity    AbilityScore
	constitution AbilityScore
	intelligence AbilityScore
	wisdom       AbilityScore
	charisma     AbilityScore
}

func NewAbilityScoreList(
	abilityScoreValueList AbilityScoreValueList,
	abilityScoreImprovementList AbilityScoreImprovementList,
) AbilityScoreList {
	strength := NewAbilityScore(STRENGTH, abilityScoreValueList.strengthValue)
	strength.CalculateFinalvalue(abilityScoreImprovementList.strengthImprovements)

	dexterity := NewAbilityScore(DEXTERITY, abilityScoreValueList.dexterityValue)
	dexterity.CalculateFinalvalue(abilityScoreImprovementList.dexterityImprovements)

	constitution := NewAbilityScore(CONSTITUTION, abilityScoreValueList.constitutionValue)
	constitution.CalculateFinalvalue(abilityScoreImprovementList.constitutionImprovements)

	intelligence := NewAbilityScore(INTELLIGENCE, abilityScoreValueList.intelligenceValue)
	intelligence.CalculateFinalvalue(abilityScoreImprovementList.intelligenceImprovements)

	wisdom := NewAbilityScore(WISDOM, abilityScoreValueList.wisdomValue)
	wisdom.CalculateFinalvalue(abilityScoreImprovementList.wisdomImprovements)

	charisma := NewAbilityScore(CHARISMA, abilityScoreValueList.charismaValue)
	charisma.CalculateFinalvalue(abilityScoreImprovementList.charismaImprovements)

	return AbilityScoreList{
		strength:     strength,
		dexterity:    dexterity,
		constitution: constitution,
		intelligence: intelligence,
		wisdom:       wisdom,
		charisma:     charisma,
	}
}
