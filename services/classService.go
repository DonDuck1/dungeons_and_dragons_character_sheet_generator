package services

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
	"dungeons_and_dragons_character_sheet_generator/infrastructure"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

type ClassService struct {
	dndApiGateway *infrastructure.DndApiGateway
}

const (
	NIL_DND_API_CLASS_WITH_LEVELS              string = "dndApiClassWithLevels provided is a nil value"
	INITIALISED_SPELL_LIST_REQUIRED            string = "a %s should have an initialised spell list"
	INITIALISED_SPELL_LIST_REQUIRED_WARLOCK    string = "a warlock should have an initialised spell list"
	UNKNOWN_CLASS_WITH_SPELLCASTING            string = "unknown class (with spellcasting) detected, character creation cannot continue"
	SPELLCASTING_NOT_DEFINED_PER_LEVEL         string = "according to the API, the class has spellcasting, but the spellcasting is not defined per level; character creation cannot continue"
	NON_SPELLCASTER_HAS_SPELLCASTING_PER_LEVEL string = "according to the API, the class does not have spellcasting, but the (non-existent) spellcasting is somehow actually defined per level; character creation cannot continue"
)

func NewClassService(dndApiGateway *infrastructure.DndApiGateway) *ClassService {
	return &ClassService{dndApiGateway: dndApiGateway}
}

func getSpellSlotAmountFromDndApiClassLevelSpellcasting(dndApiClassLevelSpellcasting *infrastructure.DndApiClassLevelSpellcasting) [9]int {
	spellSlotAmount := [9]int{
		dndApiClassLevelSpellcasting.SpellSlotsLevel1,
		dndApiClassLevelSpellcasting.SpellSlotsLevel2,
		dndApiClassLevelSpellcasting.SpellSlotsLevel3,
		dndApiClassLevelSpellcasting.SpellSlotsLevel4,
		dndApiClassLevelSpellcasting.SpellSlotsLevel5,
		0,
		0,
		0,
		0,
	}
	if dndApiClassLevelSpellcasting.SpellSlotsLevel6 != nil {
		spellSlotAmount[5] = *dndApiClassLevelSpellcasting.SpellSlotsLevel6
	}
	if dndApiClassLevelSpellcasting.SpellSlotsLevel7 != nil {
		spellSlotAmount[6] = *dndApiClassLevelSpellcasting.SpellSlotsLevel7
	}
	if dndApiClassLevelSpellcasting.SpellSlotsLevel8 != nil {
		spellSlotAmount[7] = *dndApiClassLevelSpellcasting.SpellSlotsLevel8
	}
	if dndApiClassLevelSpellcasting.SpellSlotsLevel9 != nil {
		spellSlotAmount[8] = *dndApiClassLevelSpellcasting.SpellSlotsLevel9
	}

	return spellSlotAmount
}

func getWarlockSpellSlotInfo(spellSlotAmount [9]int) (int, int) {
	warlockSpellSlotAmount := 0
	warlockSpellSlotLevel := 0

	for i, levelSpellSlotAmount := range spellSlotAmount {
		if levelSpellSlotAmount != 0 {
			warlockSpellSlotAmount = levelSpellSlotAmount
			warlockSpellSlotLevel = i + 1
			break
		}
	}

	return warlockSpellSlotAmount, warlockSpellSlotLevel
}

