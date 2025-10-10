package services

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
	"dungeons_and_dragons_character_sheet_generator/infrastructure"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ClassService struct {
	dndApiGateway *infrastructure.DndApiGateway
}

func NewClassService(dndApiGateway *infrastructure.DndApiGateway) *ClassService {
	return &ClassService{dndApiGateway: dndApiGateway}
}

func CreateClassFromDndApiClassWithLevels(dndApiClassWithLevels *infrastructure.DndApiClassWithLevels, level int, proficiencyBonus int, abilityScoreList domain.AbilityScoreList) domain.Class {
	if dndApiClassWithLevels == nil {
		err := fmt.Errorf("dndApiClassWithLevels provided is a nil value")
		log.Fatal(err)
	}

	mainClassTypedName, err := domain.ClassNameFromUntypedPotentialClassName(dndApiClassWithLevels.Name)
	if err != nil {
		log.Fatal(err)
	}

	dndApiClassLevel, err := dndApiClassWithLevels.GetClassLevelByLevel(level)
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

	var classSpellcastingInfo *domain.ClassSpellcastingInfo
	var classWarlockCastingInfo *domain.ClassWarlockCastingInfo
	if dndApiClassWithLevels.Spellcasting != nil && dndApiClassLevel.Spellcasting != nil {
		maxKnownCantrips := 0
		if dndApiClassLevel.Spellcasting.CantripsKnown != nil {
			maxKnownCantrips = *dndApiClassLevel.Spellcasting.CantripsKnown
		}

		spellList := domain.NewEmptySpellList()

		spellSlotAmount := [9]int{
			dndApiClassLevel.Spellcasting.SpellSlotsLevel1,
			dndApiClassLevel.Spellcasting.SpellSlotsLevel2,
			dndApiClassLevel.Spellcasting.SpellSlotsLevel3,
			dndApiClassLevel.Spellcasting.SpellSlotsLevel4,
			dndApiClassLevel.Spellcasting.SpellSlotsLevel5,
			0,
			0,
			0,
			0,
		}
		if dndApiClassLevel.Spellcasting.SpellSlotsLevel6 != nil {
			spellSlotAmount[5] = *dndApiClassLevel.Spellcasting.SpellSlotsLevel6
		}
		if dndApiClassLevel.Spellcasting.SpellSlotsLevel7 != nil {
			spellSlotAmount[6] = *dndApiClassLevel.Spellcasting.SpellSlotsLevel7
		}
		if dndApiClassLevel.Spellcasting.SpellSlotsLevel8 != nil {
			spellSlotAmount[7] = *dndApiClassLevel.Spellcasting.SpellSlotsLevel8
		}
		if dndApiClassLevel.Spellcasting.SpellSlotsLevel9 != nil {
			spellSlotAmount[8] = *dndApiClassLevel.Spellcasting.SpellSlotsLevel9
		}

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

		switch mainClassTypedName {
		case domain.BARD:
			classSpellcastingInfoValue := domain.NewClassSpellcastingInfo(
				maxKnownCantrips,
				dndApiClassLevel.Spellcasting.SpellsKnown,
				nil,
				spellList,
				spellSlotAmount,
				spellcastingAbilityScore,
				spellSaveDC,
				spellAttackBonus,
			)
			classSpellcastingInfo = &classSpellcastingInfoValue
		case domain.CLERIC:
			maxPreparedSpells := max(1, spellcastingAbilityScore.Modifier+level)

			classSpellcastingInfoValue := domain.NewClassSpellcastingInfo(
				maxKnownCantrips,
				dndApiClassLevel.Spellcasting.SpellsKnown,
				&maxPreparedSpells,
				spellList,
				spellSlotAmount,
				spellcastingAbilityScore,
				spellSaveDC,
				spellAttackBonus,
			)
			classSpellcastingInfo = &classSpellcastingInfoValue
		case domain.DRUID:
			maxPreparedSpells := max(1, spellcastingAbilityScore.Modifier+level)

			classSpellcastingInfoValue := domain.NewClassSpellcastingInfo(
				maxKnownCantrips,
				dndApiClassLevel.Spellcasting.SpellsKnown,
				&maxPreparedSpells,
				spellList,
				spellSlotAmount,
				spellcastingAbilityScore,
				spellSaveDC,
				spellAttackBonus,
			)
			classSpellcastingInfo = &classSpellcastingInfoValue
		case domain.PALADIN:
			maxPreparedSpells := max(1, spellcastingAbilityScore.Modifier+level)

			classSpellcastingInfoValue := domain.NewClassSpellcastingInfo(
				maxKnownCantrips,
				dndApiClassLevel.Spellcasting.SpellsKnown,
				&maxPreparedSpells,
				spellList,
				spellSlotAmount,
				spellcastingAbilityScore,
				spellSaveDC,
				spellAttackBonus,
			)
			classSpellcastingInfo = &classSpellcastingInfoValue
		case domain.RANGER:
			classSpellcastingInfoValue := domain.NewClassSpellcastingInfo(
				maxKnownCantrips,
				dndApiClassLevel.Spellcasting.SpellsKnown,
				nil,
				spellList,
				spellSlotAmount,
				spellcastingAbilityScore,
				spellSaveDC,
				spellAttackBonus,
			)
			classSpellcastingInfo = &classSpellcastingInfoValue
		case domain.SORCERER:
			classSpellcastingInfoValue := domain.NewClassSpellcastingInfo(
				maxKnownCantrips,
				dndApiClassLevel.Spellcasting.SpellsKnown,
				nil,
				spellList,
				spellSlotAmount,
				spellcastingAbilityScore,
				spellSaveDC,
				spellAttackBonus,
			)
			classSpellcastingInfo = &classSpellcastingInfoValue
		case domain.WIZARD:
			maxPreparedSpells := max(1, spellcastingAbilityScore.Modifier+level)

			classSpellcastingInfoValue := domain.NewClassSpellcastingInfo(
				maxKnownCantrips,
				dndApiClassLevel.Spellcasting.SpellsKnown,
				&maxPreparedSpells,
				spellList,
				spellSlotAmount,
				spellcastingAbilityScore,
				spellSaveDC,
				spellAttackBonus,
			)
			classSpellcastingInfo = &classSpellcastingInfoValue
		case domain.WARLOCK:
			warlockSpellSlotAmount := 0
			warlockSpellSlotLevel := 0

			for i, levelSpellSlotAmount := range spellSlotAmount {
				if levelSpellSlotAmount != 0 {
					warlockSpellSlotAmount = levelSpellSlotAmount
					warlockSpellSlotLevel = i + 1
					break
				}
			}

			classWarlockCastingInfoValue := domain.NewClassWarlockCastingInfo(
				maxKnownCantrips,
				*dndApiClassLevel.Spellcasting.SpellsKnown,
				spellList,
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
			err = fmt.Errorf("unknown class (with spellcasting) detected, character creation cannot continue")
			log.Fatal(err)
		}
	} else if dndApiClassWithLevels.Spellcasting != nil {
		err = fmt.Errorf("according to the API, the class has spellcasting, but the spellcasting is not defined per level; character creation cannot continue")
		log.Fatal(err)
	} else if dndApiClassLevel.Spellcasting != nil {
		err = fmt.Errorf("according to the API, the class does not have spellcasting, but the (non-existent) spellcasting is somehow actually defined per level; character creation cannot continue")
		log.Fatal(err)
	}

	return domain.NewClass(
		mainClassTypedName,
		level,
		skillProficiencies,
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

		spellSlotAmount := [9]int{
			dndApiClassLevel.Spellcasting.SpellSlotsLevel1,
			dndApiClassLevel.Spellcasting.SpellSlotsLevel2,
			dndApiClassLevel.Spellcasting.SpellSlotsLevel3,
			dndApiClassLevel.Spellcasting.SpellSlotsLevel4,
			dndApiClassLevel.Spellcasting.SpellSlotsLevel5,
			0,
			0,
			0,
			0,
		}
		if dndApiClassLevel.Spellcasting.SpellSlotsLevel6 != nil {
			spellSlotAmount[5] = *dndApiClassLevel.Spellcasting.SpellSlotsLevel6
		}
		if dndApiClassLevel.Spellcasting.SpellSlotsLevel7 != nil {
			spellSlotAmount[6] = *dndApiClassLevel.Spellcasting.SpellSlotsLevel7
		}
		if dndApiClassLevel.Spellcasting.SpellSlotsLevel8 != nil {
			spellSlotAmount[7] = *dndApiClassLevel.Spellcasting.SpellSlotsLevel8
		}
		if dndApiClassLevel.Spellcasting.SpellSlotsLevel9 != nil {
			spellSlotAmount[8] = *dndApiClassLevel.Spellcasting.SpellSlotsLevel9
		}

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

			warlockSpellSlotAmount := 0
			warlockSpellSlotLevel := 0

			for i, levelSpellSlotAmount := range spellSlotAmount {
				if levelSpellSlotAmount != 0 {
					warlockSpellSlotAmount = levelSpellSlotAmount
					warlockSpellSlotLevel = i + 1
					break
				}
			}

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
	dndApiClassList := []infrastructure.DndApiClass{}
	for _, body := range bodies {
		var dndApiClass infrastructure.DndApiClass
		err = json.Unmarshal(body, &dndApiClass)
		if err != nil {
			log.Fatal(err)
		}

		dndApiClassList = append(dndApiClassList, dndApiClass)
	}

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
	dndApiClassLevelsList := [][]infrastructure.DndApiClassLevel{}
	for _, body := range bodies {
		var dndApiClassLevels []infrastructure.DndApiClassLevel
		err = json.Unmarshal(body, &dndApiClassLevels)
		if err != nil {
			log.Fatal(err)
		}

		dndApiClassLevelsList = append(dndApiClassLevelsList, dndApiClassLevels)
	}

	dndApiClassWithLevelsList := []infrastructure.DndApiClassWithLevels{}
	for i, dndApiClass := range dndApiClassList {
		dndApiClassWithLevels := infrastructure.NewDndApiClassWithLevels(
			dndApiClass.Index,
			dndApiClass.Name,
			dndApiClass.ProficiencyChoices,
			dndApiClass.ClassLevelsUrl,
			dndApiClassLevelsList[i],
			dndApiClass.Spellcasting,
		)

		dndApiClassWithLevelsList = append(dndApiClassWithLevelsList, dndApiClassWithLevels)
	}

	infrastructure.SaveDndApiClassWithLevelsListAsJson("./data/classes.json", &dndApiClassWithLevelsList)
}
