package services

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
	"dungeons_and_dragons_character_sheet_generator/infrastructure"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type RaceService struct {
	dndApiGateway *infrastructure.DndApiGateway
}

func NewRaceService(dndApiGateway *infrastructure.DndApiGateway) *RaceService {
	return &RaceService{dndApiGateway: dndApiGateway}
}

func (raceService *RaceService) InitialiseRaces() {
	body, err := raceService.dndApiGateway.Get("/api/2014/races")
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

	bodies, errors := raceService.dndApiGateway.GetMultipleOrdered(endpoints)
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
