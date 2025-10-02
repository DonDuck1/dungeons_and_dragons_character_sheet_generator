package domain

type AbilityScoreList struct {
	Strength     AbilityScore
	Dexterity    AbilityScore
	Constitution AbilityScore
	Intelligence AbilityScore
	Wisdom       AbilityScore
	Charisma     AbilityScore
}

func NewAbilityScoreList(
	abilityScoreValueList AbilityScoreValueList,
	abilityScoreImprovementList AbilityScoreImprovementList,
) AbilityScoreList {
	strength := NewAbilityScore(STRENGTH, abilityScoreValueList.StrengthValue)
	strength.CalculateFinalvalue(abilityScoreImprovementList.strengthImprovements)

	dexterity := NewAbilityScore(DEXTERITY, abilityScoreValueList.DexterityValue)
	dexterity.CalculateFinalvalue(abilityScoreImprovementList.dexterityImprovements)

	constitution := NewAbilityScore(CONSTITUTION, abilityScoreValueList.ConstitutionValue)
	constitution.CalculateFinalvalue(abilityScoreImprovementList.constitutionImprovements)

	intelligence := NewAbilityScore(INTELLIGENCE, abilityScoreValueList.IntelligenceValue)
	intelligence.CalculateFinalvalue(abilityScoreImprovementList.intelligenceImprovements)

	wisdom := NewAbilityScore(WISDOM, abilityScoreValueList.WisdomValue)
	wisdom.CalculateFinalvalue(abilityScoreImprovementList.wisdomImprovements)

	charisma := NewAbilityScore(CHARISMA, abilityScoreValueList.CharismaValue)
	charisma.CalculateFinalvalue(abilityScoreImprovementList.charismaImprovements)

	return AbilityScoreList{
		Strength:     strength,
		Dexterity:    dexterity,
		Constitution: constitution,
		Intelligence: intelligence,
		Wisdom:       wisdom,
		Charisma:     charisma,
	}
}
