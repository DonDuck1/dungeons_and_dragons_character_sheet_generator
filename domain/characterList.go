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
