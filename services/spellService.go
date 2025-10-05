package services

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
	"dungeons_and_dragons_character_sheet_generator/infrastructure"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type SpellService struct {
	dndApiGateway *infrastructure.DndApiGateway
}

func NewSpellService(dndApiGateway *infrastructure.DndApiGateway) *SpellService {
	return &SpellService{dndApiGateway: dndApiGateway}
}

func (spellService *SpellService) InitialiseSpells() {
	body, err := spellService.dndApiGateway.Get("/api/2014/spells")
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

	bodies, errors := spellService.dndApiGateway.GetMultipleOrdered(endpoints)
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
