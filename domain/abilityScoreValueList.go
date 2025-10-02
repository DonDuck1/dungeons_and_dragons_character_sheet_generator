package domain

type AbilityScoreValueList struct {
	StrengthValue     int
	DexterityValue    int
	ConstitutionValue int
	IntelligenceValue int
	WisdomValue       int
	CharismaValue     int
}

func NewAbilityScoreValueList(
	strengthValue int,
	dexterityValue int,
	constitutionValue int,
	intelligenceValue int,
	wisdomValue int,
	charismaValue int,
) AbilityScoreValueList {
	return AbilityScoreValueList{
		StrengthValue:     strengthValue,
		DexterityValue:    dexterityValue,
		ConstitutionValue: constitutionValue,
		IntelligenceValue: intelligenceValue,
		WisdomValue:       wisdomValue,
		CharismaValue:     charismaValue,
	}
}
