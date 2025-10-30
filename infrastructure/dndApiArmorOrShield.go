package infrastructure

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
	"fmt"
	"strings"
)

type DndApiArmorOrShield struct {
	Index         string           `json:"index"`
	Name          string           `json:"name"`
	ArmorCategory string           `json:"armor_category"`
	ArmorClass    DndApiArmorClass `json:"armor_class"`
}

const (
	SHIELD_CAN_NOT_BE_ARMOR string = "%s cannot be initialised as armor, as it is a shield"
	ARMOR_CAN_NOT_BE_SHIELD string = "%s cannot be initialised as shield, as it is armor"
)

func (dndApiArmorOrShield DndApiArmorOrShield) IsShield() bool {
	return strings.EqualFold(dndApiArmorOrShield.ArmorCategory, "Shield")
}

func (dndApiArmor DndApiArmorOrShield) AsArmor() (*domain.Armor, error) {
	if dndApiArmor.IsShield() {
		err := fmt.Errorf(SHIELD_CAN_NOT_BE_ARMOR, dndApiArmor.Name)
		return nil, err
	}

	if dndApiArmor.ArmorClass.DexBonus {
		armorDexterityModifier := domain.NewArmorDexterityModifier(dndApiArmor.ArmorClass.MaxBonus)

		armor := domain.NewArmor(
			dndApiArmor.Name,
			dndApiArmor.ArmorClass.Base,
			armorDexterityModifier,
		)

		return &armor, nil
	}

	armor := domain.NewArmor(
		dndApiArmor.Name,
		dndApiArmor.ArmorClass.Base,
		nil,
	)

	return &armor, nil
}

func (dndApiShield DndApiArmorOrShield) AsShield() (*domain.Shield, error) {
	if !dndApiShield.IsShield() {
		err := fmt.Errorf(ARMOR_CAN_NOT_BE_SHIELD, dndApiShield.Name)
		return nil, err
	}

	shield := domain.NewShield(
		dndApiShield.Name,
		dndApiShield.ArmorClass.Base,
	)

	return &shield, nil
}
