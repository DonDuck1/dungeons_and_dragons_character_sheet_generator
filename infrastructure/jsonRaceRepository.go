package infrastructure

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type JsonRaceRepository struct {
	filepath                   string
	dndApiRaceWithSubRacesList *[]DndApiRaceWithSubRaces
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

	var dndApiRaceWithSubRacesList []DndApiRaceWithSubRaces
	if err := json.Unmarshal(fileBytes, &dndApiRaceWithSubRacesList); err != nil {
		return nil, err
	}

	return &JsonRaceRepository{
		filepath:                   filepath,
		dndApiRaceWithSubRacesList: &dndApiRaceWithSubRacesList,
	}, nil
}

func SaveRaceListAsJson(filepath string, dndApiRaceWithSubRacesList *[]DndApiRaceWithSubRaces) error {
	jsonBytes, err := json.MarshalIndent(dndApiRaceWithSubRacesList, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath, jsonBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (jsonRaceRepository JsonRaceRepository) GetAll() *[]DndApiRaceWithSubRaces {
	return jsonRaceRepository.dndApiRaceWithSubRacesList
}

func (jsonRaceRepository JsonRaceRepository) GetByName(name string) (*DndApiRaceWithSubRaces, error) {
	if jsonRaceRepository.dndApiRaceWithSubRacesList == nil {
		err := fmt.Errorf("no races have been found, please run the init command first")
		return nil, err
	}

	dndApiRaceWithSubRacesList := *jsonRaceRepository.dndApiRaceWithSubRacesList
	for i, race := range dndApiRaceWithSubRacesList {
		if strings.EqualFold(race.Name, name) {
			return &dndApiRaceWithSubRacesList[i], nil // Use index to point to actual object, not the temporary copy of the loop
		}
		for _, subRace := range dndApiRaceWithSubRacesList[i].SubRaceList {
			if strings.EqualFold(subRace.Name, name) {
				return &dndApiRaceWithSubRacesList[i], nil // Use index to point to actual object, not the temporary copy of the loop
			}
		}
	}

	err := fmt.Errorf("could not find race with name '%s'", name)
	return nil, err
}
