package services

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
	"dungeons_and_dragons_character_sheet_generator/infrastructure"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type CharacterService struct {
	jsonArmorRepository      *infrastructure.JsonArmorRepository
	jsonBackgroundRepository *infrastructure.JsonBackgroundRepository
	jsonCharacterRepository  *infrastructure.JsonCharacterRepository
	jsonClassRepository      *infrastructure.JsonClassRepository
	jsonRaceRepository       *infrastructure.JsonRaceRepository
	jsonShieldRepository     *infrastructure.JsonShieldRepository
	jsonSpellRepository      *infrastructure.JsonSpellRepository
	jsonWeaponRepository     *infrastructure.JsonWeaponRepository
}

const (
	CHARACTER_SERVICE_REQUIRED_REPOSITORY_MISSING string = "the character service has been created uncorrectly, as a required repository is missing"
	NOT_UNIQUE_CHARACTER_NAME                     string = "another character with the same name already exists"
	WRONGLY_INITIALISED_SPELLCASTER               string = "a %s should be able to spellcast, but somehow wasn't initialised that way"
	UNINITIALISED_KNOWN_SPELLS_LIMIT              string = "a %s should have a limit of known spells, but somehow wasn't initialised that way"
	UNITIALISED_PREPARED_SPELLS_LIMIT             string = "a %s should have a limit of prepared spells, but somehow wasn't initialised that way"
	REACHED_LIMIT_OF_PREPARED_SPELLS              string = "%s has already reached their limit of prepared spells"
	ALREADY_KNOWS_SPELL                           string = "%s already knows spell %s"
	REACHED_LIMIT_OF_CANTRIPS                     string = "%s has already reached their limit of known cantrips"
	REACHED_LIMIT_OF_KNOWN_SPELLS                 string = "%s has already reached their limit of known spells"
	SPELL_LEVEL_TOO_HIGH                          string = "the spell has higher level than the available spell slots"
	SPELL_INVALID_FOR_CLASS                       string = "the spell cannot be used by class %s"
	WRONGLY_INITIALISED_WARLOCK                   string = "a Warlock should be able to spellcast, but somehow wasn't initialised that way"
	PREPARED_CASTER_CAN_NOT_LEARN                 string = "this class prepares spells and can't learn them"
	CAN_NOT_CAST_SPELLS                           string = "this class can't cast spells"
	PREPARED_CASTER_CAN_NOT_FORGET                string = "this class prepares spells and can't forget them"
	CANTRIP_CAN_NOT_BE_PREPARED                   string = "%s is a cantrip and cannot be prepared, only learnt or forgotten"
	LEARNED_CASTER_CAN_NOT_PREPARE                string = "this class learns spells and can't prepare them"
	SPELL_ALREADY_PREPARED                        string = "%s is already prepared"
	SPELL_ALREADY_UNPREPARED                      string = "%s is already unprepared"
)

func NewCharacterService(
	jsonArmorRepository *infrastructure.JsonArmorRepository,
	jsonBackgroundRepository *infrastructure.JsonBackgroundRepository,
	jsonCharacterRepository *infrastructure.JsonCharacterRepository,
	jsonClassRepository *infrastructure.JsonClassRepository,
	jsonRaceRepository *infrastructure.JsonRaceRepository,
	jsonShieldRepository *infrastructure.JsonShieldRepository,
	jsonSpellRepository *infrastructure.JsonSpellRepository,
	jsonWeaponRepository *infrastructure.JsonWeaponRepository,
) *CharacterService {
	return &CharacterService{
		jsonArmorRepository:      jsonArmorRepository,
		jsonBackgroundRepository: jsonBackgroundRepository,
		jsonCharacterRepository:  jsonCharacterRepository,
		jsonClassRepository:      jsonClassRepository,
		jsonRaceRepository:       jsonRaceRepository,
		jsonShieldRepository:     jsonShieldRepository,
		jsonSpellRepository:      jsonSpellRepository,
		jsonWeaponRepository:     jsonWeaponRepository,
	}
}

