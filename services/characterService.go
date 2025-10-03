package services

import "dungeons_and_dragons_character_sheet_generator/infrastructure"

type CharacterService struct {
	dndApiGateway           infrastructure.DndApiGateway
	jsonCharacterRepository infrastructure.JsonCharacterRepository
}

// func NewCharacterService(
// 	equipmentCsvFilepath string,
// 	spellCsvFilepath string,
// 	dndApiUrl string,
// 	characterJsonFilepath string,
// ) CharacterService {
// 	return CharacterService{
// 		equipmentRepository:     equipmentRepository,
// 		spellRepository:         spellRepository,
// 		dndApiGateway:           dndApiGateway,
// 		jsonCharacterRepository: jsonCharacterRepository,
// 	}
// }

// func (characterService CharacterService) NewCharacter() {
// 	characterService.jsonCharacterRepository.NewCharacter()
// }
