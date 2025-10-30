package services

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

type WeaponService struct {
	csvEquipmentRepository *infrastructure.CsvEquipmentRepository
	dndApiGateway          *infrastructure.DndApiGateway
}

func NewWeaponService(csvEquipmentRepository *infrastructure.CsvEquipmentRepository, dndApiGateway *infrastructure.DndApiGateway) *WeaponService {
	return &WeaponService{csvEquipmentRepository: csvEquipmentRepository, dndApiGateway: dndApiGateway}
}

func getWeaponListFromResponses(bodies [][]byte) []domain.Weapon {
	weaponList := []domain.Weapon{}
	for _, body := range bodies {
		var dndApiWeapon infrastructure.DndApiWeapon
		err := json.Unmarshal(body, &dndApiWeapon)
		if err != nil {
			log.Fatal(err)
		}
		weaponList = append(weaponList, dndApiWeapon.AsWeapon())
	}

	return weaponList
}

func (weaponService *WeaponService) InitialiseWeapons() {
	csvWeaponList, err := weaponService.csvEquipmentRepository.GetByEquipmentType("Weapon")
	if err != nil {
		log.Fatal(err)
	}

	endpoints := []string{}
	for _, weapon := range *csvWeaponList {
		endpoint := "/api/2014/equipment?name=" + url.QueryEscape(weapon.Name)
		endpoints = append(endpoints, endpoint)
	}

	bodies, errors := weaponService.dndApiGateway.GetMultipleOrdered(endpoints)
	if len(errors) != 0 {
		for _, err := range errors {
			fmt.Println(err)
		}
		os.Exit(1)
	}

	endpoints = []string{}
	for i, body := range bodies {
		var dndApiReferenceList infrastructure.DndApiReferenceList
		err = json.Unmarshal(body, &dndApiReferenceList)
		if err != nil {
			log.Fatal(err)
		}
		for _, result := range dndApiReferenceList.Results {
			if strings.EqualFold((*csvWeaponList)[i].Name, result.Name) {
				endpoints = append(endpoints, result.Url)
			}
		}
	}

	bodies, errors = weaponService.dndApiGateway.GetMultipleOrdered(endpoints)
	if len(errors) != 0 {
		for _, err := range errors {
			fmt.Println(err)
		}
		os.Exit(1)
	}

	weaponList := getWeaponListFromResponses(bodies)

	infrastructure.SaveWeaponListAsJson("./data/weapons.json", &weaponList)
}