func createCastingInfoFromDndApiClassWithLevels(dndApiClassWithLevels *infrastructure.DndApiClassWithLevels, level int, classTypedName domain.ClassName, jsonSpellRepository *infrastructure.JsonSpellRepository, abilityScoreList domain.AbilityScoreList, proficiencyBonus int) (*domain.ClassSpellcastingInfo, *domain.ClassWarlockCastingInfo) {
	var classSpellcastingInfo *domain.ClassSpellcastingInfo
	var classWarlockCastingInfo *domain.ClassWarlockCastingInfo

	dndApiClassLevel, err := dndApiClassWithLevels.GetClassLevelByLevel(level)
	if err != nil {
		log.Fatal(err)
	}

	if dndApiClassWithLevels.Spellcasting != nil && dndApiClassLevel.Spellcasting != nil {
		maxKnownCantrips := 0
		if dndApiClassLevel.Spellcasting.CantripsKnown != nil {
			maxKnownCantrips = *dndApiClassLevel.Spellcasting.CantripsKnown
		}

		spellList, err := CreateInitialSpellListForClass(classTypedName, jsonSpellRepository)
		if err != nil {
			log.Fatal(err)
		}

		spellSlotAmount := getSpellSlotAmountFromDndApiClassLevelSpellcasting(dndApiClassLevel.Spellcasting)

		spellcastingAbilityScoreName, err := dndApiClassWithLevels.Spellcasting.GetSpellcastingAbilityAsAbilityScoreName()
		if err != nil {
			log.Fatal(err)
		}
		spellcastingAbilityScore, err := abilityScoreList.GetByName(*spellcastingAbilityScoreName)
		if err != nil {
			log.Fatal(err)
		}

		spellSaveDC := 8 + proficiencyBonus + spellcastingAbilityScore.Modifier

		spellAttackBonus := proficiencyBonus + spellcastingAbilityScore.Modifier

		switch classTypedName {
		case domain.BARD, domain.RANGER, domain.SORCERER:
			if spellList == nil {
				err := fmt.Errorf(INITIALISED_SPELL_LIST_REQUIRED, string(classTypedName))
				log.Fatal(err)
			}

			classSpellcastingInfoValue := domain.NewClassSpellcastingInfo(
				maxKnownCantrips,
				dndApiClassLevel.Spellcasting.SpellsKnown,
				nil,
				*spellList,
				spellSlotAmount,
				spellcastingAbilityScore,
				spellSaveDC,
				spellAttackBonus,
			)
			classSpellcastingInfo = &classSpellcastingInfoValue
		case domain.CLERIC, domain.DRUID, domain.PALADIN, domain.WIZARD:
			if spellList == nil {
				err := fmt.Errorf(INITIALISED_SPELL_LIST_REQUIRED, string(classTypedName))
				log.Fatal(err)
			}

			maxPreparedSpells := max(1, spellcastingAbilityScore.Modifier+level)

			classSpellcastingInfoValue := domain.NewClassSpellcastingInfo(
				maxKnownCantrips,
				dndApiClassLevel.Spellcasting.SpellsKnown,
				&maxPreparedSpells,
				*spellList,
				spellSlotAmount,
				spellcastingAbilityScore,
				spellSaveDC,
				spellAttackBonus,
			)
			classSpellcastingInfo = &classSpellcastingInfoValue
		case domain.WARLOCK:
			if spellList == nil {
				err := errors.New(INITIALISED_SPELL_LIST_REQUIRED_WARLOCK)
				log.Fatal(err)
			}

			warlockSpellSlotAmount, warlockSpellSlotLevel := getWarlockSpellSlotInfo(spellSlotAmount)

			classWarlockCastingInfoValue := domain.NewClassWarlockCastingInfo(
				maxKnownCantrips,
				*dndApiClassLevel.Spellcasting.SpellsKnown,
				*spellList,
				warlockSpellSlotAmount,
				warlockSpellSlotLevel,
				spellcastingAbilityScore,
				spellSaveDC,
				spellAttackBonus,
			)
			classWarlockCastingInfo = &classWarlockCastingInfoValue
		case domain.BARBARIAN, domain.FIGHTER, domain.MONK, domain.ROGUE:
			fmt.Println("class should not have spellcasting, but it does according to the API. The API will be ignored in this case")
		default:
			err = errors.New(UNKNOWN_CLASS_WITH_SPELLCASTING)
			log.Fatal(err)
		}
	} else if dndApiClassWithLevels.Spellcasting != nil {
		err := errors.New(SPELLCASTING_NOT_DEFINED_PER_LEVEL)
		log.Fatal(err)
	} else if dndApiClassLevel.Spellcasting != nil {
		err := errors.New(NON_SPELLCASTER_HAS_SPELLCASTING_PER_LEVEL)
		log.Fatal(err)
	}

	return classSpellcastingInfo, classWarlockCastingInfo
}

