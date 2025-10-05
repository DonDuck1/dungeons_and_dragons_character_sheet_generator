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

func NewJsonSpellRepository(filepath string) (*JsonSpellRepository, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	fileBytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	if len(fileBytes) == 0 {
		err := fmt.Errorf("spell JSON has not been initialised, run the 'main.go init' command first")
		return nil, err
	}

	var spells []domain.Spell
	if err := json.Unmarshal(fileBytes, &spells); err != nil {
		return nil, err
	}

	return &JsonSpellRepository{
		filepath: filepath,
		spells:   &spells,
	}, nil
}

func SaveSpellsAsJson(filepath string, spells *[]domain.Spell) error {
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
	if jsonSpellRepository.spells == nil {
		err := fmt.Errorf("no spells have been found, please run the init command first")
		return nil, err
	}

	spells := *jsonSpellRepository.spells
	for i, spell := range spells {
		if strings.EqualFold(spell.Name, name) {
			return &spells[i], nil // Use index to point to actual object, not the temporary copy of the loop
		}
	}

	err := fmt.Errorf("could not find spell with name '%s'", name)
	return nil, err
}
