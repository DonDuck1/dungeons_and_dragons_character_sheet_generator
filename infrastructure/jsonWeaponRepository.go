package infrastructure

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type JsonWeaponRepository struct {
	filepath   string
	weaponList *[]domain.Weapon
}

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
		err := fmt.Errorf("weapon JSON has not been initialised, run the 'main.go init' command first")
		return nil, err
	}

	var weaponList []domain.Weapon
	if err := json.Unmarshal(fileBytes, &weaponList); err != nil {
		return nil, err
	}

	return &JsonWeaponRepository{
		filepath:   filepath,
		weaponList: &weaponList,
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

func (jsonWeaponRepository JsonWeaponRepository) GetAll() *[]domain.Weapon {
	return jsonWeaponRepository.weaponList
}

func (jsonWeaponRepository JsonWeaponRepository) GetByName(name string) (*domain.Weapon, error) {
	for _, weapon := range *jsonWeaponRepository.weaponList {
		if strings.EqualFold(weapon.Name, name) {
			return &weapon, nil
		}
	}

	err := fmt.Errorf("could not find weapon with name '%s'", name)
	return nil, err
}