func CreateClassFromDndApiClassWithLevels(dndApiClassWithLevels *infrastructure.DndApiClassWithLevels, level int, proficiencyBonus int, abilityScoreList domain.AbilityScoreList, jsonSpellRepository *infrastructure.JsonSpellRepository) domain.Class {
	if dndApiClassWithLevels == nil {
		err := errors.New(NIL_DND_API_CLASS_WITH_LEVELS)
		log.Fatal(err)
	}

	classTypedName, err := domain.ClassNameFromUntypedPotentialClassName(dndApiClassWithLevels.Name)
	if err != nil {
		log.Fatal(err)
	}

	skillProficiencies := []domain.SkillProficiencyName{}
	skillProficiencyChoices := dndApiClassWithLevels.GetSkillProficiencyChoices()
	for i := 0; i < skillProficiencyChoices.Choose; i++ {
		skillProficiencyName, err := domain.SkillProficiencyNameFromApiIndex(skillProficiencyChoices.From.Options[i].Item.Index)
		if err != nil {
			log.Fatal(err)
		}

		skillProficiencies = append(skillProficiencies, skillProficiencyName)
	}

	unarmoredArmorClassAbilityScoreModifierNameList := []domain.AbilityScoreName{domain.DEXTERITY}
	switch classTypedName {
	case domain.BARBARIAN:
		unarmoredArmorClassAbilityScoreModifierNameList = append(unarmoredArmorClassAbilityScoreModifierNameList, domain.CONSTITUTION)
	case domain.MONK:
		unarmoredArmorClassAbilityScoreModifierNameList = append(unarmoredArmorClassAbilityScoreModifierNameList, domain.WISDOM)
	}

	classSpellcastingInfo, classWarlockCastingInfo := createCastingInfoFromDndApiClassWithLevels(dndApiClassWithLevels, level, classTypedName, jsonSpellRepository, abilityScoreList, proficiencyBonus)

	return domain.NewClass(
		classTypedName,
		dndApiClassWithLevels.HitDie,
		level,
		skillProficiencies,
		unarmoredArmorClassAbilityScoreModifierNameList,
		classSpellcastingInfo,
		classWarlockCastingInfo,
	)
}

func EditClass(class *domain.Class, level int, proficiencyBonus int, abilityScoreList *domain.AbilityScoreList, dndApiClassWithLevels *infrastructure.DndApiClassWithLevels) {
	class.Level = level

	dndApiClassLevel, err := dndApiClassWithLevels.GetClassLevelByLevel(level)
	if err != nil {
		log.Fatal(err)
	}

	if dndApiClassLevel.Spellcasting != nil {
		maxKnownCantrips := 0
		if dndApiClassLevel.Spellcasting.CantripsKnown != nil {
			maxKnownCantrips = *dndApiClassLevel.Spellcasting.CantripsKnown
		}

		spellSlotAmount := getSpellSlotAmountFromDndApiClassLevelSpellcasting(dndApiClassLevel.Spellcasting)

		if class.ClassSpellcastingInfo != nil {
			spellcastingAbilityModifier := class.ClassSpellcastingInfo.SpellcastingAbility.Modifier

			var maxPreparedSpells *int
			if class.ClassSpellcastingInfo.MaxPreparedSpells != nil {
				maxPreparedSpellsValue := max(1, spellcastingAbilityModifier+level)
				maxPreparedSpells = &maxPreparedSpellsValue
			}

			spellSaveDC := 8 + proficiencyBonus + spellcastingAbilityModifier

			spellAttackBonus := proficiencyBonus + spellcastingAbilityModifier

			class.ClassSpellcastingInfo.MaxKnownCantrips = maxKnownCantrips
			class.ClassSpellcastingInfo.MaxKnownSpells = dndApiClassLevel.Spellcasting.SpellsKnown
			class.ClassSpellcastingInfo.MaxPreparedSpells = maxPreparedSpells
			class.ClassSpellcastingInfo.SpellSlotAmount = spellSlotAmount
			class.ClassSpellcastingInfo.SpellSaveDC = spellSaveDC
			class.ClassSpellcastingInfo.SpellAttackBonus = spellAttackBonus
		} else if class.ClassWarlockCastingInfo != nil {
			spellcastingAbilityModifier := class.ClassWarlockCastingInfo.SpellcastingAbility.Modifier

			warlockSpellSlotAmount, warlockSpellSlotLevel := getWarlockSpellSlotInfo(spellSlotAmount)

			spellSaveDC := 8 + proficiencyBonus + spellcastingAbilityModifier

			spellAttackBonus := proficiencyBonus + spellcastingAbilityModifier

			class.ClassWarlockCastingInfo.MaxKnownCantrips = maxKnownCantrips
			class.ClassWarlockCastingInfo.MaxKnownSpells = *dndApiClassLevel.Spellcasting.SpellsKnown
			class.ClassWarlockCastingInfo.SpellSlotAmount = warlockSpellSlotAmount
			class.ClassWarlockCastingInfo.SpellSlotLevel = warlockSpellSlotLevel
			class.ClassWarlockCastingInfo.SpellSaveDC = spellSaveDC
			class.ClassWarlockCastingInfo.SpellAttackBonus = spellAttackBonus
		} else {
			fmt.Println("according to the API, the class has spellcasting, but the spellcasting is not already defined; API ignored in this case")
		}
	}
}

