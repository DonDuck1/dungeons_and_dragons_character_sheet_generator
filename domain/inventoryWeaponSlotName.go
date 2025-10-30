package domain

import (
	"fmt"
	"strings"
)

type InventoryWeaponSlotName string

const (
	MAIN_HAND InventoryWeaponSlotName = "Main hand"
	OFF_HAND  InventoryWeaponSlotName = "Off hand"
)

const (
	NO_WEAPON_SLOT_WITH_NAME string = "no weapon hand slot with name '%s' found"
)

func InventoryWeaponSlotNameFromUntypedPotentialInventoryWeaponSlotName(name string) (InventoryWeaponSlotName, error) {
	switch strings.ToLower(name) {
	case "main hand":
		return MAIN_HAND, nil
	case "off hand":
		return OFF_HAND, nil
	default:
		return OFF_HAND, fmt.Errorf(NO_WEAPON_SLOT_WITH_NAME, name)
	}
}
