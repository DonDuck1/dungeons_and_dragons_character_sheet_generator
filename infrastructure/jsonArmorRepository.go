package infrastructure

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type JsonArmorRepository struct {
	filepath  string
	armorList []domain.Armor
}

func NewJsonArmorRepository(filepath string) (*JsonArmorRepository, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	fileBytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	if len(fileBytes) == 0 {
		err := fmt.Errorf("armor JSON has not been initialised, run the 'main.go init' command first")
		return nil, err
	}

	var armorList []domain.Armor
	if err := json.Unmarshal(fileBytes, &armorList); err != nil {
		return nil, err
	}

	return &JsonArmorRepository{
		filepath:  filepath,
		armorList: armorList,
	}, nil
}

func SaveArmorListAsJson(filepath string, armorList *[]domain.Armor) error {
	jsonBytes, err := json.MarshalIndent(armorList, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath, jsonBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (jsonArmorRepository JsonArmorRepository) GetCopiesOfAll() *[]domain.Armor {
	deepCopiedArmorList := make([]domain.Armor, len(jsonArmorRepository.armorList))

	for i, armor := range jsonArmorRepository.armorList {
		deepCopiedArmorList[i] = armor.GetDeepCopy()
	}

	return &deepCopiedArmorList
}

func (jsonArmorRepository JsonArmorRepository) GetCopyByName(name string) (*domain.Armor, error) {
	if jsonArmorRepository.armorList == nil {
		err := fmt.Errorf("no armor has been found, please run the init command first")
		return nil, err
	}

	for _, armor := range jsonArmorRepository.armorList {
		if strings.EqualFold(armor.Name, name) {
			deepCopiedArmor := armor.GetDeepCopy()

			return &deepCopiedArmor, nil
		}
	}

	err := fmt.Errorf("could not find armor with name '%s'", name)
	return nil, err
}
