package infrastructure

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type JsonShieldRepository struct {
	filepath   string
	shieldList []domain.Shield
}

func NewJsonShieldRepository(filepath string) (*JsonShieldRepository, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	fileBytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	if len(fileBytes) == 0 {
		err := fmt.Errorf("shield JSON has not been initialised, run the 'main.go init' command first")
		return nil, err
	}

	var shieldList []domain.Shield
	if err := json.Unmarshal(fileBytes, &shieldList); err != nil {
		return nil, err
	}

	return &JsonShieldRepository{
		filepath:   filepath,
		shieldList: shieldList,
	}, nil
}

func SaveShieldListAsJson(filepath string, shieldList *[]domain.Shield) error {
	jsonBytes, err := json.MarshalIndent(shieldList, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath, jsonBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (jsonShieldRepository JsonShieldRepository) GetCopiesOfAll() *[]domain.Shield {
	copiedShieldList := jsonShieldRepository.shieldList
	return &copiedShieldList
}

func (jsonShieldRepository JsonShieldRepository) GetCopyByName(name string) (*domain.Shield, error) {
	if jsonShieldRepository.shieldList == nil {
		err := fmt.Errorf("no shields have been found, please run the init command first")
		return nil, err
	}

	for _, shield := range jsonShieldRepository.shieldList {
		if strings.EqualFold(shield.Name, name) {
			return &shield, nil
		}
	}

	err := fmt.Errorf("could not find shield with name '%s'", name)
	return nil, err
}