func (characterService CharacterService) CreateNewCharacter(
	characterName string,
	potentialRaceName string,
	potentialClassName string,
	level int,
	strengthValue int,
	dexterityValue int,
	constitutionValue int,
	intelligenceValue int,
	wisdomValue int,
	charismaValue int,
) {
	if characterService.jsonBackgroundRepository == nil || characterService.jsonCharacterRepository == nil || characterService.jsonClassRepository == nil || characterService.jsonRaceRepository == nil || characterService.jsonSpellRepository == nil {
		err := errors.New(CHARACTER_SERVICE_REQUIRED_REPOSITORY_MISSING)
		log.Fatal(err)
	}

	if !characterService.jsonCharacterRepository.IsCharacterNameUnique(characterName) {
		err := errors.New(NOT_UNIQUE_CHARACTER_NAME)
		log.Fatal(err)
	}

	dndApiRaceWithSubraces, err := characterService.jsonRaceRepository.GetCopyByName(potentialRaceName)
	if err != nil {
		log.Fatal(err)
	}

	race, err := CreateRaceFromDndApiRaceWithSubRaces(potentialRaceName, *dndApiRaceWithSubraces)
	if err != nil {
		log.Fatal(err)
	}

	dndApiClassWithLevels, err := characterService.jsonClassRepository.GetCopyByName(potentialClassName)
	if err != nil {
		log.Fatal(err)
	}

	proficiencyBonus := int(math.Ceil(float64(level)/4)) + 1

	abilityScoreImprovements := race.GetChosenAbilityScoreImprovements()
	abilityScoreImprovementList := domain.NewAbilityScoreImprovementList(abilityScoreImprovements)
	abilityScoreValueList := domain.NewAbilityScoreValueList(strengthValue, dexterityValue, constitutionValue, intelligenceValue, wisdomValue, charismaValue)
	abilityScoreList := domain.NewAbilityScoreList(abilityScoreValueList, abilityScoreImprovementList)

	class := CreateClassFromDndApiClassWithLevels(dndApiClassWithLevels, level, proficiencyBonus, abilityScoreList, characterService.jsonSpellRepository)

	background, err := characterService.jsonBackgroundRepository.GetRandomCopy()
	if err != nil {
		log.Fatal(err)
	}

	skillProficiencies := []domain.SkillProficiencyName{}
	skillProficiencies = append(skillProficiencies, class.SkillProficiencies...)
	skillProficiencies = append(skillProficiencies, background.SkillProficiencies...)
	skillProficiencyList := domain.NewSkillProficiencyList(&abilityScoreList, skillProficiencies, proficiencyBonus)

	inventory := domain.NewEmptyInventory()

	unarmoredArmorClassModifier := 0
	for _, unarmoredArmorClassAbilityScoreModifierName := range class.UnarmoredArmorClassAbilityScoreModifierNameList {
		abilityScore, err := abilityScoreList.GetByName(unarmoredArmorClassAbilityScoreModifierName)
		if err != nil {
			log.Fatal(err)
		}

		unarmoredArmorClassModifier += abilityScore.Modifier
	}
	armorClass := inventory.GetArmorClass(abilityScoreList.Dexterity.Modifier, unarmoredArmorClassModifier)

	initiative := abilityScoreList.Dexterity.Modifier

	passivePerception := 10 + skillProficiencyList.Perception.Modifier

	maxHitPoints := calculateMaxHitPoints(class, abilityScoreList.Constitution.Modifier, *race)

	character := domain.NewCharacter(
		characterName,
		*race,
		class,
		*background,
		proficiencyBonus,
		abilityScoreList,
		skillProficiencyList,
		armorClass,
		initiative,
		passivePerception,
		inventory,
		maxHitPoints,
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
		err := errors.New(CHARACTER_SERVICE_REQUIRED_REPOSITORY_MISSING)
		log.Fatal(err)
	}

	character, err := characterService.jsonCharacterRepository.GetByName(characterName)
	if err != nil {
		log.Fatal(err)
	}

	dndApiClassWithLevels, err := characterService.jsonClassRepository.GetCopyByName(string(character.Class.Name))
	if err != nil {
		log.Fatal(err)
	}

	proficiencyBonus := int(math.Ceil(float64(level)/4)) + 1
	character.ProficiencyBonus = proficiencyBonus

	EditClass(&character.Class, level, proficiencyBonus, &character.AbilityScoreList, dndApiClassWithLevels)

	character.SkillProficiencyList.UpdateSkillProficiencies(proficiencyBonus)

	character.PassivePerception = 10 + character.SkillProficiencyList.Perception.Modifier

	character.MaxHitPoints = calculateMaxHitPoints(character.Class, character.AbilityScoreList.Constitution.Modifier, character.Race)

	err = characterService.jsonCharacterRepository.SaveCharacterList()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("character succesfully updated to level %d!\n", level)
	os.Exit(0)
}

