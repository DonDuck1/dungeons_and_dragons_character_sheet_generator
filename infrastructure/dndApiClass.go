package infrastructure

import "strings"

type DndApiClass struct {
	Index              string                         `json:"index"`
	Name               string                         `json:"name"`
	ProficiencyChoices []DndApiClassProficiencyChoice `json:"proficiency_choices"`
	ClassLevelsUrl     string                         `json:"class_levels"`
	Spellcasting       *DndApiClassSpellcasting       `json:"spellcasting"`
}

func (dndApiClass DndApiClass) GetSkillProficiencyChoices() *DndApiClassProficiencyChoice {
	for _, dndApiClassProficiencyChoice := range dndApiClass.ProficiencyChoices {
		if strings.HasPrefix(dndApiClassProficiencyChoice.From.Options[0].Item.Index, "skill-") {
			return &dndApiClassProficiencyChoice
		}
	}

	return nil
}
