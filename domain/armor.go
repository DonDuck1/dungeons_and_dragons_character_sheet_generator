package domain

type Armor struct {
	name                   string
	baseArmorClass         int
	armorDexterityModifier *ArmorDexterityModifier
}

func NewArmor(name string, baseArmorClass int, armorDexterityModifier *ArmorDexterityModifier) Armor {
	return Armor{name: name, baseArmorClass: baseArmorClass, armorDexterityModifier: armorDexterityModifier}
}

func (armor Armor) CalculateArmorClassModifierOfArmor(dexterityModifier int) int {
	if armor.armorDexterityModifier == nil {
		return armor.baseArmorClass
	}

	if armor.armorDexterityModifier.armorMaxDexterityModifier == nil {
		return armor.baseArmorClass + dexterityModifier
	}

	armorMaxDexterityModifier := *armor.armorDexterityModifier.armorMaxDexterityModifier
	if dexterityModifier <= armorMaxDexterityModifier {
		return armor.baseArmorClass + dexterityModifier
	}

	return armor.baseArmorClass + armorMaxDexterityModifier
}