func getDndApiClassListFromResponses(bodies [][]byte) []infrastructure.DndApiClass {
	dndApiClassList := []infrastructure.DndApiClass{}
	for _, body := range bodies {
		var dndApiClass infrastructure.DndApiClass
		err := json.Unmarshal(body, &dndApiClass)
		if err != nil {
			log.Fatal(err)
		}

		dndApiClassList = append(dndApiClassList, dndApiClass)
	}

	return dndApiClassList
}

func getDndApiClassWithLevelsListFromResponses(bodies [][]byte) [][]infrastructure.DndApiClassLevel {
	dndApiClassLevelsList := [][]infrastructure.DndApiClassLevel{}
	for _, body := range bodies {
		var dndApiClassLevels []infrastructure.DndApiClassLevel
		err := json.Unmarshal(body, &dndApiClassLevels)
		if err != nil {
			log.Fatal(err)
		}

		dndApiClassLevelsList = append(dndApiClassLevelsList, dndApiClassLevels)
	}

	return dndApiClassLevelsList
}

func (classService *ClassService) InitialiseClasses() {
	body, err := classService.dndApiGateway.Get("/api/2014/classes")
	if err != nil {
		log.Fatal(err)
	}

	endpoints := []string{}
	var dndApiReferenceList infrastructure.DndApiReferenceList
	err = json.Unmarshal(body, &dndApiReferenceList)
	if err != nil {
		log.Fatal(err)
	}
	for _, result := range dndApiReferenceList.Results {
		endpoints = append(endpoints, result.Url)
	}

	bodies, errors := classService.dndApiGateway.GetMultipleOrdered(endpoints)
	if len(errors) != 0 {
		for _, err := range errors {
			fmt.Println(err)
		}
		os.Exit(1)
	}

	dndApiClassList := getDndApiClassListFromResponses(bodies)

	endpoints = []string{}
	for _, dndApiClass := range dndApiClassList {
		endpoints = append(endpoints, dndApiClass.ClassLevelsUrl)
	}

	bodies, errors = classService.dndApiGateway.GetMultipleOrdered(endpoints)
	if len(errors) != 0 {
		for _, err := range errors {
			fmt.Println(err)
		}
		os.Exit(1)
	}

	dndApiClassLevelsList := getDndApiClassWithLevelsListFromResponses(bodies)

	dndApiClassWithLevelsList := []infrastructure.DndApiClassWithLevels{}
	for i, dndApiClass := range dndApiClassList {
		dndApiClassWithLevels := infrastructure.NewDndApiClassWithLevels(
			dndApiClass.Index,
			dndApiClass.Name,
			dndApiClass.HitDie,
			dndApiClass.ProficiencyChoices,
			dndApiClass.ClassLevelsUrl,
			dndApiClassLevelsList[i],
			dndApiClass.Spellcasting,
		)

		dndApiClassWithLevelsList = append(dndApiClassWithLevelsList, dndApiClassWithLevels)
	}

	infrastructure.SaveDndApiClassWithLevelsListAsJson("./data/classes.json", &dndApiClassWithLevelsList)
}
