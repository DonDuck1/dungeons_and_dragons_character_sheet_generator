package infrastructure

import "dungeons_and_dragons_character_sheet_generator/domain"

type DndApiAbilityBonus struct {
	AbilityScore DndApiReference `json:"ability_score"`
	Bonus        int             `json:"bonus"`
}

func (dndApiAbilityBonus DndApiAbilityBonus) AsAbilityScoreImprovement() (*domain.AbilityScoreImprovement, error) {
	abilityScoreName, err := domain.AbilityScoreNameFromApiIndex(dndApiAbilityBonus.AbilityScore.Index)
	if err != nil {
		return nil, err
	}

	abilityScoreImprovement := domain.NewAbilityScoreImprovement(
		abilityScoreName,
		dndApiAbilityBonus.Bonus,
	)
	return &abilityScoreImprovement, nil
}
