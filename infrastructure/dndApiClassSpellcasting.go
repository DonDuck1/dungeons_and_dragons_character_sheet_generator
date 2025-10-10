package infrastructure

import "dungeons_and_dragons_character_sheet_generator/domain"

type DndApiClassSpellcasting struct {
	SpellcastingAbility DndApiReference `json:"spellcasting_ability"`
}

func NewDndApiClassSpellcasting(spellcastingAbility DndApiReference) DndApiClassSpellcasting {
	return DndApiClassSpellcasting{SpellcastingAbility: spellcastingAbility}
}

func (dndApiClassSpellcasting DndApiClassSpellcasting) GetSpellcastingAbilityAsAbilityScoreName() (*domain.AbilityScoreName, error) {
	abilityScoreName, err := domain.AbilityScoreNameFromApiIndex(dndApiClassSpellcasting.SpellcastingAbility.Index)
	if err != nil {
		return nil, err
	}

	return &abilityScoreName, nil
}

func (dndApiClassSpellcasting DndApiClassSpellcasting) GetDeepCopy() DndApiClassSpellcasting {
	return NewDndApiClassSpellcasting(dndApiClassSpellcasting.SpellcastingAbility)
}
