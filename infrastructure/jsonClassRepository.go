package infrastructure

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type JsonClassRepository struct {
	filepath  string
	classList *[]domain.Class
}

func NewJsonClassRepository(filepath string) (*JsonClassRepository, error) {
	_, err := os.Stat(filepath)
	if !(err == nil) {
		return nil, err
	}

	fileBytes, err := os.ReadFile(filepath)
	if !(err == nil) {
		return nil, err
	}

	if len(fileBytes) == 0 {
		err := fmt.Errorf("class JSON has not been initialised, run the 'main.go init' command first")
		return nil, err
	}

	var classList []domain.Class
	if err := json.Unmarshal(fileBytes, &classList); !(err == nil) {
		return nil, err
	}

	return &JsonClassRepository{
		filepath:  filepath,
		classList: &classList,
	}, nil
}

func SaveClassListAsJson(filepath string, classList *[]domain.Class) error {
	jsonBytes, err := json.MarshalIndent(classList, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath, jsonBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (jsonClassRepository JsonClassRepository) GetAll() *[]domain.Class {
	return jsonClassRepository.classList
}

func (jsonClassRepository JsonClassRepository) GetByName(name string) (*domain.Class, error) {
	for _, class := range *jsonClassRepository.classList {
		if strings.EqualFold(class.Name, name) {
			return &class, nil
		}
	}

	err := fmt.Errorf("could not find class with name '%s'", name)
	return nil, err
}
