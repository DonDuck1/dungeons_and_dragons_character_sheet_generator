package infrastructure

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
	"strings"
)

type DndApiRaceWithSubRaces struct {
	Index               string                        `json:"index"`
	Name                string                        `json:"name"`
	AbilityBonusList    []DndApiAbilityBonus          `json:"ability_bonuses"`
	AbilityBonusOptions *DndApiRaceAbilityScoreChoice `json:"ability_bonus_options"`
	SubRaceReferences   *[]DndApiReference            `json:"subrace_references"`
	SubRaceList         []DndApiSubRace               `json:"subraces"`
}

func NewDndApiRaceWithSubRaces(
	index string,
	name string,
	AbilityBonusList []DndApiAbilityBonus,
	AbilityBonusOptions *DndApiRaceAbilityScoreChoice,
	SubRaceReferences *[]DndApiReference,
	SubRaceList []DndApiSubRace,
) DndApiRaceWithSubRaces {
	return DndApiRaceWithSubRaces{
		Index:               index,
		Name:                name,
		AbilityBonusList:    AbilityBonusList,
		AbilityBonusOptions: AbilityBonusOptions,
		SubRaceReferences:   SubRaceReferences,
		SubRaceList:         SubRaceList,
	}
}

func (dndApiRaceWithSubRaces DndApiRaceWithSubRaces) AsRace(chosenRaceName string) (*domain.Race, error) {
	raceAbilityScoreImprovements := []domain.AbilityScoreImprovement{}
	for _, dndApiAbilityBonus := range dndApiRaceWithSubRaces.AbilityBonusList {
		abilityScoreImprovement, err := dndApiAbilityBonus.AsAbilityScoreImprovement()
		if err != nil {
			return nil, err
		}
		raceAbilityScoreImprovements = append(raceAbilityScoreImprovements, *abilityScoreImprovement)
	}

	var chosenDndApiSubRace *DndApiSubRace
	for i, subRace := range dndApiRaceWithSubRaces.SubRaceList {
		if strings.EqualFold(subRace.Name, chosenRaceName) {
			chosenDndApiSubRace = &dndApiRaceWithSubRaces.SubRaceList[i] // Use index to point to actual object, not the temporary copy of the loop
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
			2,
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
		2,
	)

	return &race, nil
}
