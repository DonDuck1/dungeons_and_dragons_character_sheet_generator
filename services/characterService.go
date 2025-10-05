package services

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
	"dungeons_and_dragons_character_sheet_generator/infrastructure"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type CharacterService struct {
	jsonBackgroundRepository *infrastructure.JsonBackgroundRepository
	jsonCharacterRepository  *infrastructure.JsonCharacterRepository
	jsonClassRepository      *infrastructure.JsonClassRepository
	jsonRaceRepository       *infrastructure.JsonRaceRepository
}

func NewCharacterService(
	jsonBackgroundRepository *infrastructure.JsonBackgroundRepository,
	jsonCharacterRepository *infrastructure.JsonCharacterRepository,
	jsonClassRepository *infrastructure.JsonClassRepository,
	jsonRaceRepository *infrastructure.JsonRaceRepository,
) *CharacterService {
	return &CharacterService{
		jsonBackgroundRepository: jsonBackgroundRepository,
		jsonCharacterRepository:  jsonCharacterRepository,
		jsonClassRepository:      jsonClassRepository,
		jsonRaceRepository:       jsonRaceRepository,
	}
}

func (characterService CharacterService) CreateNewCharacter(
	characterName string,
	potentialRaceName string,
	potentialMainClassName string,
	level int,
	strengthValue int,
	dexterityValue int,
	constitutionValue int,
	intelligenceValue int,
	wisdomValue int,
	charismaValue int,
) {
	if characterService.jsonBackgroundRepository == nil || characterService.jsonCharacterRepository == nil || characterService.jsonClassRepository == nil || characterService.jsonRaceRepository == nil {
		err := fmt.Errorf("the character service has been created uncorrectly, as a required repository is missing")
		log.Fatal(err)
	}

	if !characterService.jsonCharacterRepository.IsCharacterNameUnique(characterName) {
		err := fmt.Errorf("another character with the same name already exists")
		log.Fatal(err)
	}

	dndApiRaceWithSubraces, err := characterService.jsonRaceRepository.GetByName(potentialRaceName)
	if err != nil {
		log.Fatal(err)
	}

	race, err := dndApiRaceWithSubraces.AsRace(potentialRaceName)
	if err != nil {
		log.Fatal(err)
	}

	dndApiClassWithLevels, err := characterService.jsonClassRepository.GetByName(potentialMainClassName)
	if err != nil {
		log.Fatal(err)
	}

	mainClassTypedName, err := domain.ClassNameFromUntypedPotentialClassName(dndApiClassWithLevels.Name)
	if err != nil {
		log.Fatal(err)
	}

	proficiencyBonus := int(math.Ceil(float64(level)/4)) + 1

	abilityScoreImprovements := race.GetChosenAbilityScoreImprovements()
	abilityScoreImprovementList := domain.NewAbilityScoreImprovementList(abilityScoreImprovements)
	abilityScoreValueList := domain.NewAbilityScoreValueList(strengthValue, dexterityValue, constitutionValue, intelligenceValue, wisdomValue, charismaValue)
	abilityScoreList := domain.NewAbilityScoreList(abilityScoreValueList, abilityScoreImprovementList)

	mainClass := CreateClass(mainClassTypedName, level, proficiencyBonus, abilityScoreList, dndApiClassWithLevels)

	background, err := characterService.jsonBackgroundRepository.GetRandom()
	if err != nil {
		log.Fatal(err)
	}

	skillProficiencies := []domain.SkillProficiencyName{}
	skillProficiencies = append(skillProficiencies, mainClass.SkillProficiencies...)
	skillProficiencies = append(skillProficiencies, background.SkillProficiencies...)
	skillProficiencyList := domain.NewSkillProficiencyList(&abilityScoreList, skillProficiencies, proficiencyBonus)

	inventory := domain.NewEmptyInventory(race.NumberOfHandSlots)

	armorClass := inventory.GetArmorClass(abilityScoreList.Dexterity.Modifier)

	initiative := abilityScoreList.Dexterity.Modifier

	passivePerception := 10 + skillProficiencyList.Perception.Modifier

	character := domain.NewCharacter(
		characterName,
		*race,
		mainClass,
		*background,
		proficiencyBonus,
		abilityScoreList,
		skillProficiencyList,
		armorClass,
		initiative,
		passivePerception,
		inventory,
	)

	characterService.jsonCharacterRepository.AddCharacter(character)
	err = characterService.jsonCharacterRepository.SaveCharacterList()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("saved character %s\n", characterName)
	os.Exit(0)
}

