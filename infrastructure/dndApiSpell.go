package infrastructure

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
)

type DndApiSpell struct {
	Index      string            `json:"index"`
	Name       string            `json:"name"`
	SpellRange string            `json:"range"`
	Level      int               `json:"level"`
	School     DndApiReference   `json:"school"`
	Classes    []DndApiReference `json:"classes"`
}

func (dndApiSpell DndApiSpell) AsSpell() (*domain.Spell, error) {
	classNameList := []domain.ClassName{}
	for _, dndApiClass := range dndApiSpell.Classes {
		className, err := domain.ClassNameFromApiIndex(dndApiClass.Index)
		if err != nil {
			return nil, err
		}

		classNameList = append(classNameList, className)
	}

	spell := domain.NewSpell(
		dndApiSpell.Name,
		dndApiSpell.Level,
		classNameList,
		dndApiSpell.School.Name,
		dndApiSpell.SpellRange,
		false,
	)

	return &spell, nil
}
