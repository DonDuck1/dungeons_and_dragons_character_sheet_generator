package infrastructure

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type JsonBackgroundRepository struct {
	filepath       string
	backgroundList *[]domain.Background
}

func NewJsonBackgroundRepository(filepath string) (*JsonBackgroundRepository, error) {
	_, err := os.Stat(filepath)
	if !(err == nil) {
		return nil, err
	}

	fileBytes, err := os.ReadFile(filepath)
	if !(err == nil) {
		return nil, err
	}

	if len(fileBytes) == 0 {
		err := fmt.Errorf("background JSON has not been initialised, run the 'main.go init' command first")
		return nil, err
	}

	var backgroundList []domain.Background
	if err := json.Unmarshal(fileBytes, &backgroundList); !(err == nil) {
		return nil, err
	}

	return &JsonBackgroundRepository{
		filepath:       filepath,
		backgroundList: &backgroundList,
	}, nil
}

func SaveBackgroundListAsJson(filepath string, backgroundList *[]domain.Background) error {
	jsonBytes, err := json.MarshalIndent(backgroundList, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath, jsonBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (jsonBackgroundRepository JsonBackgroundRepository) GetAll() *[]domain.Background {
	return jsonBackgroundRepository.backgroundList
}

func (jsonBackgroundRepository JsonBackgroundRepository) GetByName(name string) (*domain.Background, error) {
	for _, background := range *jsonBackgroundRepository.backgroundList {
		if strings.EqualFold(background.Name, name) {
			return &background, nil
		}
	}

	err := fmt.Errorf("could not find background with name '%s'", name)
	return nil, err
}
