package infrastructure

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type JsonSpellRepository struct {
	filepath     string
	dndApiSpells []DndApiSpell
}

const (
	UNINITIALISED_SPELL_JSON string = "spell JSON has not been initialised, run the 'main.go init' command first"
	NO_SPELL_WITH_NAME       string = "could not find spell with name '%s'"
	NO_SPELLS_WITH_CLASS     string = "no spells found for class %s"
)

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
		err := errors.New(UNINITIALISED_SPELL_JSON)
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
		err := errors.New(UNINITIALISED_SPELL_JSON)
		return nil, err
	}

	for _, dndApiSpell := range jsonSpellRepository.dndApiSpells {
		if strings.EqualFold(dndApiSpell.Name, name) {
			return &dndApiSpell, nil
		}
	}

	err := fmt.Errorf(NO_SPELL_WITH_NAME, name)
	return nil, err
}

func (jsonSpellRepository JsonSpellRepository) GetCopiesByClass(className string) (*[]DndApiSpell, error) {
	if jsonSpellRepository.dndApiSpells == nil {
		err := errors.New(UNINITIALISED_SPELL_JSON)
		return nil, err
	}

	var copiedSpellsForClass []DndApiSpell
	for _, dndApiSpell := range jsonSpellRepository.dndApiSpells {
		for _, dndApiSpellValidClass := range dndApiSpell.Classes {
			if strings.EqualFold(dndApiSpellValidClass.Name, className) {
				copiedSpellsForClass = append(copiedSpellsForClass, dndApiSpell)
			}
		}
	}

	if len(copiedSpellsForClass) == 0 {
		err := fmt.Errorf(NO_SPELLS_WITH_CLASS, className)
		return nil, err
	}

	return &copiedSpellsForClass, nil
}
