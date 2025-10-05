package domain

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

func (characterList *CharacterList) IsCharacterNameUnique(name string) bool {
	for _, character := range characterList.Characters {
		if character.Name == name {
			return false
		}
	}
	return true
}

func (characterList *CharacterList) AddCharacter(character *Character) {
	characterList.Characters = append(characterList.Characters, *character)
}
