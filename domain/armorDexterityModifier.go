package domain

type ArmorDexterityModifier struct {
	ArmorMaxDexterityModifier *int
}

func NewArmorDexterityModifier(armorMaxDexterityModifier *int) *ArmorDexterityModifier {
	return &ArmorDexterityModifier{ArmorMaxDexterityModifier: armorMaxDexterityModifier}
}

func (armorDexterityModifier ArmorDexterityModifier) GetDeepCopy() *ArmorDexterityModifier {
	var deepCopiedMaxDexMod *int
	if armorDexterityModifier.ArmorMaxDexterityModifier != nil {
		value := *armorDexterityModifier.ArmorMaxDexterityModifier
		deepCopiedMaxDexMod = &value
	}

	return NewArmorDexterityModifier(deepCopiedMaxDexMod)
}
