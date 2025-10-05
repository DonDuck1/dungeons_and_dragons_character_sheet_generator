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
	for i, dndApiClassProficiencyChoice := range dndApiClass.ProficiencyChoices {
		if strings.HasPrefix(dndApiClassProficiencyChoice.From.Options[0].Item.Index, "skill-") {
			return &dndApiClass.ProficiencyChoices[i] // Use index to point to actual object, not the temporary copy of the loop
		}
	}

	return nil
}
