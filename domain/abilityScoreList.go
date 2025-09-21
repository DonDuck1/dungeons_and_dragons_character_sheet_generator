package domain

type AbilityScoreList struct {
	abilityScores [6]AbilityScore
}

func NewAbilityScoreList(
	abilityScoreValueList AbilityScoreValueList,
	abilityScoreImprovementList AbilityScoreImprovementList,
) AbilityScoreList {
	strength := NewAbilityScore(Strength, abilityScoreValueList.strengthValue)
	strength.CalculateFinalvalue(abilityScoreImprovementList.strengthImprovements)

	dexterity := NewAbilityScore(Dexterity, abilityScoreValueList.dexterityValue)
	dexterity.CalculateFinalvalue(abilityScoreImprovementList.dexterityImprovements)

	constitution := NewAbilityScore(Constitution, abilityScoreValueList.constitutionValue)
	constitution.CalculateFinalvalue(abilityScoreImprovementList.constitutionImprovements)

	intelligence := NewAbilityScore(Intelligence, abilityScoreValueList.intelligenceValue)
	intelligence.CalculateFinalvalue(abilityScoreImprovementList.intelligenceImprovements)

	wisdom := NewAbilityScore(Wisdom, abilityScoreValueList.wisdomValue)
	wisdom.CalculateFinalvalue(abilityScoreImprovementList.wisdomImprovements)

	charisma := NewAbilityScore(Charisma, abilityScoreValueList.charismaValue)
	charisma.CalculateFinalvalue(abilityScoreImprovementList.charismaImprovements)

	return AbilityScoreList{[6]AbilityScore{strength, dexterity, constitution, intelligence, wisdom, charisma}}
}
