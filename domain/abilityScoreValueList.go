package domain

type AbilityScoreValueList struct {
	strengthValue     int
	dexterityValue    int
	constitutionValue int
	intelligenceValue int
	wisdomValue       int
	charismaValue     int
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
		strengthValue:     strengthValue,
		dexterityValue:    dexterityValue,
		constitutionValue: constitutionValue,
		intelligenceValue: intelligenceValue,
		wisdomValue:       wisdomValue,
		charismaValue:     charismaValue,
	}
}
