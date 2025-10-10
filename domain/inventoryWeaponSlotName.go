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

func InventoryWeaponSlotNameFromUntypedPotentialInventoryWeaponSlotName(name string) (InventoryWeaponSlotName, error) {
	switch strings.ToLower(name) {
	case "main hand":
		return MAIN_HAND, nil
	case "off hand":
		return OFF_HAND, nil
	default:
		return OFF_HAND, fmt.Errorf("no weapon hand slot with name '%s' found", name)
	}
}
