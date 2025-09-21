package domain

type Armor struct {
	name                   string
	baseArmorClass         int
	applyDexterityModifier bool
	maxDexterityModifier   int
}

func NewArmor(name string, baseArmorClass int, applyDexterityModifier bool, maxDexterityModifier int) Armor {
	return Armor{name: name, baseArmorClass: baseArmorClass, applyDexterityModifier: applyDexterityModifier, maxDexterityModifier: maxDexterityModifier}
}

func (armor Armor) CalculateArmorClassModifierOfArmor(dexterityModifier int) int {
	armorClassModifier := armor.baseArmorClass

	if armor.applyDexterityModifier {
		if dexterityModifier <= armor.maxDexterityModifier {
			armorClassModifier += dexterityModifier
		} else {
			armorClassModifier += armor.maxDexterityModifier
		}
	}

	return armorClassModifier
}
