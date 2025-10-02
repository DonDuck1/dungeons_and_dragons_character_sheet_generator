package domain

type Armor struct {
	Name                   string
	BaseArmorClass         int
	ArmorDexterityModifier *ArmorDexterityModifier
}

func NewArmor(name string, baseArmorClass int, armorDexterityModifier *ArmorDexterityModifier) Armor {
	return Armor{Name: name, BaseArmorClass: baseArmorClass, ArmorDexterityModifier: armorDexterityModifier}
}

func (armor Armor) GetArmorClassModifierOfArmor(dexterityModifier int) int {
	if armor.ArmorDexterityModifier == nil {
		return armor.BaseArmorClass
	}

	if armor.ArmorDexterityModifier.ArmorMaxDexterityModifier == nil {
		return armor.BaseArmorClass + dexterityModifier
	}

	armorMaxDexterityModifier := *armor.ArmorDexterityModifier.ArmorMaxDexterityModifier
	if dexterityModifier <= armorMaxDexterityModifier {
		return armor.BaseArmorClass + dexterityModifier
	}

	return armor.BaseArmorClass + armorMaxDexterityModifier
}
