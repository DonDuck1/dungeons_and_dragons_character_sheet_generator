package infrastructure

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type JsonClassRepository struct {
	filepath                  string
	dndApiClassWithLevelsList *[]DndApiClassWithLevels
}

func NewJsonClassRepository(filepath string) (*JsonClassRepository, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	fileBytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	if len(fileBytes) == 0 {
		err := fmt.Errorf("class JSON has not been initialised, run the 'main.go init' command first")
		return nil, err
	}

	var dndApiClassWithLevelsList []DndApiClassWithLevels
	if err := json.Unmarshal(fileBytes, &dndApiClassWithLevelsList); err != nil {
		return nil, err
	}

	return &JsonClassRepository{
		filepath:                  filepath,
		dndApiClassWithLevelsList: &dndApiClassWithLevelsList,
	}, nil
}

func SaveDndApiClassWithLevelsListAsJson(filepath string, dndApiClassWithLevelsList *[]DndApiClassWithLevels) error {
	jsonBytes, err := json.MarshalIndent(dndApiClassWithLevelsList, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath, jsonBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (jsonClassRepository JsonClassRepository) GetAll() *[]DndApiClassWithLevels {
	return jsonClassRepository.dndApiClassWithLevelsList
}

func (jsonClassRepository JsonClassRepository) GetByName(name string) (*DndApiClassWithLevels, error) {
	if jsonClassRepository.dndApiClassWithLevelsList == nil {
		err := fmt.Errorf("no class has been found, please run the init command first")
		return nil, err
	}

	dndApiClassWithLevelsList := *jsonClassRepository.dndApiClassWithLevelsList
	for i, dndApiClassWithLevels := range dndApiClassWithLevelsList {
		if strings.EqualFold(dndApiClassWithLevels.Name, name) {
			return &dndApiClassWithLevelsList[i], nil // Use index to point to actual object, not the temporary copy of the loop
		}
	}

	err := fmt.Errorf("could not find class with name '%s'", name)
	return nil, err
}
