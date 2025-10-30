package infrastructure

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type JsonShieldRepository struct {
	filepath   string
	shieldList []domain.Shield
}

const (
	UNINITIALISED_SHIELD_JSON string = "shield JSON has not been initialised, run the 'main.go init' command first"
	NO_SHIELD_WITH_NAME       string = "could not find shield with name '%s'"
)

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
		err := errors.New(UNINITIALISED_SHIELD_JSON)
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
		err := errors.New(UNINITIALISED_SHIELD_JSON)
		return nil, err
	}

	for _, shield := range jsonShieldRepository.shieldList {
		if strings.EqualFold(shield.Name, name) {
			return &shield, nil
		}
	}

	err := fmt.Errorf(NO_SHIELD_WITH_NAME, name)
	return nil, err
}
