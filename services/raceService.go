package services

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
	"dungeons_and_dragons_character_sheet_generator/infrastructure"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type RaceService struct {
	dndApiGateway *infrastructure.DndApiGateway
}

func NewRaceService(dndApiGateway *infrastructure.DndApiGateway) *RaceService {
	return &RaceService{dndApiGateway: dndApiGateway}
}

func CreateRaceFromdndApiRaceWithSubRaces(chosenRaceName string, dndApiRaceWithSubRaces infrastructure.DndApiRaceWithSubRaces) (*domain.Race, error) {
	raceAbilityScoreImprovements := []domain.AbilityScoreImprovement{}
	for _, dndApiAbilityBonus := range dndApiRaceWithSubRaces.AbilityBonusList {
		abilityScoreImprovement, err := dndApiAbilityBonus.AsAbilityScoreImprovement()
		if err != nil {
			return nil, err
		}
		raceAbilityScoreImprovements = append(raceAbilityScoreImprovements, *abilityScoreImprovement)
	}

	var chosenDndApiSubRace *infrastructure.DndApiSubRace
	for _, subRace := range dndApiRaceWithSubRaces.SubRaceList {
		if strings.EqualFold(subRace.Name, chosenRaceName) {
			chosenDndApiSubRace = &subRace
		}
	}

	if dndApiRaceWithSubRaces.AbilityBonusOptions != nil {
		optionalRaceAbilityScoreImprovements := []domain.AbilityScoreImprovement{}
		for _, dndApiOptionalAbilityBonus := range dndApiRaceWithSubRaces.AbilityBonusOptions.From.Options {
			optionalAbilityScoreImprovement, err := dndApiOptionalAbilityBonus.AsAbilityScoreImprovement()
			if err != nil {
				return nil, err
			}

			optionalRaceAbilityScoreImprovements = append(optionalRaceAbilityScoreImprovements, *optionalAbilityScoreImprovement)
		}
		optionalAbilityScoreImprovementList := domain.NewOptionalAbilityScoreImprovementList(optionalRaceAbilityScoreImprovements, dndApiRaceWithSubRaces.AbilityBonusOptions.Choose)

		chosenOptionalRaceAbilityScoreImprovements := optionalAbilityScoreImprovementList.ChooseRandomAbilityScoreImprovements()
		raceAbilityScoreImprovements = append(raceAbilityScoreImprovements, chosenOptionalRaceAbilityScoreImprovements...)
	}

	if chosenDndApiSubRace == nil {
		race := domain.NewRace(
			dndApiRaceWithSubRaces.Name,
			raceAbilityScoreImprovements,
			nil,
		)

		return &race, nil
	}

	chosenSubRace, err := chosenDndApiSubRace.AsSubRace()
	if err != nil {
		return nil, err
	}

	race := domain.NewRace(
		dndApiRaceWithSubRaces.Name,
		raceAbilityScoreImprovements,
		chosenSubRace,
	)

	return &race, nil
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
	dndApiRaceList := []infrastructure.DndApiRace{}
	for _, body := range bodies {
		var dndApiRace infrastructure.DndApiRace
		err := json.Unmarshal(body, &dndApiRace)
		if err != nil {
			log.Fatal(err)
		}

		dndApiRaceList = append(dndApiRaceList, dndApiRace)
	}

	dndApiRaceWithSubRacesList := []infrastructure.DndApiRaceWithSubRaces{}
	for _, dndApiRace := range dndApiRaceList {
		endpoints := []string{}
		for _, subRaceReference := range *dndApiRace.SubRaceReferences {
			endpoints = append(endpoints, subRaceReference.Url)
		}

		bodies, errors = raceService.dndApiGateway.GetMultipleOrdered(endpoints)
		if len(errors) != 0 {
			for _, err := range errors {
				fmt.Println(err)
			}
			os.Exit(1)
		}
		dndApiSubRaceList := []infrastructure.DndApiSubRace{}
		for _, body := range bodies {
			var dndApiSubRace infrastructure.DndApiSubRace
			err := json.Unmarshal(body, &dndApiSubRace)
			if err != nil {
				log.Fatal(err)
			}

			dndApiSubRaceList = append(dndApiSubRaceList, dndApiSubRace)
		}

		dndApiRaceWithSubRaces := infrastructure.NewDndApiRaceWithSubRaces(
			dndApiRace.Index,
			dndApiRace.Name,
			dndApiRace.AbilityBonusList,
			dndApiRace.AbilityBonusOptions,
			dndApiRace.SubRaceReferences,
			dndApiSubRaceList,
		)

		dndApiRaceWithSubRacesList = append(dndApiRaceWithSubRacesList, dndApiRaceWithSubRaces)
	}

	infrastructure.SaveRaceListAsJson("./data/races.json", &dndApiRaceWithSubRacesList)
}
