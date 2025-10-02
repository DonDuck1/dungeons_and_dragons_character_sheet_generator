package domain

type ArmorDexterityModifier struct {
	ArmorMaxDexterityModifier *int
}

func NewArmorDexterityModifier(armorMaxDexterityModifier *int) *ArmorDexterityModifier {
	return &ArmorDexterityModifier{ArmorMaxDexterityModifier: armorMaxDexterityModifier}
}
