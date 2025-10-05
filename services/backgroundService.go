package services

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
	"dungeons_and_dragons_character_sheet_generator/infrastructure"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type BackgroundService struct {
	dndApiGateway *infrastructure.DndApiGateway
}

func NewBackgroundService(dndApiGateway *infrastructure.DndApiGateway) *BackgroundService {
	return &BackgroundService{dndApiGateway: dndApiGateway}
}

func (backgroundService *BackgroundService) InitialiseBackgrounds() {
	body, err := backgroundService.dndApiGateway.Get("/api/2014/backgrounds")
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

	bodies, errors := backgroundService.dndApiGateway.GetMultipleOrdered(endpoints)
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
