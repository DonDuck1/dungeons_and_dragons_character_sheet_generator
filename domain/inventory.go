package domain

import "fmt"

type Inventory struct {
	OpenHandSlots int
	Weapons       []Weapon
	Armor         *Armor
	Shield        *Shield
}

func NewEmptyInventory(numberOfHandSlots int) Inventory {
	return Inventory{OpenHandSlots: numberOfHandSlots, Weapons: []Weapon{}, Armor: nil, Shield: nil}
}

func (inventory Inventory) GetArmorClass(dexterityModifier int) int {
	armorClass := 0

	if inventory.Armor == nil {
		armorClass = 10 + dexterityModifier
	} else {
		armorClass = inventory.Armor.GetArmorClassModifierOfArmor(dexterityModifier)
	}

	if !(inventory.Shield == nil) {
		armorClass += inventory.Shield.ArmorClassModifier
	}

	return armorClass
}

func (inventory *Inventory) AddWeapon(weapon Weapon) error {
	requiredOpenHandSlots := weapon.GetNumberOfOccupiedHandSlots()

	if inventory.OpenHandSlots >= requiredOpenHandSlots {
		inventory.OpenHandSlots -= requiredOpenHandSlots
		inventory.Weapons = append(inventory.Weapons, weapon)

		return nil
	}

	missingOpenHandSlots := requiredOpenHandSlots - inventory.OpenHandSlots
	err := fmt.Errorf("not enough open hand slots, %d more open hand slots are required", missingOpenHandSlots)
	return err
}

func (inventory *Inventory) RemoveWeapon(name string) error {
	for index, weapon := range inventory.Weapons {
		if weapon.Name == name {
			inventory.OpenHandSlots += weapon.GetNumberOfOccupiedHandSlots()
			inventory.Weapons = append(inventory.Weapons[:index], inventory.Weapons[index+1:]...)
			return nil
		}
	}

	err := fmt.Errorf("no weapon found with name %s", name)
	return err
}

func (inventory *Inventory) AddArmor(armor *Armor) error {
	if inventory.Armor == nil {
		inventory.Armor = armor
		return nil
	}

	err := fmt.Errorf("character already has armor ('%s') equipped, please remove it first", inventory.Armor.Name)
	return err
}

func (inventory *Inventory) RemoveArmor() {
	inventory.Armor = nil
}

func (inventory *Inventory) AddShield(shield *Shield) error {
	if inventory.Shield == nil {
		inventory.OpenHandSlots -= shield.GetNumberOfOccupiedHandSlots()
		inventory.Shield = shield
		return nil
	}

	err := fmt.Errorf("character already has a shield ('%s') equipped, please remove it first", inventory.Shield.Name)
	return err
}

func (inventory *Inventory) RemoveShield() {
	if !(inventory.Shield == nil) {
		inventory.OpenHandSlots += inventory.Shield.GetNumberOfOccupiedHandSlots()
		inventory.Shield = nil
	}
}
