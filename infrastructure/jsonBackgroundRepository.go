package infrastructure

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

type JsonBackgroundRepository struct {
	filepath       string
	backgroundList []domain.Background
}

func NewJsonBackgroundRepository(filepath string) (*JsonBackgroundRepository, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	fileBytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	if len(fileBytes) == 0 {
		err := fmt.Errorf("background JSON has not been initialised, run the 'main.go init' command first")
		return nil, err
	}

	var backgroundList []domain.Background
	if err := json.Unmarshal(fileBytes, &backgroundList); err != nil {
		return nil, err
	}

	return &JsonBackgroundRepository{
		filepath:       filepath,
		backgroundList: backgroundList,
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

func (jsonBackgroundRepository JsonBackgroundRepository) GetCopiesOfAll() *[]domain.Background {
	copiedBackgroundList := jsonBackgroundRepository.backgroundList
	return &copiedBackgroundList
}

func (jsonBackgroundRepository JsonBackgroundRepository) GetCopyByName(name string) (*domain.Background, error) {
	if jsonBackgroundRepository.backgroundList == nil {
		err := fmt.Errorf("no backgrounds have been found, please run the init command first")
		return nil, err
	}

	for _, background := range jsonBackgroundRepository.backgroundList {
		if strings.EqualFold(background.Name, name) {
			return &background, nil
		}
	}

	err := fmt.Errorf("could not find background with name '%s'", name)
	return nil, err
}

func (jsonBackgroundRepository JsonBackgroundRepository) GetRandomCopy() (*domain.Background, error) {
	if jsonBackgroundRepository.backgroundList == nil {
		err := fmt.Errorf("no backgrounds have been found, please run the init command first")
		return nil, err
	} else if len(jsonBackgroundRepository.backgroundList) == 0 {
		err := fmt.Errorf("no backgrounds have been found, please re-run the init command")
		return nil, err
	}

	amountOfBackgrounds := len(jsonBackgroundRepository.backgroundList)

	shuffledList := make([]domain.Background, amountOfBackgrounds)
	copy(shuffledList, jsonBackgroundRepository.backgroundList)

	rand.Shuffle(amountOfBackgrounds, func(i, j int) {
		shuffledList[i], shuffledList[j] = shuffledList[j], shuffledList[i]
	})

	return &shuffledList[0], nil
}
