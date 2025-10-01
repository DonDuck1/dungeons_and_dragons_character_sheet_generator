package domain

import "fmt"

type Inventory struct {
	openHandSlots int
	weapons       []Weapon
	armor         *Armor
	shield        *Shield
}

func NewEmptyInventory(numberOfHandSlots int) Inventory {
	return Inventory{openHandSlots: numberOfHandSlots, weapons: []Weapon{}, armor: nil, shield: nil}
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

func (inventory *Inventory) AddWeapon(weapon Weapon) error {
	requiredOpenHandSlots := weapon.GetNumberOfOccupiedHandSlots()

	if inventory.openHandSlots >= requiredOpenHandSlots {
		inventory.openHandSlots -= requiredOpenHandSlots
		inventory.weapons = append(inventory.weapons, weapon)

		return nil
	}

	missingOpenHandSlots := requiredOpenHandSlots - inventory.openHandSlots
	err := fmt.Errorf("Not enough open hand slots. %d more open hand slots are required.", missingOpenHandSlots)
	return err
}

func (inventory *Inventory) RemoveWeapon(name string) error {
	for index, weapon := range inventory.weapons {
		if weapon.name == name {
			inventory.openHandSlots += weapon.GetNumberOfOccupiedHandSlots()
			inventory.weapons = append(inventory.weapons[:index], inventory.weapons[index+1:]...)
			return nil
		}
	}

	err := fmt.Errorf("No weapon found with name %s", name)
	return err
}

func (inventory *Inventory) AddArmor(armor *Armor) error {
	if inventory.armor == nil {
		inventory.armor = armor
		return nil
	}

	err := fmt.Errorf("Character already has armor ('%s') equipped. Please remove it first.", inventory.armor.name)
	return err
}

func (inventory *Inventory) RemoveArmor() {
	inventory.armor = nil
}

func (inventory *Inventory) AddShield(shield *Shield) error {
	if inventory.shield == nil {
		inventory.openHandSlots -= shield.GetNumberOfOccupiedHandSlots()
		inventory.shield = shield
		return nil
	}

	err := fmt.Errorf("Character already has a shield ('%s') equipped. Please remove it first.", inventory.shield.name)
	return err
}

func (inventory *Inventory) RemoveShield() {
	if !(inventory.shield == nil) {
		inventory.openHandSlots += inventory.shield.GetNumberOfOccupiedHandSlots()
		inventory.shield = nil
	}
}
