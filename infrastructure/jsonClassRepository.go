package infrastructure

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type JsonClassRepository struct {
	filepath                  string
	dndApiClassWithLevelsList []DndApiClassWithLevels
}

const (
	UNINITIALISED_CLASS_JSON string = "class JSON has not been initialised, run the 'main.go init' command first"
	NO_CLASS_WITH_NAME       string = "could not find class with name '%s'"
)

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
		err := errors.New(UNINITIALISED_CLASS_JSON)
		return nil, err
	}

	var dndApiClassWithLevelsList []DndApiClassWithLevels
	if err := json.Unmarshal(fileBytes, &dndApiClassWithLevelsList); err != nil {
		return nil, err
	}

	return &JsonClassRepository{
		filepath:                  filepath,
		dndApiClassWithLevelsList: dndApiClassWithLevelsList,
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

func (jsonClassRepository JsonClassRepository) GetCopiesOfAll() *[]DndApiClassWithLevels {
	deepCopiedDndApiClassWithLevelsList := make([]DndApiClassWithLevels, len(jsonClassRepository.dndApiClassWithLevelsList))
	for i, dndApiClassWithlevels := range jsonClassRepository.dndApiClassWithLevelsList {
		deepCopiedDndApiClassWithLevelsList[i] = dndApiClassWithlevels.GetDeepCopy()
	}

	return &deepCopiedDndApiClassWithLevelsList
}

func (jsonClassRepository JsonClassRepository) GetCopyByName(name string) (*DndApiClassWithLevels, error) {
	if jsonClassRepository.dndApiClassWithLevelsList == nil {
		err := errors.New(UNINITIALISED_CLASS_JSON)
		return nil, err
	}

	for _, dndApiClassWithLevels := range jsonClassRepository.dndApiClassWithLevelsList {
		if strings.EqualFold(dndApiClassWithLevels.Name, name) {
			deepCopiedDndApiClassWithLevels := dndApiClassWithLevels.GetDeepCopy()

			return &deepCopiedDndApiClassWithLevels, nil
		}
	}

	err := fmt.Errorf(NO_CLASS_WITH_NAME, name)
	return nil, err
}
