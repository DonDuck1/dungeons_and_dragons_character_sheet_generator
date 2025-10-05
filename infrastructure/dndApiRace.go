package infrastructure

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
)

type DndApiRace struct {
	Index               string                        `json:"index"`
	Name                string                        `json:"name"`
	AbilityBonusList    []DndApiAbilityBonus          `json:"ability_bonuses"`
	AbilityBonusOptions *DndApiRaceAbilityScoreChoice `json:"ability_bonus_options"`
}

func (dndApiRace DndApiRace) AsRace() (*domain.Race, error) {
	abilityScoreImprovements := []domain.AbilityScoreImprovement{}
	for _, dndApiAbilityBonus := range dndApiRace.AbilityBonusList {
		abilityScoreImprovement, err := dndApiAbilityBonus.AsAbilityScoreImprovement()
		if err != nil {
			return nil, err
		}
		abilityScoreImprovements = append(abilityScoreImprovements, *abilityScoreImprovement)
	}

	if dndApiRace.AbilityBonusOptions == nil {
		race := domain.NewRace(
			dndApiRace.Name,
			abilityScoreImprovements,
			nil,
			2,
		)

		return &race, nil
	}

	optionalAbilityScoreImprovements := []domain.AbilityScoreImprovement{}
	for _, dndApiOptionalAbilityBonus := range dndApiRace.AbilityBonusOptions.From.Options {
		optionalAbilityScoreImprovement, err := dndApiOptionalAbilityBonus.AsAbilityScoreImprovement()
		if err != nil {
			return nil, err
		}

		optionalAbilityScoreImprovements = append(optionalAbilityScoreImprovements, *optionalAbilityScoreImprovement)
	}
	optionalAbilityScoreImprovementList := domain.NewOptionalAbilityScoreImprovementList(optionalAbilityScoreImprovements, dndApiRace.AbilityBonusOptions.Choose)

	race := domain.NewRace(
		dndApiRace.Name,
		abilityScoreImprovements,
		&optionalAbilityScoreImprovementList,
		2,
	)

	return &race, nil
}
