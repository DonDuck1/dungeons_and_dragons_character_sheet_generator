package infrastructure

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type JsonRaceRepository struct {
	filepath string
	raceList *[]domain.Race
}

func NewJsonRaceRepository(filepath string) (*JsonRaceRepository, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	fileBytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	if len(fileBytes) == 0 {
		err := fmt.Errorf("race JSON has not been initialised, run the 'main.go init' command first")
		return nil, err
	}

	var raceList []domain.Race
	if err := json.Unmarshal(fileBytes, &raceList); err != nil {
		return nil, err
	}

	return &JsonRaceRepository{
		filepath: filepath,
		raceList: &raceList,
	}, nil
}

func SaveRaceListAsJson(filepath string, raceList *[]domain.Race) error {
	jsonBytes, err := json.MarshalIndent(raceList, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath, jsonBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (jsonRaceRepository JsonRaceRepository) GetAll() *[]domain.Race {
	return jsonRaceRepository.raceList
}

func (jsonRaceRepository JsonRaceRepository) GetByName(name string) (*domain.Race, error) {
	if jsonRaceRepository.raceList == nil {
		err := fmt.Errorf("no races have been found, please run the init command first")
		return nil, err
	}

	raceList := *jsonRaceRepository.raceList
	for i, race := range raceList {
		if strings.EqualFold(race.Name, name) {
			return &raceList[i], nil // Use index to point to actual object, not the temporary copy of the loop
		}
	}

	err := fmt.Errorf("could not find race with name '%s'", name)
	return nil, err
}
