package infrastructure

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type JsonCharacterRepository struct {
	filepath      string
	characterList *domain.CharacterList
}

const (
	CHARACTER_NOT_FOUND      string = "character \"%s\" not found"
	NO_CHARACTER_CREATED_YET string = "no characters have been found, please create one first"
)

func NewJsonCharacterRepository(filepath string) (*JsonCharacterRepository, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	fileBytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	if len(fileBytes) == 0 {
		characterList := domain.NewEmptyCharacterList()
		return &JsonCharacterRepository{
			filepath:      filepath,
			characterList: &characterList,
		}, nil
	}

	var characters []domain.Character
	if err := json.Unmarshal(fileBytes, &characters); err != nil {
		return nil, err
	}

	characterList := domain.NewFilledCharacterList(characters)
	return &JsonCharacterRepository{
		filepath:      filepath,
		characterList: &characterList,
	}, nil
}

func (jsonCharacterRepository *JsonCharacterRepository) IsCharacterNameUnique(name string) bool {
	return jsonCharacterRepository.characterList.IsCharacterNameUnique(name)
}

func (jsonCharacterRepository *JsonCharacterRepository) AddCharacter(character *domain.Character) {
	jsonCharacterRepository.characterList.AddCharacter(character)
}

func (jsonCharacterRepository *JsonCharacterRepository) SaveCharacterList() error {
	characters := jsonCharacterRepository.characterList.Characters

	jsonBytes, err := json.MarshalIndent(characters, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(jsonCharacterRepository.filepath, jsonBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (jsonCharacterRepository *JsonCharacterRepository) GetAll() *[]domain.Character {
	return &jsonCharacterRepository.characterList.Characters
}

func (jsonCharacterRepository *JsonCharacterRepository) GetByName(name string) (*domain.Character, error) {
	if jsonCharacterRepository.characterList == nil {
		err := errors.New(NO_CHARACTER_CREATED_YET)
		return nil, err
	}

	for i, character := range jsonCharacterRepository.characterList.Characters {
		if strings.EqualFold(character.Name, name) {
			return &jsonCharacterRepository.characterList.Characters[i], nil // Use index to point to actual object, not the temporary copy of the loop
		}
	}

	err := fmt.Errorf(CHARACTER_NOT_FOUND, name)
	return nil, err
}

func (jsonCharacterRepository *JsonCharacterRepository) DeleteCharacter(name string) error {
	return jsonCharacterRepository.characterList.DeleteCharacter(name)
}