func (characterService CharacterService) ChangeLevelOfCharacter(characterName string, level int) {
	if characterService.jsonCharacterRepository == nil || characterService.jsonClassRepository == nil {
		err := fmt.Errorf("the character service has been created uncorrectly, as a required repository is missing")
		log.Fatal(err)
	}

	character, err := characterService.jsonCharacterRepository.GetByName(characterName)
	if err != nil {
		log.Fatal(err)
	}

	dndApiClassWithLevels, err := characterService.jsonClassRepository.GetByName(string(character.MainClass.Name))
	if err != nil {
		log.Fatal(err)
	}

	proficiencyBonus := int(math.Ceil(float64(level)/4)) + 1
	character.ProficiencyBonus = proficiencyBonus

	EditClass(&character.MainClass, level, proficiencyBonus, &character.AbilityScoreList, dndApiClassWithLevels)

	character.SkillProficiencyList.UpdateSkillProficiencies(proficiencyBonus)

	character.PassivePerception = 10 + character.SkillProficiencyList.Perception.Modifier

	err = characterService.jsonCharacterRepository.SaveCharacterList()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("character succesfully updated to level %d!\n", level)
	os.Exit(0)
}

func (characterService CharacterService) DeleteCharacter(characterName string) {
	if characterService.jsonCharacterRepository == nil {
		err := fmt.Errorf("the character service has been created uncorrectly, as a required repository is missing")
		log.Fatal(err)
	}

	err := characterService.jsonCharacterRepository.DeleteCharacter(characterName)
	if err != nil {
		log.Fatal(err)
	}

	err = characterService.jsonCharacterRepository.SaveCharacterList()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("deleted %s\n", characterName)
	os.Exit(0)
}

func (characterService CharacterService) ViewCharacter(characterName string) {
	if characterService.jsonCharacterRepository == nil {
		err := fmt.Errorf("the character service has been created uncorrectly, as a required repository is missing")
		log.Fatal(err)
	}

	character, err := characterService.jsonCharacterRepository.GetByName(characterName)
	if err != nil {
		log.Fatal(err)
	}

	proficientSkillProficiencies := character.SkillProficiencyList.GetSkillProficienciesThatAreProficient()
	proficientSkillProficiencyNames := []string{}
	for _, proficientSkillProficiency := range *proficientSkillProficiencies {
		for i := 0; i < proficientSkillProficiency.TimesProficiencyIsApplied; i++ {
			proficientSkillProficiencyNames = append(proficientSkillProficiencyNames, strings.ToLower(string(proficientSkillProficiency.Name)))
		}
	}

	fmt.Printf("Name: %s\n", characterName)
	fmt.Printf("Class: %s\n", strings.ToLower(string(character.MainClass.Name)))
	fmt.Printf("Race: %s\n", strings.ToLower(character.Race.Name))
	fmt.Printf("Background: %s\n", strings.ToLower(character.Background.Name))
	fmt.Printf("Level: %d\n", character.MainClass.Level)
	fmt.Println("Ability scores:")
	fmt.Printf("  STR: %d (%+d)\n", character.AbilityScoreList.Strength.Final_value, character.AbilityScoreList.Strength.Modifier)
	fmt.Printf("  DEX: %d (%+d)\n", character.AbilityScoreList.Dexterity.Final_value, character.AbilityScoreList.Dexterity.Modifier)
	fmt.Printf("  CON: %d (%+d)\n", character.AbilityScoreList.Constitution.Final_value, character.AbilityScoreList.Constitution.Modifier)
	fmt.Printf("  INT: %d (%+d)\n", character.AbilityScoreList.Intelligence.Final_value, character.AbilityScoreList.Intelligence.Modifier)
	fmt.Printf("  WIS: %d (%+d)\n", character.AbilityScoreList.Wisdom.Final_value, character.AbilityScoreList.Wisdom.Modifier)
	fmt.Printf("  CHA: %d (%+d)\n", character.AbilityScoreList.Charisma.Final_value, character.AbilityScoreList.Charisma.Modifier)
	fmt.Printf("Proficiency bonus: %+d\n", character.ProficiencyBonus)
	fmt.Printf("Skill proficiencies: %s\n", strings.Join(proficientSkillProficiencyNames, ", "))

	os.Exit(0)
}

func (characterService CharacterService) ListCharacters() {
	if characterService.jsonCharacterRepository == nil {
		err := fmt.Errorf("the character service has been created uncorrectly, as a required repository is missing")
		log.Fatal(err)
	}

	characters := characterService.jsonCharacterRepository.GetAll()
	if len(*characters) <= 0 {
		fmt.Println("there are no characters yet!")
		os.Exit(0)
	}

	fmt.Println("all characters:")
	for _, character := range *characters {
		fmt.Printf("%s, Lv%d %s, %s, %s\n", character.Name, character.MainClass.Level, character.MainClass.Name, character.Race.Name, character.Background.Name)
	}

	os.Exit(0)
}
