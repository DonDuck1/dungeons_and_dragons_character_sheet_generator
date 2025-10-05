package infrastructure

import "dungeons_and_dragons_character_sheet_generator/domain"

type DndApiSubRace struct {
	Index            string               `json:"index"`
	Name             string               `json:"name"`
	Race             DndApiReference      `json:"race"`
	AbilityBonusList []DndApiAbilityBonus `json:"ability_bonuses"`
}

func (dndApiSubRace DndApiSubRace) AsSubRace() (*domain.SubRace, error) {
	abilityScoreImprovements := []domain.AbilityScoreImprovement{}
	for _, dndApiAbilityBonus := range dndApiSubRace.AbilityBonusList {
		abilityScoreImprovement, err := dndApiAbilityBonus.AsAbilityScoreImprovement()
		if err != nil {
			return nil, err
		}
		abilityScoreImprovements = append(abilityScoreImprovements, *abilityScoreImprovement)
	}

	subRace := domain.NewSubRace(dndApiSubRace.Name, abilityScoreImprovements)
	return &subRace, nil
}
