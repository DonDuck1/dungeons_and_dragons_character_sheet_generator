package main

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
	"dungeons_and_dragons_character_sheet_generator/infrastructure"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

func InitialiseArmorAndShields(csvEquipmentRepository infrastructure.CsvEquipmentRepository, dndApiGateway infrastructure.DndApiGateway) {
	csvArmorAndShieldList, err := csvEquipmentRepository.GetByEquipmentType("Armor")
	if err != nil {
		log.Fatal(err)
	}

	endpoints := []string{}
	for _, armorOrShield := range *csvArmorAndShieldList {
		endpoint := "/api/2014/equipment?name=" + url.QueryEscape(armorOrShield.Name)
		endpoints = append(endpoints, endpoint)
	}

	bodies, errors := dndApiGateway.GetMultipleOrdered(endpoints)
	if len(errors) != 0 {
		for _, err := range errors {
			fmt.Println(err)
		}
		os.Exit(1)
	}

	endpoints = []string{}
	for index, body := range bodies {
		var dndApiReferenceList infrastructure.DndApiReferenceList
		err = json.Unmarshal(body, &dndApiReferenceList)
		if err != nil {
			log.Fatal(err)
		}
		for _, result := range dndApiReferenceList.Results {
			if strings.EqualFold((*csvArmorAndShieldList)[index].Name, result.Name) {
				endpoints = append(endpoints, result.Url)
			}
		}
	}

	bodies, errors = dndApiGateway.GetMultipleOrdered(endpoints)
	if len(errors) != 0 {
		for _, err := range errors {
			fmt.Println(err)
		}
		os.Exit(1)
	}

	shieldList := []domain.Shield{}
	armorList := []domain.Armor{}
	for _, body := range bodies {
		var dndApiArmorOrShield infrastructure.DndApiArmorOrShield
		err = json.Unmarshal(body, &dndApiArmorOrShield)
		if err != nil {
			log.Fatal(err)
		}
		if dndApiArmorOrShield.IsShield() {
			shield, err := dndApiArmorOrShield.AsShield()
			if err != nil {
				log.Fatal(err)
			}
			shieldList = append(shieldList, *shield)
		} else {
			armor, err := dndApiArmorOrShield.AsArmor()
			if err != nil {
				log.Fatal(err)
			}
			armorList = append(armorList, *armor)
		}
	}

	infrastructure.SaveShieldListAsJson("./data/shields.json", &shieldList)
	infrastructure.SaveArmorListAsJson("./data/armor.json", &armorList)
}

func InitialiseBackgrounds(dndApiGateway infrastructure.DndApiGateway) {
	body, err := dndApiGateway.Get("/api/2014/backgrounds")
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

	bodies, errors := dndApiGateway.GetMultipleOrdered(endpoints)
	if len(errors) != 0 {
		for _, err := range errors {
			fmt.Println(err)
		}
		os.Exit(1)
	}

	backgroundList := []domain.Background{}
	for _, body := range bodies {
		var dndApiBackground infrastructure.DndApiBackground
		err = json.Unmarshal(body, &dndApiBackground)
		if err != nil {
			log.Fatal(err)
		}
		background, err := dndApiBackground.AsBackground()
		if err != nil {
			log.Fatal(err)
		}

		backgroundList = append(backgroundList, *background)
	}

	infrastructure.SaveBackgroundListAsJson("./data/backgrounds.json", &backgroundList)
}

func InitialiseRaces(dndApiGateway infrastructure.DndApiGateway) {
	body, err := dndApiGateway.Get("/api/2014/races")
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

	bodies, errors := dndApiGateway.GetMultipleOrdered(endpoints)
	if len(errors) != 0 {
		for _, err := range errors {
			fmt.Println(err)
		}
		os.Exit(1)
	}

	raceList := []domain.Race{}
	for _, body := range bodies {
		var dndApiRace infrastructure.DndApiRace
		err = json.Unmarshal(body, &dndApiRace)
		if err != nil {
			log.Fatal(err)
		}
		race, err := dndApiRace.AsRace()
		if err != nil {
			log.Fatal(err)
		}

		raceList = append(raceList, *race)
	}

	infrastructure.SaveRaceListAsJson("./data/races.json", &raceList)
}

func InitialiseSpells(dndApiGateway infrastructure.DndApiGateway) {
	body, err := dndApiGateway.Get("/api/2014/spells")
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

	bodies, errors := dndApiGateway.GetMultipleOrdered(endpoints)
	if len(errors) != 0 {
		for _, err := range errors {
			fmt.Println(err)
		}
		os.Exit(1)
	}

	spells := []domain.Spell{}
	for _, body := range bodies {
		var dndApiSpell infrastructure.DndApiSpell
		err = json.Unmarshal(body, &dndApiSpell)
		if err != nil {
			log.Fatal(err)
		}
		spell, err := dndApiSpell.AsSpell()
		if err != nil {
			log.Fatal(err)
		}

		spells = append(spells, *spell)
	}

	infrastructure.SaveSpellsAsJson("./data/spells.json", &spells)
}

func InitialiseWeapons(csvEquipmentRepository infrastructure.CsvEquipmentRepository, dndApiGateway infrastructure.DndApiGateway) {
	csvWeaponList, err := csvEquipmentRepository.GetByEquipmentType("Weapon")
	if err != nil {
		log.Fatal(err)
	}

	endpoints := []string{}
	for _, weapon := range *csvWeaponList {
		endpoint := "/api/2014/equipment?name=" + url.QueryEscape(weapon.Name)
		endpoints = append(endpoints, endpoint)
	}

	bodies, errors := dndApiGateway.GetMultipleOrdered(endpoints)
	if len(errors) != 0 {
		for _, err := range errors {
			fmt.Println(err)
		}
		os.Exit(1)
	}

	endpoints = []string{}
	for index, body := range bodies {
		var dndApiReferenceList infrastructure.DndApiReferenceList
		err = json.Unmarshal(body, &dndApiReferenceList)
		if err != nil {
			log.Fatal(err)
		}
		for _, result := range dndApiReferenceList.Results {
			if strings.EqualFold((*csvWeaponList)[index].Name, result.Name) {
				endpoints = append(endpoints, result.Url)
			}
		}
	}

	bodies, errors = dndApiGateway.GetMultipleOrdered(endpoints)
	if len(errors) != 0 {
		for _, err := range errors {
			fmt.Println(err)
		}
		os.Exit(1)
	}

	weaponList := []domain.Weapon{}
	for _, body := range bodies {
		var dndApiWeapon infrastructure.DndApiWeapon
		err = json.Unmarshal(body, &dndApiWeapon)
		if err != nil {
			log.Fatal(err)
		}
		weaponList = append(weaponList, dndApiWeapon.AsWeapon())
	}

	infrastructure.SaveWeaponListAsJson("./data/weapons.json", &weaponList)
}
