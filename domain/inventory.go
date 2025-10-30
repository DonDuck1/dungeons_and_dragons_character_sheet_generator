package domain

import (
	"errors"
	"fmt"
)

type Inventory struct {
	OpenHandSlots int
	WeaponSlots   InventoryWeaponSlots
	Armor         *Armor
	Shield        *Shield
}

const (
	MAIN_HAND_OCCUPIED           string = "main hand already occupied"
	OFF_HAND_OCCUPIED            string = "off hand already occupied"
	INSUFFICIENT_OPEN_HAND_SLOTS string = "not enough open hand slots, %d more open hand slots are required"
	NO_WEAPON_IN_MAIN_HAND       string = "no weapon equipped in main hand"
	NO_WEAPON_IN_OFF_HAND        string = "no weapon equipped in off hand"
	NO_INVENTORY_SLOT_WITH_NAME  string = "no inventory slot with name %s found"
	ALREADY_EQUIPPED_ARMOR       string = "character already has armor ('%s') equipped, please remove it first"
	NO_ARMOR_EQUIPPED            string = "no armor has been equipped yet"
	ALREADY_EQUIPPED_SHIELD      string = "character already has a shield ('%s') equipped, please remove it first"
	NO_SHIELD_EQUIPPED           string = "no shield has been equipped yet"
)

func NewEmptyInventory() Inventory {
	return Inventory{OpenHandSlots: 2, WeaponSlots: NewEmptyInventoryWeaponSlots(), Armor: nil, Shield: nil}
}

func (inventory Inventory) GetArmorClass(dexterityModifier int, unarmoredArmorClassModifier int) int {
	armorClass := 0

	if inventory.Armor == nil {
		armorClass = 10 + unarmoredArmorClassModifier
	} else {
		armorClass = inventory.Armor.GetArmorClassModifierOfArmor(dexterityModifier)
	}

	if inventory.Shield != nil {
		armorClass += inventory.Shield.ArmorClassModifier
	}

	return armorClass
}

func (inventory *Inventory) AddWeapon(weapon *Weapon, inventoryWeaponSlotName InventoryWeaponSlotName) error {
	requiredOpenHandSlots := weapon.GetNumberOfOccupiedHandSlots()

	switch inventoryWeaponSlotName {
	case MAIN_HAND:
		if inventory.WeaponSlots.MainHand != nil {
			err := errors.New(MAIN_HAND_OCCUPIED)
			return err
		}

		if inventory.OpenHandSlots >= requiredOpenHandSlots {
			inventory.OpenHandSlots -= requiredOpenHandSlots
			inventory.WeaponSlots.MainHand = weapon

			return nil
		}
	case OFF_HAND:
		if inventory.WeaponSlots.OffHand != nil {
			err := errors.New(OFF_HAND_OCCUPIED)
			return err
		}

		if inventory.OpenHandSlots >= requiredOpenHandSlots {
			inventory.OpenHandSlots -= requiredOpenHandSlots
			inventory.WeaponSlots.OffHand = weapon

			return nil
		}
	}

	missingOpenHandSlots := requiredOpenHandSlots - inventory.OpenHandSlots
	err := fmt.Errorf(INSUFFICIENT_OPEN_HAND_SLOTS, missingOpenHandSlots)
	return err
}

func (inventory *Inventory) RemoveWeapon(inventoryWeaponSlotName InventoryWeaponSlotName) error {
	switch inventoryWeaponSlotName {
	case MAIN_HAND:
		if inventory.WeaponSlots.MainHand == nil {
			err := errors.New(NO_WEAPON_IN_MAIN_HAND)
			return err
		}
		inventory.OpenHandSlots += inventory.WeaponSlots.MainHand.GetNumberOfOccupiedHandSlots()
		inventory.WeaponSlots.MainHand = nil
		return nil
	case OFF_HAND:
		if inventory.WeaponSlots.OffHand == nil {
			err := errors.New(NO_WEAPON_IN_OFF_HAND)
			return err
		}
		inventory.OpenHandSlots += inventory.WeaponSlots.OffHand.GetNumberOfOccupiedHandSlots()
		inventory.WeaponSlots.OffHand = nil
		return nil
	}

	err := fmt.Errorf(NO_INVENTORY_SLOT_WITH_NAME, inventoryWeaponSlotName)
	return err
}

func (inventory *Inventory) AddArmor(armor *Armor) error {
	if inventory.Armor == nil {
		inventory.Armor = armor
		return nil
	}

	err := fmt.Errorf(ALREADY_EQUIPPED_ARMOR, inventory.Armor.Name)
	return err
}

func (inventory *Inventory) RemoveArmor() error {
	if inventory.Armor == nil {
		err := errors.New(NO_ARMOR_EQUIPPED)
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

		err := fmt.Errorf(ALREADY_EQUIPPED_SHIELD, inventory.Shield.Name)
		return err
	}

	missingOpenHandSlots := requiredOpenHandSlots - inventory.OpenHandSlots
	err := fmt.Errorf(INSUFFICIENT_OPEN_HAND_SLOTS, missingOpenHandSlots)
	return err
}

func (inventory *Inventory) RemoveShield() error {
	if inventory.Shield == nil {
		err := errors.New(NO_SHIELD_EQUIPPED)
		return err
	}

	inventory.OpenHandSlots += inventory.Shield.GetNumberOfOccupiedHandSlots()
	inventory.Shield = nil
	return nil
}
