package domain

type Inventory struct {
	weapons *[]Weapon
	armor   *Armor
	shield  *Shield
}

func NewEmptyInventory() Inventory {
	return Inventory{weapons: nil, armor: nil, shield: nil}
}

func (inventory Inventory) GetArmorClass(dexterityModifier int) int {
	armorClass := 0

	if inventory.armor == nil {
		armorClass = 10 + dexterityModifier
	} else {
		armorClass = inventory.armor.GetArmorClassModifierOfArmor(dexterityModifier)
	}

	if !(inventory.shield == nil) {
		armorClass += inventory.shield.armorClassModifier
	}

	return armorClass
}
