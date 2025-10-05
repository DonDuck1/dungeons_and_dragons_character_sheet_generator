package main

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
	"dungeons_and_dragons_character_sheet_generator/infrastructure"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
)

func CreateClassUsingApi(name domain.ClassName, level int, proficiencyBonus int, abilityScoreList domain.AbilityScoreList, dndApiGateway *infrastructure.DndApiGateway) domain.Class {
	body, err := dndApiGateway.Get("/api/2014/classes?name=" + url.QueryEscape(string(name)))
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
		if strings.EqualFold(string(name), result.Name) {
			endpoints = append(endpoints, result.Url)
		}
	}
	if len(endpoints) > 1 {
		fmt.Printf("something went wrong: the API provided %d classes with the exact same name", len(endpoints))
		fmt.Println("The first occurence of the class will be used")
	}

	body, err = dndApiGateway.Get(endpoints[0])
	if err != nil {
		log.Fatal(err)
	}
	var dndApiClass infrastructure.DndApiClass
	err = json.Unmarshal(body, &dndApiClass)
	if err != nil {
		log.Fatal(err)
	}

	body, err = dndApiGateway.Get(dndApiClass.ClassLevelsUrl + "/" + strconv.Itoa(level))
	if err != nil {
		log.Fatal(err)
	}
	var dndApiClassLevel infrastructure.DndApiClassLevel
	err = json.Unmarshal(body, &dndApiClassLevel)
	if err != nil {
		log.Fatal(err)
	}

	skillProficiencies := []domain.SkillProficiencyName{}
	skillProficiencyChoices := dndApiClass.GetSkillProficiencyChoices()
	for i := 0; i < skillProficiencyChoices.Choose; i++ {
		skillProficiencyName, err := domain.SkillProficiencyNameFromApiIndex(skillProficiencyChoices.From.Options[i].Item.Index)
		if err != nil {
			log.Fatal(err)
		}

		skillProficiencies = append(skillProficiencies, skillProficiencyName)
	}

	var classSpellcastingInfo *domain.ClassSpellcastingInfo
	var classWarlockCastingInfo *domain.ClassWarlockCastingInfo
	if dndApiClass.Spellcasting != nil && dndApiClassLevel.Spellcasting != nil {
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

		spellcastingAbilityScoreName, err := dndApiClass.Spellcasting.GetSpellcastingAbilityAsAbilityScoreName()
		if err != nil {
			log.Fatal(err)
		}
		spellcastingAbilityScore, err := abilityScoreList.GetByName(*spellcastingAbilityScoreName)
		if err != nil {
			log.Fatal(err)
		}

		spellSaveDC := 8 + proficiencyBonus + spellcastingAbilityScore.Modifier

		spellAttackBonus := proficiencyBonus + spellcastingAbilityScore.Modifier

		switch name {
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
				}
				warlockSpellSlotLevel = i + 1
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
			fmt.Println("Class should not have spellcasting, but it does according to the API. The API will be ignored in this case")
		default:
			err = fmt.Errorf("unknown class (with spellcasting) detected, character creation cannot continue")
			log.Fatal(err)
		}
	} else if dndApiClass.Spellcasting != nil {
		err = fmt.Errorf("according to the API, the class has spellcasting, but the spellcasting is not defined per level; character creation cannot continue")
		log.Fatal(err)
	} else if dndApiClassLevel.Spellcasting != nil {
		err = fmt.Errorf("according to the API, the class does not have spellcasting, but the (non-existent) spellcasting is somehow actually defined per level; character creation cannot continue")
		log.Fatal(err)
	}

	return domain.NewClass(
		name,
		level,
		skillProficiencies,
		classSpellcastingInfo,
		classWarlockCastingInfo,
	)
}
