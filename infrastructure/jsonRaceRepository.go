package infrastructure

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type JsonRaceRepository struct {
	filepath                   string
	dndApiRaceWithSubRacesList []DndApiRaceWithSubRaces
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
		dndApiRaceWithSubRacesList: dndApiRaceWithSubRacesList,
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

func (jsonRaceRepository JsonRaceRepository) GetCopiesOfAll() *[]DndApiRaceWithSubRaces {
	deepCopiedDndApiRaceWithSubRacesList := make([]DndApiRaceWithSubRaces, len(jsonRaceRepository.dndApiRaceWithSubRacesList))
	for i, dndApiRaceWithSubRaces := range jsonRaceRepository.dndApiRaceWithSubRacesList {
		deepCopiedDndApiRaceWithSubRacesList[i] = dndApiRaceWithSubRaces.GetDeepCopy()
	}

	return &deepCopiedDndApiRaceWithSubRacesList
}

func (jsonRaceRepository JsonRaceRepository) GetCopyByName(name string) (*DndApiRaceWithSubRaces, error) {
	if jsonRaceRepository.dndApiRaceWithSubRacesList == nil {
		err := fmt.Errorf("no races have been found, please run the init command first")
		return nil, err
	}

	for _, race := range jsonRaceRepository.dndApiRaceWithSubRacesList {
		if strings.EqualFold(race.Name, name) {
			deepCopiedRace := race.GetDeepCopy()
			return &deepCopiedRace, nil
		}
		for _, subRace := range race.SubRaceList {
			if strings.EqualFold(subRace.Name, name) {
				deepCopiedRace := race.GetDeepCopy()
				return &deepCopiedRace, nil
			}
		}
	}

	err := fmt.Errorf("could not find race with name '%s'", name)
	return nil, err
}
