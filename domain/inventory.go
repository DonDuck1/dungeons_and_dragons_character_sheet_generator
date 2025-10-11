package domain

import "fmt"

type Inventory struct {
	OpenHandSlots int
	WeaponSlots   InventoryWeaponSlots
	Armor         *Armor
	Shield        *Shield
}

func NewEmptyInventory() Inventory {
	return Inventory{OpenHandSlots: 2, WeaponSlots: NewEmptyInventoryWeaponSlots(), Armor: nil, Shield: nil}
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

func (inventory *Inventory) AddWeapon(weapon *Weapon, inventoryWeaponSlotName InventoryWeaponSlotName) error {
	requiredOpenHandSlots := weapon.GetNumberOfOccupiedHandSlots()

	switch inventoryWeaponSlotName {
	case MAIN_HAND:
		if inventory.WeaponSlots.MainHand != nil {
			err := fmt.Errorf("main hand already occupied")
			return err
		}

		if inventory.OpenHandSlots >= requiredOpenHandSlots {
			inventory.OpenHandSlots -= requiredOpenHandSlots
			inventory.WeaponSlots.MainHand = weapon

			return nil
		}
	case OFF_HAND:
		if inventory.WeaponSlots.OffHand != nil {
			err := fmt.Errorf("off hand already occupied")
			return err
		}

		if inventory.OpenHandSlots >= requiredOpenHandSlots {
			inventory.OpenHandSlots -= requiredOpenHandSlots
			inventory.WeaponSlots.OffHand = weapon

			return nil
		}
	}

	missingOpenHandSlots := requiredOpenHandSlots - inventory.OpenHandSlots
	err := fmt.Errorf("not enough open hand slots, %d more open hand slots are required", missingOpenHandSlots)
	return err
}

func (inventory *Inventory) RemoveWeapon(inventoryWeaponSlotName InventoryWeaponSlotName) error {
	switch inventoryWeaponSlotName {
	case MAIN_HAND:
		if inventory.WeaponSlots.MainHand == nil {
			err := fmt.Errorf("no weapon equipped in main hand")
			return err
		}
		inventory.OpenHandSlots += inventory.WeaponSlots.MainHand.GetNumberOfOccupiedHandSlots()
		inventory.WeaponSlots.MainHand = nil
		return nil
	case OFF_HAND:
		if inventory.WeaponSlots.OffHand == nil {
			err := fmt.Errorf("no weapon equipped in off hand")
			return err
		}
		inventory.OpenHandSlots += inventory.WeaponSlots.OffHand.GetNumberOfOccupiedHandSlots()
		inventory.WeaponSlots.OffHand = nil
		return nil
	}

	err := fmt.Errorf("no inventory slot with name %s found", inventoryWeaponSlotName)
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

func (inventory *Inventory) RemoveArmor() error {
	if inventory.Armor == nil {
		err := fmt.Errorf("no armor has been equipped yet")
		return err
	}

	inventory.Armor = nil
	return nil
}

func (inventory *Inventory) AddShield(shield *Shield) error {
	requiredOpenHandSlots := shield.GetNumberOfOccupiedHandSlots()

	if inventory.OpenHandSlots >= requiredOpenHandSlots {
		if inventory.Shield == nil {
			inventory.OpenHandSlots -= shield.GetNumberOfOccupiedHandSlots()
			inventory.Shield = shield
			return nil
		}

		err := fmt.Errorf("character already has a shield ('%s') equipped, please remove it first", inventory.Shield.Name)
		return err
	}

	missingOpenHandSlots := requiredOpenHandSlots - inventory.OpenHandSlots
	err := fmt.Errorf("not enough open hand slots, %d more open hand slots are required", missingOpenHandSlots)
	return err
}

func (inventory *Inventory) RemoveShield() error {
	if inventory.Shield == nil {
		err := fmt.Errorf("no shield has been equipped yet")
		return err
	}

	inventory.OpenHandSlots += inventory.Shield.GetNumberOfOccupiedHandSlots()
	inventory.Shield = nil
	return nil
}
