package infrastructure

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type JsonSpellRepository struct {
	filepath string
	spells   *[]domain.Spell
}

func NewJsonArmSpellRepository(filepath string) (*JsonSpellRepository, error) {
	_, err := os.Stat(filepath)
	if !(err == nil) {
		return nil, err
	}

	fileBytes, err := os.ReadFile(filepath)
	if !(err == nil) {
		return nil, err
	}

	if len(fileBytes) == 0 {
		err := fmt.Errorf("spell JSON has not been initialised, run the 'main.go init' command first")
		return nil, err
	}

	var spells []domain.Spell
	if err := json.Unmarshal(fileBytes, &spells); !(err == nil) {
		return nil, err
	}

	return &JsonSpellRepository{
		filepath: filepath,
		spells:   &spells,
	}, nil
}

func SaveSpellListAsJson(filepath string, spells *[]domain.Spell) error {
	jsonBytes, err := json.MarshalIndent(spells, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath, jsonBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (jsonSpellRepository JsonSpellRepository) GetAll() *[]domain.Spell {
	return jsonSpellRepository.spells
}

func (jsonSpellRepository JsonSpellRepository) GetByName(name string) (*domain.Spell, error) {
	for _, spell := range *jsonSpellRepository.spells {
		if strings.EqualFold(spell.Name, name) {
			return &spell, nil
		}
	}

	err := fmt.Errorf("could not find spell with name '%s'", name)
	return nil, err
}
