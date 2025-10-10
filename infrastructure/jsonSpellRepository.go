package infrastructure

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type JsonSpellRepository struct {
	filepath     string
	dndApiSpells []DndApiSpell
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

	var dndApiSpells []DndApiSpell
	if err := json.Unmarshal(fileBytes, &dndApiSpells); err != nil {
		return nil, err
	}

	return &JsonSpellRepository{
		filepath:     filepath,
		dndApiSpells: dndApiSpells,
	}, nil
}

func SaveSpellsAsJson(filepath string, spells *[]DndApiSpell) error {
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

func (jsonSpellRepository JsonSpellRepository) GetCopiesOfAll() *[]DndApiSpell {
	copiedDndApiSpells := jsonSpellRepository.dndApiSpells
	return &copiedDndApiSpells
}

func (jsonSpellRepository JsonSpellRepository) GetCopyByName(name string) (*DndApiSpell, error) {
	if jsonSpellRepository.dndApiSpells == nil {
		err := fmt.Errorf("no spells have been found, please run the init command first")
		return nil, err
	}

	for _, dndApiSpell := range jsonSpellRepository.dndApiSpells {
		if strings.EqualFold(dndApiSpell.Name, name) {
			return &dndApiSpell, nil
		}
	}

	err := fmt.Errorf("could not find spell with name '%s'", name)
	return nil, err
}
