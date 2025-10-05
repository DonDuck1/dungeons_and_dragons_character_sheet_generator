package infrastructure

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
	"encoding/json"
	"os"
)

type JsonCharacterRepository struct {
	filepath      string
	characterList *domain.CharacterList
}

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
