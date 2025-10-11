package domain

import (
	"fmt"
	"strings"
)

type CharacterList struct {
	Characters []Character
}

func NewEmptyCharacterList() CharacterList {
	characters := []Character{}
	return CharacterList{Characters: characters}
}

func NewFilledCharacterList(characters []Character) CharacterList {
	return CharacterList{Characters: characters}
}

func (characterList *CharacterList) IsCharacterNameUnique(characterName string) bool {
	for _, character := range characterList.Characters {
		if strings.EqualFold(character.Name, characterName) {
			return false
		}
	}
	return true
}

func (characterList *CharacterList) AddCharacter(character *Character) {
	characterList.Characters = append(characterList.Characters, *character)
}

func (characterList *CharacterList) DeleteCharacter(characterName string) error {
	for i, character := range characterList.Characters {
		if strings.EqualFold(character.Name, characterName) {
			characterList.Characters = append(characterList.Characters[:i], characterList.Characters[i+1:]...)
			return nil
		}
	}

	err := fmt.Errorf("character \"%s\" not found", characterName)
	return err
}
