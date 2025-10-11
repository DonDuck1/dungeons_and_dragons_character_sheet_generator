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
	jsonArmorRepository      *infrastructure.JsonArmorRepository
	jsonBackgroundRepository *infrastructure.JsonBackgroundRepository
	jsonCharacterRepository  *infrastructure.JsonCharacterRepository
	jsonClassRepository      *infrastructure.JsonClassRepository
	jsonRaceRepository       *infrastructure.JsonRaceRepository
	jsonShieldRepository     *infrastructure.JsonShieldRepository
	jsonSpellRepository      *infrastructure.JsonSpellRepository
	jsonWeaponRepository     *infrastructure.JsonWeaponRepository
}

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
		err := fmt.Errorf("the character service has been created uncorrectly, as a required repository is missing")
		log.Fatal(err)
	}

	if !characterService.jsonCharacterRepository.IsCharacterNameUnique(characterName) {
		err := fmt.Errorf("another character with the same name already exists")
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

	dndApiClassWithLevels, err := characterService.jsonClassRepository.GetCopyByName(string(character.Class.Name))
	if err != nil {
		log.Fatal(err)
	}

	proficiencyBonus := int(math.Ceil(float64(level)/4)) + 1
	character.ProficiencyBonus = proficiencyBonus

	EditClass(&character.Class, level, proficiencyBonus, &character.AbilityScoreList, dndApiClassWithLevels)

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
	fmt.Printf("  STR: %d (%+d)\n", character.AbilityScoreList.Strength.Final_value, character.AbilityScoreList.Strength.Modifier)
	fmt.Printf("  DEX: %d (%+d)\n", character.AbilityScoreList.Dexterity.Final_value, character.AbilityScoreList.Dexterity.Modifier)
	fmt.Printf("  CON: %d (%+d)\n", character.AbilityScoreList.Constitution.Final_value, character.AbilityScoreList.Constitution.Modifier)
	fmt.Printf("  INT: %d (%+d)\n", character.AbilityScoreList.Intelligence.Final_value, character.AbilityScoreList.Intelligence.Modifier)
	fmt.Printf("  WIS: %d (%+d)\n", character.AbilityScoreList.Wisdom.Final_value, character.AbilityScoreList.Wisdom.Modifier)
	fmt.Printf("  CHA: %d (%+d)\n", character.AbilityScoreList.Charisma.Final_value, character.AbilityScoreList.Charisma.Modifier)
	fmt.Printf("Proficiency bonus: %+d\n", character.ProficiencyBonus)
	fmt.Printf("Skill proficiencies: %s\n", strings.Join(proficientSkillProficiencyNames, ", "))
	if character.Inventory.WeaponSlots.MainHand != nil {
		if character.Inventory.WeaponSlots.MainHand.TwoHanded {
			fmt.Printf("Main hand: %s (two-handed)\n", strings.ToLower(character.Inventory.WeaponSlots.MainHand.Name))
		} else {
			fmt.Printf("Main hand: %s\n", strings.ToLower(character.Inventory.WeaponSlots.MainHand.Name))
		}
	}
	if character.Inventory.WeaponSlots.OffHand != nil {
		if character.Inventory.WeaponSlots.OffHand.TwoHanded {
			fmt.Printf("Off hand: %s (two-handed)\n", strings.ToLower(character.Inventory.WeaponSlots.OffHand.Name))
		} else {
			fmt.Printf("Off hand: %s\n", strings.ToLower(character.Inventory.WeaponSlots.OffHand.Name))
		}
	}
	if character.Inventory.Armor != nil {
		fmt.Printf("Armor: %s\n", strings.ToLower(character.Inventory.Armor.Name))
	}
	if character.Inventory.Shield != nil {
		fmt.Printf("Shield: %s\n", strings.ToLower(character.Inventory.Shield.Name))
	}
	if character.Class.ClassSpellcastingInfo != nil {
		if character.Class.ClassSpellcastingInfo.SpellSlotAmount[0] != 0 || character.Class.ClassSpellcastingInfo.MaxKnownCantrips != 0 {
			fmt.Println("Spell slots:")
		}
		if character.Class.ClassSpellcastingInfo.MaxKnownCantrips != 0 {
			fmt.Printf("  Level 0: %d\n", character.Class.ClassSpellcastingInfo.MaxKnownCantrips)
		}
		for i, spellSlotLevelAmount := range character.Class.ClassSpellcastingInfo.SpellSlotAmount {
			if spellSlotLevelAmount != 0 {
				fmt.Printf("  Level %d: %d\n", i+1, spellSlotLevelAmount)
			}
		}
		fmt.Printf("Spellcasting ability: %s\n", strings.ToLower(string(character.Class.ClassSpellcastingInfo.SpellcastingAbility.Name)))
		fmt.Printf("Spell save DC: %d\n", character.Class.ClassSpellcastingInfo.SpellSaveDC)
		fmt.Printf("Spell attack bonus: %+d\n", character.Class.ClassSpellcastingInfo.SpellAttackBonus)
	} else if character.Class.ClassWarlockCastingInfo != nil {
		fmt.Println("Spell slots:")
		fmt.Printf("  Level 0: %d\n", character.Class.ClassWarlockCastingInfo.MaxKnownCantrips)
		fmt.Printf("  Level %d: %d\n", character.Class.ClassWarlockCastingInfo.SpellSlotLevel, character.Class.ClassWarlockCastingInfo.SpellSlotAmount)
		fmt.Printf("Spellcasting ability: %s\n", strings.ToLower(string(character.Class.ClassWarlockCastingInfo.SpellcastingAbility.Name)))
		fmt.Printf("Spell save DC: %d\n", character.Class.ClassWarlockCastingInfo.SpellSaveDC)
		fmt.Printf("Spell attack bonus: %+d\n", character.Class.ClassWarlockCastingInfo.SpellAttackBonus)
	}
	fmt.Printf("Armor class: %d\n", character.ArmorClass)
	fmt.Printf("Initiative bonus: %d\n", character.Initiative)
	fmt.Printf("Passive perception: %d\n", character.PassivePerception)

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
		err := fmt.Errorf("the character service has been created uncorrectly, as a required repository is missing")
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
		err := fmt.Errorf("the character service has been created uncorrectly, as a required repository is missing")
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
		err := fmt.Errorf("the character service has been created uncorrectly, as a required repository is missing")
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
		err := fmt.Errorf("the character service has been created uncorrectly, as a required repository is missing")
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
		err := fmt.Errorf("the character service has been created uncorrectly, as a required repository is missing")
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
		err := fmt.Errorf("the character service has been created uncorrectly, as a required repository is missing")
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

func (characterService CharacterService) MakeCharacterLearnSpell(characterName string, spellName string) {
	if characterService.jsonCharacterRepository == nil || characterService.jsonSpellRepository == nil {
		err := fmt.Errorf("the character service has been created uncorrectly, as a required repository is missing")
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
		if character.Class.ClassSpellcastingInfo == nil {
			err := fmt.Errorf("a %s should be able to spellcast, but somehow wasn't initialised that way", character.Class.Name)
			log.Fatal(err)
		}

		if character.Class.ClassSpellcastingInfo.MaxKnownSpells == nil {
			err := fmt.Errorf("a %s should have a limit of known spells, but somehow wasn't initialised that way", character.Class.Name)
			log.Fatal(err)
		}

		existingSpell, _ := character.Class.ClassSpellcastingInfo.SpellList.GetByName(spellName)
		if existingSpell != nil {
			err := fmt.Errorf("%s already knows spell %s", character.Name, spellName)
			log.Fatal(err)
		}

		if spell.Level == 0 {
			if character.Class.ClassSpellcastingInfo.MaxKnownCantrips <= character.Class.ClassSpellcastingInfo.SpellList.GetAmountOfKnownCantrips() {
				err := fmt.Errorf("%s has already reached their limit of known cantrips", character.Name)
				log.Fatal(err)
			}
		} else {
			if *character.Class.ClassSpellcastingInfo.MaxKnownSpells <= character.Class.ClassSpellcastingInfo.SpellList.GetAmountOfKnownSpells() {
				err := fmt.Errorf("%s has already reached their limit of known spells", character.Name)
				log.Fatal(err)
			}
		}

		if character.Class.ClassSpellcastingInfo.GetHighestSpellSlotLevel() < spell.Level {
			err := fmt.Errorf("the spell has higher level than the available spell slots")
			log.Fatal(err)
		}

		if !spell.CanBeUsedByClass(character.Class.Name) {
			err := fmt.Errorf("the spell cannot be used by class %s", character.Class.Name)
			log.Fatal(err)
		}

		character.Class.ClassSpellcastingInfo.SpellList.AddSpell(*spell)
	case domain.WARLOCK:
		if character.Class.ClassWarlockCastingInfo == nil {
			err := fmt.Errorf("a Warlock should be able to spellcast, but somehow wasn't initialised that way")
			log.Fatal(err)
		}

		existingSpell, _ := character.Class.ClassWarlockCastingInfo.SpellList.GetByName(spellName)
		if existingSpell != nil {
			err := fmt.Errorf("%s already knows spell %s", character.Name, spellName)
			log.Fatal(err)
		}

		if spell.Level == 0 {
			if character.Class.ClassWarlockCastingInfo.MaxKnownCantrips <= character.Class.ClassWarlockCastingInfo.SpellList.GetAmountOfKnownCantrips() {
				err := fmt.Errorf("%s has already reached their limit of known cantrips", character.Name)
				log.Fatal(err)
			}
		} else {
			if character.Class.ClassWarlockCastingInfo.MaxKnownSpells <= character.Class.ClassWarlockCastingInfo.SpellList.GetAmountOfKnownSpells() {
				err := fmt.Errorf("%s has already reached their limit of known spells", character.Name)
				log.Fatal(err)
			}
		}

		if character.Class.ClassWarlockCastingInfo.SpellSlotLevel < spell.Level {
			err := fmt.Errorf("the spell has higher level than the available spell slots")
			log.Fatal(err)
		}

		if !spell.CanBeUsedByClass(character.Class.Name) {
			err := fmt.Errorf("the spell cannot be used by class %s", character.Class.Name)
			log.Fatal(err)
		}

		character.Class.ClassWarlockCastingInfo.SpellList.AddSpell(*spell)
	case domain.CLERIC, domain.DRUID, domain.WIZARD:
		if character.Class.ClassSpellcastingInfo == nil {
			err := fmt.Errorf("a %s should be able to spellcast, but somehow wasn't initialised that way", character.Class.Name)
			log.Fatal(err)
		}

		if spell.Level != 0 {
			err := fmt.Errorf("this class prepares spells and can't learn them")
			log.Fatal(err)
		}

		existingSpell, _ := character.Class.ClassSpellcastingInfo.SpellList.GetByName(spellName)
		if existingSpell != nil {
			err := fmt.Errorf("%s already knows spell %s", character.Name, spellName)
			log.Fatal(err)
		}

		if character.Class.ClassSpellcastingInfo.MaxKnownCantrips <= character.Class.ClassSpellcastingInfo.SpellList.GetAmountOfKnownCantrips() {
			err := fmt.Errorf("%s has already reached their limit of known cantrips", character.Name)
			log.Fatal(err)
		}

		if !spell.CanBeUsedByClass(character.Class.Name) {
			err := fmt.Errorf("the spell cannot be used by class %s", character.Class.Name)
			log.Fatal(err)
		}

		character.Class.ClassSpellcastingInfo.SpellList.AddSpell(*spell)
	case domain.PALADIN:
		err := fmt.Errorf("this class prepares spells and can't learn them")
		log.Fatal(err)
	case domain.BARBARIAN, domain.FIGHTER, domain.MONK, domain.ROGUE:
		err := fmt.Errorf("this class can't cast spells")
		log.Fatal(err)
	}

	err = characterService.jsonCharacterRepository.SaveCharacterList()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Learned spell %s\n", spellName)
	os.Exit(0)
}

func (characterService CharacterService) MakeCharacterForgetSpell(characterName string, spellName string) {
	if characterService.jsonCharacterRepository == nil {
		err := fmt.Errorf("the character service has been created uncorrectly, as a required repository is missing")
		log.Fatal(err)
	}

	character, err := characterService.jsonCharacterRepository.GetByName(characterName)
	if err != nil {
		log.Fatal(err)
	}

	switch character.Class.Name {
	case domain.BARD, domain.RANGER, domain.SORCERER:
		if character.Class.ClassSpellcastingInfo == nil {
			err := fmt.Errorf("a %s should be able to spellcast, but somehow wasn't initialised that way", character.Class.Name)
			log.Fatal(err)
		}

		character.Class.ClassSpellcastingInfo.SpellList.ForgetSpell(spellName)
	case domain.WARLOCK:
		if character.Class.ClassWarlockCastingInfo == nil {
			err := fmt.Errorf("a Warlock should be able to spellcast, but somehow wasn't initialised that way")
			log.Fatal(err)
		}

		character.Class.ClassWarlockCastingInfo.SpellList.ForgetSpell(spellName)
	case domain.CLERIC, domain.DRUID, domain.WIZARD:
		dndApiSpell, err := characterService.jsonSpellRepository.GetCopyByName(spellName)
		if err != nil {
			log.Fatal(err)
		}

		if dndApiSpell.Level != 0 {
			err := fmt.Errorf("this class prepares spells and can't forget them")
			log.Fatal(err)
		}

		if character.Class.ClassSpellcastingInfo == nil {
			err := fmt.Errorf("a %s should be able to spellcast, but somehow wasn't initialised that way", character.Class.Name)
			log.Fatal(err)
		}

		character.Class.ClassSpellcastingInfo.SpellList.ForgetSpell(spellName)
	case domain.PALADIN:
		err := fmt.Errorf("this class prepares spells and can't forget them")
		log.Fatal(err)
	case domain.BARBARIAN, domain.FIGHTER, domain.MONK, domain.ROGUE:
		err := fmt.Errorf("this class can't cast spells")
		log.Fatal(err)
	}

	err = characterService.jsonCharacterRepository.SaveCharacterList()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Forgot spell %s\n", spellName)
	os.Exit(0)
}

func (characterService CharacterService) MakeCharacterPrepareSpell(characterName string, spellName string) {
	if characterService.jsonCharacterRepository == nil || characterService.jsonSpellRepository == nil {
		err := fmt.Errorf("the character service has been created uncorrectly, as a required repository is missing")
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
		err := fmt.Errorf("%s is a cantrip and cannot be prepared, only learnt or forgotten", spellName)
		log.Fatal(err)
	}

	switch character.Class.Name {
	case domain.CLERIC, domain.DRUID, domain.PALADIN, domain.WIZARD:
		if character.Class.ClassSpellcastingInfo == nil {
			err := fmt.Errorf("a %s should be able to spellcast, but somehow wasn't initialised that way", character.Class.Name)
			log.Fatal(err)
		}

		spell, err := character.Class.ClassSpellcastingInfo.SpellList.GetByName(spellName)
		if err != nil {
			log.Fatal(err)
		}

		if character.Class.ClassSpellcastingInfo.MaxPreparedSpells == nil {
			err := fmt.Errorf("a %s should have a limit of prepared spells, but somehow wasn't initialised that way", character.Class.Name)
			log.Fatal(err)
		}

		if *character.Class.ClassSpellcastingInfo.MaxPreparedSpells <= character.Class.ClassSpellcastingInfo.SpellList.GetAmountOfPreparedSpells() {
			err := fmt.Errorf("%s has already reached their limit of prepared spells", character.Name)
			log.Fatal(err)
		}

		if character.Class.ClassSpellcastingInfo.GetHighestSpellSlotLevel() < spell.Level {
			err := fmt.Errorf("the spell has higher level than the available spell slots")
			log.Fatal(err)
		}

		if !spell.CanBeUsedByClass(character.Class.Name) {
			err := fmt.Errorf("the spell cannot be used by class %s", character.Class.Name)
			log.Fatal(err)
		}

		if spell.Prepared {
			err := fmt.Errorf("%s is already prepared", spellName)
			log.Fatal(err)
		}

		spell.Prepared = true
	case domain.BARD, domain.RANGER, domain.SORCERER, domain.WARLOCK:
		err := fmt.Errorf("this class learns spells and can't prepare them")
		log.Fatal(err)
	case domain.BARBARIAN, domain.FIGHTER, domain.MONK, domain.ROGUE:
		err := fmt.Errorf("this class can't cast spells")
		log.Fatal(err)
	}

	err = characterService.jsonCharacterRepository.SaveCharacterList()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Prepared spell %s\n", spellName)
	os.Exit(0)
}

func (characterService CharacterService) MakeCharacterUnprepareSpell(characterName string, spellName string) {
	if characterService.jsonCharacterRepository == nil || characterService.jsonSpellRepository == nil {
		err := fmt.Errorf("the character service has been created uncorrectly, as a required repository is missing")
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
		err := fmt.Errorf("%s is a cantrip and cannot be prepared, only learnt or forgotten", spellName)
		log.Fatal(err)
	}

	switch character.Class.Name {
	case domain.CLERIC, domain.DRUID, domain.PALADIN, domain.WIZARD:
		if character.Class.ClassSpellcastingInfo == nil {
			err := fmt.Errorf("a %s should be able to spellcast, but somehow wasn't initialised that way", character.Class.Name)
			log.Fatal(err)
		}

		spell, err := character.Class.ClassSpellcastingInfo.SpellList.GetByName(spellName)
		if err != nil {
			log.Fatal(err)
		}

		if !spell.Prepared {
			err := fmt.Errorf("%s is already unprepared", spellName)
			log.Fatal(err)
		}

		spell.Prepared = false
	case domain.BARD, domain.RANGER, domain.SORCERER, domain.WARLOCK:
		err := fmt.Errorf("this class learns spells and can't prepare them")
		log.Fatal(err)
	case domain.BARBARIAN, domain.FIGHTER, domain.MONK, domain.ROGUE:
		err := fmt.Errorf("this class can't cast spells")
		log.Fatal(err)
	}

	err = characterService.jsonCharacterRepository.SaveCharacterList()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Unprepared spell %s\n", spellName)
	os.Exit(0)
}