func (characterService CharacterService) DeleteCharacter(characterName string) {
	if characterService.jsonCharacterRepository == nil {
		err := errors.New(CHARACTER_SERVICE_REQUIRED_REPOSITORY_MISSING)
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

func printInventory(inventory domain.Inventory) {
	if inventory.WeaponSlots.MainHand != nil {
		if inventory.WeaponSlots.MainHand.TwoHanded {
			fmt.Printf("Main hand: %s (two-handed)\n", strings.ToLower(inventory.WeaponSlots.MainHand.Name))
		} else {
			fmt.Printf("Main hand: %s\n", strings.ToLower(inventory.WeaponSlots.MainHand.Name))
		}
	}
	if inventory.WeaponSlots.OffHand != nil {
		if inventory.WeaponSlots.OffHand.TwoHanded {
			fmt.Printf("Off hand: %s (two-handed)\n", strings.ToLower(inventory.WeaponSlots.OffHand.Name))
		} else {
			fmt.Printf("Off hand: %s\n", strings.ToLower(inventory.WeaponSlots.OffHand.Name))
		}
	}
	if inventory.Armor != nil {
		fmt.Printf("Armor: %s\n", strings.ToLower(inventory.Armor.Name))
	}
	if inventory.Shield != nil {
		fmt.Printf("Shield: %s\n", strings.ToLower(inventory.Shield.Name))
	}
}

func printClassCastingInfo(class domain.Class) {
	if class.ClassSpellcastingInfo != nil {
		if class.ClassSpellcastingInfo.SpellSlotAmount[0] != 0 || class.ClassSpellcastingInfo.MaxKnownCantrips != 0 {
			fmt.Println("Spell slots:")
			if class.ClassSpellcastingInfo.MaxKnownCantrips != 0 {
				fmt.Printf("  Level 0: %d\n", class.ClassSpellcastingInfo.MaxKnownCantrips)
			}
			for i, spellSlotLevelAmount := range class.ClassSpellcastingInfo.SpellSlotAmount {
				if spellSlotLevelAmount != 0 {
					fmt.Printf("  Level %d: %d\n", i+1, spellSlotLevelAmount)
				}
			}
			fmt.Printf("Spellcasting ability: %s\n", strings.ToLower(string(class.ClassSpellcastingInfo.SpellcastingAbility.Name)))
			fmt.Printf("Spell save DC: %d\n", class.ClassSpellcastingInfo.SpellSaveDC)
			fmt.Printf("Spell attack bonus: %+d\n", class.ClassSpellcastingInfo.SpellAttackBonus)
		}
	}
	if class.ClassWarlockCastingInfo != nil {
		fmt.Println("Spell slots:")
		fmt.Printf("  Level 0: %d\n", class.ClassWarlockCastingInfo.MaxKnownCantrips)
		fmt.Printf("  Level %d: %d\n", class.ClassWarlockCastingInfo.SpellSlotLevel, class.ClassWarlockCastingInfo.SpellSlotAmount)
		fmt.Printf("Spellcasting ability: %s\n", strings.ToLower(string(class.ClassWarlockCastingInfo.SpellcastingAbility.Name)))
		fmt.Printf("Spell save DC: %d\n", class.ClassWarlockCastingInfo.SpellSaveDC)
		fmt.Printf("Spell attack bonus: %+d\n", class.ClassWarlockCastingInfo.SpellAttackBonus)
	}
}

func (characterService CharacterService) ViewCharacter(characterName string) {
	if characterService.jsonCharacterRepository == nil {
		err := errors.New(CHARACTER_SERVICE_REQUIRED_REPOSITORY_MISSING)
		log.Fatal(err)
	}

	character, err := characterService.jsonCharacterRepository.GetByName(characterName)
	if err != nil {
		log.Fatal(err)
	}

	relevantRaceName := character.Race.Name
	if character.Race.SubRace != nil {
		relevantRaceName = character.Race.SubRace.Name
	}

	proficientSkillProficiencies := character.SkillProficiencyList.GetSkillProficienciesThatAreProficient()
	proficientSkillProficiencyNames := []string{}
	for _, proficientSkillProficiency := range *proficientSkillProficiencies {
		for i := 0; i < proficientSkillProficiency.TimesProficiencyIsApplied; i++ {
			proficientSkillProficiencyNames = append(proficientSkillProficiencyNames, strings.ToLower(string(proficientSkillProficiency.Name)))
		}
	}

	fmt.Printf("Name: %s\n", characterName)
	fmt.Printf("Class: %s\n", strings.ToLower(string(character.Class.Name)))
	fmt.Printf("Race: %s\n", strings.ToLower(relevantRaceName))
	fmt.Printf("Background: %s\n", strings.ToLower(character.Background.Name))
	fmt.Printf("Level: %d\n", character.Class.Level)
	fmt.Println("Ability scores:")
	fmt.Printf("  STR: %d (%+d)\n", character.AbilityScoreList.Strength.FinalValue, character.AbilityScoreList.Strength.Modifier)
	fmt.Printf("  DEX: %d (%+d)\n", character.AbilityScoreList.Dexterity.FinalValue, character.AbilityScoreList.Dexterity.Modifier)
	fmt.Printf("  CON: %d (%+d)\n", character.AbilityScoreList.Constitution.FinalValue, character.AbilityScoreList.Constitution.Modifier)
	fmt.Printf("  INT: %d (%+d)\n", character.AbilityScoreList.Intelligence.FinalValue, character.AbilityScoreList.Intelligence.Modifier)
	fmt.Printf("  WIS: %d (%+d)\n", character.AbilityScoreList.Wisdom.FinalValue, character.AbilityScoreList.Wisdom.Modifier)
	fmt.Printf("  CHA: %d (%+d)\n", character.AbilityScoreList.Charisma.FinalValue, character.AbilityScoreList.Charisma.Modifier)
	fmt.Printf("Proficiency bonus: %+d\n", character.ProficiencyBonus)
	fmt.Printf("Skill proficiencies: %s\n", strings.Join(proficientSkillProficiencyNames, ", "))
	printInventory(character.Inventory)
	printClassCastingInfo(character.Class)
	fmt.Printf("Armor class: %d\n", character.ArmorClass)
	fmt.Printf("Initiative bonus: %d\n", character.Initiative)
	fmt.Printf("Passive perception: %d\n", character.PassivePerception)
	// fmt.Printf("Max hit points: %d\n", character.MaxHitPoints) // Can unfortunately not show this, as the CodeGrade will break

	os.Exit(0)
}

func (characterService CharacterService) ListCharacters() {
	if characterService.jsonCharacterRepository == nil {
		err := errors.New(CHARACTER_SERVICE_REQUIRED_REPOSITORY_MISSING)
		log.Fatal(err)
	}

	characters := characterService.jsonCharacterRepository.GetAll()
	if len(*characters) <= 0 {
		fmt.Println("there are no characters yet!")
		os.Exit(0)
	}

	fmt.Println("all characters:")
	for _, character := range *characters {
		if character.Race.SubRace != nil {
			fmt.Printf("%s, Lv%d %s, %s, %s\n", character.Name, character.Class.Level, character.Class.Name, character.Race.SubRace.Name, character.Background.Name)
		} else {
			fmt.Printf("%s, Lv%d %s, %s, %s\n", character.Name, character.Class.Level, character.Class.Name, character.Race.Name, character.Background.Name)
		}
	}

	os.Exit(0)
}

func (characterService CharacterService) EquipWeaponToCharacter(characterName string, weaponName string, potentialInventoryWeaponSlotName string) {
	if characterService.jsonCharacterRepository == nil || characterService.jsonWeaponRepository == nil {
		err := errors.New(CHARACTER_SERVICE_REQUIRED_REPOSITORY_MISSING)
		log.Fatal(err)
	}

	character, err := characterService.jsonCharacterRepository.GetByName(characterName)
	if err != nil {
		log.Fatal(err)
	}

	weapon, err := characterService.jsonWeaponRepository.GetCopyByName(weaponName)
	if err != nil {
		log.Fatal(err)
	}

	inventoryWeaponSlotName, err := domain.InventoryWeaponSlotNameFromUntypedPotentialInventoryWeaponSlotName(potentialInventoryWeaponSlotName)
	if err != nil {
		log.Fatal(err)
	}

	err = character.Inventory.AddWeapon(weapon, inventoryWeaponSlotName)
	if err != nil {
		log.Fatal(err)
	}

	err = characterService.jsonCharacterRepository.SaveCharacterList()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Equipped weapon %s to %s\n", weaponName, potentialInventoryWeaponSlotName)
}

func (characterService CharacterService) EquipArmorToCharacter(characterName string, armorName string) {
	if characterService.jsonCharacterRepository == nil || characterService.jsonArmorRepository == nil {
		err := errors.New(CHARACTER_SERVICE_REQUIRED_REPOSITORY_MISSING)
		log.Fatal(err)
	}

	character, err := characterService.jsonCharacterRepository.GetByName(characterName)
	if err != nil {
		log.Fatal(err)
	}

	armor, err := characterService.jsonArmorRepository.GetCopyByName(armorName)
	if err != nil {
		log.Fatal(err)
	}

	err = character.Inventory.AddArmor(armor)
	if err != nil {
		log.Fatal(err)
	}

	unarmoredArmorClassModifier := 0
	for _, unarmoredArmorClassAbilityScoreModifierName := range character.Class.UnarmoredArmorClassAbilityScoreModifierNameList {
		abilityScore, err := character.AbilityScoreList.GetByName(unarmoredArmorClassAbilityScoreModifierName)
		if err != nil {
			log.Fatal(err)
		}

		unarmoredArmorClassModifier += abilityScore.Modifier
	}
	character.ArmorClass = character.Inventory.GetArmorClass(character.AbilityScoreList.Dexterity.Modifier, unarmoredArmorClassModifier)

	err = characterService.jsonCharacterRepository.SaveCharacterList()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Equipped armor %s\n", armorName)
}

func (characterService CharacterService) EquipShieldToCharacter(characterName string, shieldName string) {
	if characterService.jsonCharacterRepository == nil || characterService.jsonShieldRepository == nil {
		err := errors.New(CHARACTER_SERVICE_REQUIRED_REPOSITORY_MISSING)
		log.Fatal(err)
	}

	character, err := characterService.jsonCharacterRepository.GetByName(characterName)
	if err != nil {
		log.Fatal(err)
	}

	shield, err := characterService.jsonShieldRepository.GetCopyByName(shieldName)
	if err != nil {
		log.Fatal(err)
	}

	err = character.Inventory.AddShield(shield)
	if err != nil {
		log.Fatal(err)
	}

	unarmoredArmorClassModifier := 0
	for _, unarmoredArmorClassAbilityScoreModifierName := range character.Class.UnarmoredArmorClassAbilityScoreModifierNameList {
		abilityScore, err := character.AbilityScoreList.GetByName(unarmoredArmorClassAbilityScoreModifierName)
		if err != nil {
			log.Fatal(err)
		}

		unarmoredArmorClassModifier += abilityScore.Modifier
	}
	character.ArmorClass = character.Inventory.GetArmorClass(character.AbilityScoreList.Dexterity.Modifier, unarmoredArmorClassModifier)

	err = characterService.jsonCharacterRepository.SaveCharacterList()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Equipped shield %s\n", shieldName)
}

func (characterService CharacterService) UnequipWeaponFromCharacter(characterName string, potentialInventoryWeaponSlotName string) {
	if characterService.jsonCharacterRepository == nil || characterService.jsonWeaponRepository == nil {
		err := errors.New(CHARACTER_SERVICE_REQUIRED_REPOSITORY_MISSING)
		log.Fatal(err)
	}

	character, err := characterService.jsonCharacterRepository.GetByName(characterName)
	if err != nil {
		log.Fatal(err)
	}

	inventoryWeaponSlotName, err := domain.InventoryWeaponSlotNameFromUntypedPotentialInventoryWeaponSlotName(potentialInventoryWeaponSlotName)
	if err != nil {
		log.Fatal(err)
	}

	err = character.Inventory.RemoveWeapon(inventoryWeaponSlotName)
	if err != nil {
		log.Fatal(err)
	}

	err = characterService.jsonCharacterRepository.SaveCharacterList()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Unequipped weapon from %s of %s\n", inventoryWeaponSlotName, character.Name)
	os.Exit(0)
}

func (characterService CharacterService) UnequipArmorFromCharacter(characterName string) {
	if characterService.jsonCharacterRepository == nil || characterService.jsonArmorRepository == nil {
		err := errors.New(CHARACTER_SERVICE_REQUIRED_REPOSITORY_MISSING)
		log.Fatal(err)
	}

	character, err := characterService.jsonCharacterRepository.GetByName(characterName)
	if err != nil {
		log.Fatal(err)
	}

	err = character.Inventory.RemoveArmor()
	if err != nil {
		log.Fatal(err)
	}

	unarmoredArmorClassModifier := 0
	for _, unarmoredArmorClassAbilityScoreModifierName := range character.Class.UnarmoredArmorClassAbilityScoreModifierNameList {
		abilityScore, err := character.AbilityScoreList.GetByName(unarmoredArmorClassAbilityScoreModifierName)
		if err != nil {
			log.Fatal(err)
		}

		unarmoredArmorClassModifier += abilityScore.Modifier
	}
	character.ArmorClass = character.Inventory.GetArmorClass(character.AbilityScoreList.Dexterity.Modifier, unarmoredArmorClassModifier)

	err = characterService.jsonCharacterRepository.SaveCharacterList()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Unequipped armor from %s\n", character.Name)
	os.Exit(0)
}

func (characterService CharacterService) UnequipShieldFromCharacter(characterName string) {
	if characterService.jsonCharacterRepository == nil || characterService.jsonShieldRepository == nil {
		err := errors.New(CHARACTER_SERVICE_REQUIRED_REPOSITORY_MISSING)
		log.Fatal(err)
	}

	character, err := characterService.jsonCharacterRepository.GetByName(characterName)
	if err != nil {
		log.Fatal(err)
	}

	err = character.Inventory.RemoveShield()
	if err != nil {
		log.Fatal(err)
	}

	unarmoredArmorClassModifier := 0
	for _, unarmoredArmorClassAbilityScoreModifierName := range character.Class.UnarmoredArmorClassAbilityScoreModifierNameList {
		abilityScore, err := character.AbilityScoreList.GetByName(unarmoredArmorClassAbilityScoreModifierName)
		if err != nil {
			log.Fatal(err)
		}

		unarmoredArmorClassModifier += abilityScore.Modifier
	}
	character.ArmorClass = character.Inventory.GetArmorClass(character.AbilityScoreList.Dexterity.Modifier, unarmoredArmorClassModifier)

	err = characterService.jsonCharacterRepository.SaveCharacterList()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Unequipped shield from %s\n", character.Name)
	os.Exit(0)
}

func makeLearnedSpellcasterLearnSpell(character *domain.Character, spell *domain.Spell, spellName string) {
	if character.Class.ClassSpellcastingInfo == nil {
		err := fmt.Errorf(WRONGLY_INITIALISED_SPELLCASTER, character.Class.Name)
		log.Fatal(err)
	}

	if character.Class.ClassSpellcastingInfo.MaxKnownSpells == nil {
		err := fmt.Errorf(UNINITIALISED_KNOWN_SPELLS_LIMIT, character.Class.Name)
		log.Fatal(err)
	}

	existingSpell, _ := character.Class.ClassSpellcastingInfo.SpellList.GetByName(spellName)
	if existingSpell != nil {
		err := fmt.Errorf(ALREADY_KNOWS_SPELL, character.Name, spellName)
		log.Fatal(err)
	}

	if spell.Level == 0 {
		if character.Class.ClassSpellcastingInfo.MaxKnownCantrips <= character.Class.ClassSpellcastingInfo.SpellList.GetAmountOfKnownCantrips() {
			err := fmt.Errorf(REACHED_LIMIT_OF_CANTRIPS, character.Name)
			log.Fatal(err)
		}
	} else {
		if *character.Class.ClassSpellcastingInfo.MaxKnownSpells <= character.Class.ClassSpellcastingInfo.SpellList.GetAmountOfKnownSpells() {
			err := fmt.Errorf(REACHED_LIMIT_OF_KNOWN_SPELLS, character.Name)
			log.Fatal(err)
		}
	}

	if character.Class.ClassSpellcastingInfo.GetHighestSpellSlotLevel() < spell.Level {
		err := errors.New(SPELL_LEVEL_TOO_HIGH)
		log.Fatal(err)
	}

	if !spell.CanBeUsedByClass(character.Class.Name) {
		err := fmt.Errorf(SPELL_INVALID_FOR_CLASS, character.Class.Name)
		log.Fatal(err)
	}

	character.Class.ClassSpellcastingInfo.SpellList.AddSpell(*spell)
}

func makeWarlockLearnSpell(character *domain.Character, spell *domain.Spell, spellName string) {
	if character.Class.ClassWarlockCastingInfo == nil {
		err := errors.New(WRONGLY_INITIALISED_WARLOCK)
		log.Fatal(err)
	}

	existingSpell, _ := character.Class.ClassWarlockCastingInfo.SpellList.GetByName(spellName)
	if existingSpell != nil {
		err := fmt.Errorf(ALREADY_KNOWS_SPELL, character.Name, spellName)
		log.Fatal(err)
	}

	if spell.Level == 0 {
		if character.Class.ClassWarlockCastingInfo.MaxKnownCantrips <= character.Class.ClassWarlockCastingInfo.SpellList.GetAmountOfKnownCantrips() {
			err := fmt.Errorf(REACHED_LIMIT_OF_CANTRIPS, character.Name)
			log.Fatal(err)
		}
	} else {
		if character.Class.ClassWarlockCastingInfo.MaxKnownSpells <= character.Class.ClassWarlockCastingInfo.SpellList.GetAmountOfKnownSpells() {
			err := fmt.Errorf(REACHED_LIMIT_OF_KNOWN_SPELLS, character.Name)
			log.Fatal(err)
		}
	}

	if character.Class.ClassWarlockCastingInfo.SpellSlotLevel < spell.Level {
		err := errors.New(SPELL_LEVEL_TOO_HIGH)
		log.Fatal(err)
	}

	if !spell.CanBeUsedByClass(character.Class.Name) {
		err := fmt.Errorf(SPELL_INVALID_FOR_CLASS, character.Class.Name)
		log.Fatal(err)
	}

	character.Class.ClassWarlockCastingInfo.SpellList.AddSpell(*spell)
}

func makePreparedSpellcasterLearnSpell(character *domain.Character, spell *domain.Spell, spellName string) {
	if character.Class.ClassSpellcastingInfo == nil {
		err := fmt.Errorf(WRONGLY_INITIALISED_SPELLCASTER, character.Class.Name)
		log.Fatal(err)
	}

	if spell.Level != 0 {
		err := errors.New(PREPARED_CASTER_CAN_NOT_LEARN)
		log.Fatal(err)
	}

	existingSpell, _ := character.Class.ClassSpellcastingInfo.SpellList.GetByName(spellName)
	if existingSpell != nil {
		err := fmt.Errorf(ALREADY_KNOWS_SPELL, character.Name, spellName)
		log.Fatal(err)
	}

	if character.Class.ClassSpellcastingInfo.MaxKnownCantrips <= character.Class.ClassSpellcastingInfo.SpellList.GetAmountOfKnownCantrips() {
		err := fmt.Errorf(REACHED_LIMIT_OF_CANTRIPS, character.Name)
		log.Fatal(err)
	}

	if !spell.CanBeUsedByClass(character.Class.Name) {
		err := fmt.Errorf(SPELL_INVALID_FOR_CLASS, character.Class.Name)
		log.Fatal(err)
	}

	character.Class.ClassSpellcastingInfo.SpellList.AddSpell(*spell)
}

func (characterService CharacterService) MakeCharacterLearnSpell(characterName string, spellName string) {
	if characterService.jsonCharacterRepository == nil || characterService.jsonSpellRepository == nil {
		err := errors.New(CHARACTER_SERVICE_REQUIRED_REPOSITORY_MISSING)
		log.Fatal(err)
	}

	character, err := characterService.jsonCharacterRepository.GetByName(characterName)
	if err != nil {
		log.Fatal(err)
	}

	dndApiSpell, err := characterService.jsonSpellRepository.GetCopyByName(spellName)
	if err != nil {
		log.Fatal(err)
	}

	spell, err := CreateSpellFromDndApiSpell(*dndApiSpell, true)
	if err != nil {
		log.Fatal(err)
	}

	switch character.Class.Name {
	case domain.BARD, domain.RANGER, domain.SORCERER:
		makeLearnedSpellcasterLearnSpell(character, spell, spellName)
	case domain.WARLOCK:
		makeWarlockLearnSpell(character, spell, spellName)
	case domain.CLERIC, domain.DRUID, domain.WIZARD:
		makePreparedSpellcasterLearnSpell(character, spell, spellName)
	case domain.PALADIN:
		err := errors.New(PREPARED_CASTER_CAN_NOT_LEARN)
		log.Fatal(err)
	case domain.BARBARIAN, domain.FIGHTER, domain.MONK, domain.ROGUE:
		err := errors.New(CAN_NOT_CAST_SPELLS)
		log.Fatal(err)
	}

	err = characterService.jsonCharacterRepository.SaveCharacterList()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Learned spell %s\n", spellName)
	os.Exit(0)
}

func makeLearnedSpellcasterForgetSpell(character *domain.Character, spellName string) {
	if character.Class.ClassSpellcastingInfo == nil {
		err := fmt.Errorf(WRONGLY_INITIALISED_SPELLCASTER, character.Class.Name)
		log.Fatal(err)
	}

	character.Class.ClassSpellcastingInfo.SpellList.ForgetSpell(spellName)
}

func makeWarlockForgetSpell(character *domain.Character, spellName string) {
	if character.Class.ClassWarlockCastingInfo == nil {
		err := errors.New(WRONGLY_INITIALISED_WARLOCK)
		log.Fatal(err)
	}

	character.Class.ClassWarlockCastingInfo.SpellList.ForgetSpell(spellName)
}

func makePreparedSpellcasterForgetSpell(character *domain.Character, spellName string, jsonSpellRepository *infrastructure.JsonSpellRepository) {
	dndApiSpell, err := jsonSpellRepository.GetCopyByName(spellName)
	if err != nil {
		log.Fatal(err)
	}

	if dndApiSpell.Level != 0 {
		err := errors.New(PREPARED_CASTER_CAN_NOT_FORGET)
		log.Fatal(err)
	}

	if character.Class.ClassSpellcastingInfo == nil {
		err := fmt.Errorf(WRONGLY_INITIALISED_SPELLCASTER, character.Class.Name)
		log.Fatal(err)
	}

	character.Class.ClassSpellcastingInfo.SpellList.ForgetSpell(spellName)
}

func (characterService CharacterService) MakeCharacterForgetSpell(characterName string, spellName string) {
	if characterService.jsonCharacterRepository == nil {
		err := errors.New(CHARACTER_SERVICE_REQUIRED_REPOSITORY_MISSING)
		log.Fatal(err)
	}

	character, err := characterService.jsonCharacterRepository.GetByName(characterName)
	if err != nil {
		log.Fatal(err)
	}

	switch character.Class.Name {
	case domain.BARD, domain.RANGER, domain.SORCERER:
		makeLearnedSpellcasterForgetSpell(character, spellName)
	case domain.WARLOCK:
		makeWarlockForgetSpell(character, spellName)
	case domain.CLERIC, domain.DRUID, domain.WIZARD:
		makePreparedSpellcasterForgetSpell(character, spellName, characterService.jsonSpellRepository)
	case domain.PALADIN:
		err := errors.New(PREPARED_CASTER_CAN_NOT_FORGET)
		log.Fatal(err)
	case domain.BARBARIAN, domain.FIGHTER, domain.MONK, domain.ROGUE:
		err := errors.New(CAN_NOT_CAST_SPELLS)
		log.Fatal(err)
	}

	err = characterService.jsonCharacterRepository.SaveCharacterList()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Forgot spell %s\n", spellName)
	os.Exit(0)
}

func makePreparedSpellcasterPrepareSpell(character *domain.Character, spellName string) {
	if character.Class.ClassSpellcastingInfo == nil {
		err := fmt.Errorf(WRONGLY_INITIALISED_SPELLCASTER, character.Class.Name)
		log.Fatal(err)
	}

	spell, err := character.Class.ClassSpellcastingInfo.SpellList.GetByName(spellName)
	if err != nil {
		log.Fatal(err)
	}

	if character.Class.ClassSpellcastingInfo.MaxPreparedSpells == nil {
		err := fmt.Errorf(UNITIALISED_PREPARED_SPELLS_LIMIT, character.Class.Name)
		log.Fatal(err)
	}

	if *character.Class.ClassSpellcastingInfo.MaxPreparedSpells <= character.Class.ClassSpellcastingInfo.SpellList.GetAmountOfPreparedSpells() {
		err := fmt.Errorf(REACHED_LIMIT_OF_PREPARED_SPELLS, character.Name)
		log.Fatal(err)
	}

	if character.Class.ClassSpellcastingInfo.GetHighestSpellSlotLevel() < spell.Level {
		err := errors.New(SPELL_LEVEL_TOO_HIGH)
		log.Fatal(err)
	}

	if !spell.CanBeUsedByClass(character.Class.Name) {
		err := fmt.Errorf(SPELL_INVALID_FOR_CLASS, character.Class.Name)
		log.Fatal(err)
	}

	if spell.Prepared {
		err := fmt.Errorf(SPELL_ALREADY_PREPARED, spellName)
		log.Fatal(err)
	}

	spell.Prepared = true
}

func (characterService CharacterService) MakeCharacterPrepareSpell(characterName string, spellName string) {
	if characterService.jsonCharacterRepository == nil || characterService.jsonSpellRepository == nil {
		err := errors.New(CHARACTER_SERVICE_REQUIRED_REPOSITORY_MISSING)
		log.Fatal(err)
	}

	character, err := characterService.jsonCharacterRepository.GetByName(characterName)
	if err != nil {
		log.Fatal(err)
	}

	dndApiSpell, err := characterService.jsonSpellRepository.GetCopyByName(spellName)
	if err != nil {
		log.Fatal(err)
	}

	if dndApiSpell.Level == 0 {
		err := fmt.Errorf(CANTRIP_CAN_NOT_BE_PREPARED, spellName)
		log.Fatal(err)
	}

	switch character.Class.Name {
	case domain.CLERIC, domain.DRUID, domain.PALADIN, domain.WIZARD:
		makePreparedSpellcasterPrepareSpell(character, spellName)
	case domain.BARD, domain.RANGER, domain.SORCERER, domain.WARLOCK:
		err := errors.New(LEARNED_CASTER_CAN_NOT_PREPARE)
		log.Fatal(err)
	case domain.BARBARIAN, domain.FIGHTER, domain.MONK, domain.ROGUE:
		err := errors.New(CAN_NOT_CAST_SPELLS)
		log.Fatal(err)
	}

	err = characterService.jsonCharacterRepository.SaveCharacterList()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Prepared spell %s\n", spellName)
	os.Exit(0)
}

func makePreparedSpellcasterUnprepareSpell(class *domain.Class, spellName string) {
	if class.ClassSpellcastingInfo == nil {
		err := fmt.Errorf(WRONGLY_INITIALISED_SPELLCASTER, class.Name)
		log.Fatal(err)
	}

	spell, err := class.ClassSpellcastingInfo.SpellList.GetByName(spellName)
	if err != nil {
		log.Fatal(err)
	}

	if !spell.Prepared {
		err := fmt.Errorf(SPELL_ALREADY_UNPREPARED, spellName)
		log.Fatal(err)
	}

	spell.Prepared = false
}

func (characterService CharacterService) MakeCharacterUnprepareSpell(characterName string, spellName string) {
	if characterService.jsonCharacterRepository == nil || characterService.jsonSpellRepository == nil {
		err := errors.New(CHARACTER_SERVICE_REQUIRED_REPOSITORY_MISSING)
		log.Fatal(err)
	}

	character, err := characterService.jsonCharacterRepository.GetByName(characterName)
	if err != nil {
		log.Fatal(err)
	}

	dndApiSpell, err := characterService.jsonSpellRepository.GetCopyByName(spellName)
	if err != nil {
		log.Fatal(err)
	}

	if dndApiSpell.Level == 0 {
		err := fmt.Errorf(CANTRIP_CAN_NOT_BE_PREPARED, spellName)
		log.Fatal(err)
	}

	switch character.Class.Name {
	case domain.CLERIC, domain.DRUID, domain.PALADIN, domain.WIZARD:
		makePreparedSpellcasterUnprepareSpell(&character.Class, spellName)
	case domain.BARD, domain.RANGER, domain.SORCERER, domain.WARLOCK:
		err := errors.New(LEARNED_CASTER_CAN_NOT_PREPARE)
		log.Fatal(err)
	case domain.BARBARIAN, domain.FIGHTER, domain.MONK, domain.ROGUE:
		err := errors.New(CAN_NOT_CAST_SPELLS)
		log.Fatal(err)
	}

	err = characterService.jsonCharacterRepository.SaveCharacterList()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Unprepared spell %s\n", spellName)
	os.Exit(0)
}

func calculateMaxHitPoints(class domain.Class, constitutionModifier int, race domain.Race) int {
	maxHitPoints := class.GetStartingMaxHitPointsFromClass(constitutionModifier) + race.GetStartingMaxHitPointsFromRace(class.Level)
	maxHitPointsPerLevel := max(class.GetMaxHitPointsPerLevelFromClass(constitutionModifier)+race.GetMaxHitPointsPerLevelFromRace(class.Level), 1)

	for i := class.Level - 1; i > 0; i-- {
		maxHitPoints += maxHitPointsPerLevel
	}

	return maxHitPoints
}
