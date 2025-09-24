package domain

type ArmorDexterityModifier struct {
	armorMaxDexterityModifier *int
}

func NewArmorDexterityModifier(armorMaxDexterityModifier *int) *ArmorDexterityModifier {
	return &ArmorDexterityModifier{armorMaxDexterityModifier: armorMaxDexterityModifier}
}
