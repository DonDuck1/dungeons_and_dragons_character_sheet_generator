package infrastructure

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type JsonWeaponRepository struct {
	filepath   string
	weaponList []domain.Weapon
}

const (
	UNINITIALISED_WEAPON_JSON string = "weapon JSON has not been initialised, run the 'main.go init' command first"
	NO_WEAPON_WITH_NAME       string = "could not find weapon with name '%s'"
)

func NewJsonWeaponRepository(filepath string) (*JsonWeaponRepository, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	fileBytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	if len(fileBytes) == 0 {
		err := errors.New(UNINITIALISED_WEAPON_JSON)
		return nil, err
	}

	var weaponList []domain.Weapon
	if err := json.Unmarshal(fileBytes, &weaponList); err != nil {
		return nil, err
	}

	return &JsonWeaponRepository{
		filepath:   filepath,
		weaponList: weaponList,
	}, nil
}

func SaveWeaponListAsJson(filepath string, weaponList *[]domain.Weapon) error {
	jsonBytes, err := json.MarshalIndent(weaponList, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath, jsonBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (jsonWeaponRepository JsonWeaponRepository) GetCopiesOfAll() *[]domain.Weapon {
	copiedWeaponList := jsonWeaponRepository.weaponList
	return &copiedWeaponList
}

func (jsonWeaponRepository JsonWeaponRepository) GetCopyByName(name string) (*domain.Weapon, error) {
	if jsonWeaponRepository.weaponList == nil {
		err := errors.New(UNINITIALISED_WEAPON_JSON)
		return nil, err
	}

	for _, weapon := range jsonWeaponRepository.weaponList {
		if strings.EqualFold(weapon.Name, name) {
			return &weapon, nil
		}
	}

	err := fmt.Errorf(NO_WEAPON_WITH_NAME, name)
	return nil, err
}
