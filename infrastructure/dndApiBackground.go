package infrastructure

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
)

type DndApiBackground struct {
	Index                 string            `json:"index"`
	Name                  string            `json:"name"`
	StartingProficiencies []DndApiReference `json:"starting_proficiencies"`
}

func (dndApiBackground DndApiBackground) AsBackground() (*domain.Background, error) {
	skillProficiencies := []domain.SkillProficiencyName{}

	for _, startingProficiency := range dndApiBackground.StartingProficiencies {
		skillProficiency, err := domain.SkillProficiencyNameFromApiIndex(startingProficiency.Index)
		if err != nil {
			return nil, err
		}

		skillProficiencies = append(skillProficiencies, skillProficiency)
	}

	background := domain.NewBackground(dndApiBackground.Name, skillProficiencies)
	return &background, nil
}
