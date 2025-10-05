package infrastructure

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
	"strings"
)

type DndApiWeapon struct {
	Index          string            `json:"index"`
	Name           string            `json:"name"`
	WeaponCategory string            `json:"weapon_category"`
	WeaponRange    DndApiWeaponRange `json:"range"`
	Properties     []DndApiReference `json:"properties"`
}

func (dndApiWeapon DndApiWeapon) AsWeapon() domain.Weapon {
	twoHanded := false

	for _, property := range dndApiWeapon.Properties {
		if strings.EqualFold(property.Name, "Two-Handed") {
			twoHanded = true
		}
	}

	return domain.NewWeapon(
		dndApiWeapon.Name,
		dndApiWeapon.WeaponCategory,
		dndApiWeapon.WeaponRange.Normal,
		twoHanded,
	)
}
