package infrastructure

import (
	"fmt"
	"strings"
)

type DndApiClassWithLevels struct {
	Index              string                         `json:"index"`
	Name               string                         `json:"name"`
	ProficiencyChoices []DndApiClassProficiencyChoice `json:"proficiency_choices"`
	ClassLevelsUrl     string                         `json:"class_levels_url"`
	ClassLevelList     []DndApiClassLevel             `json:"class_level_list"`
	Spellcasting       *DndApiClassSpellcasting       `json:"spellcasting"`
}

func NewDndApiClassWithLevels(
	index string,
	name string,
	proficiencyChoices []DndApiClassProficiencyChoice,
	classLevelsUrl string,
	classLevelList []DndApiClassLevel,
	spellcasting *DndApiClassSpellcasting,
) DndApiClassWithLevels {
	return DndApiClassWithLevels{
		Index:              index,
		Name:               name,
		ProficiencyChoices: proficiencyChoices,
		ClassLevelsUrl:     classLevelsUrl,
		ClassLevelList:     classLevelList,
		Spellcasting:       spellcasting,
	}
}

func (dndApiFullClass DndApiClassWithLevels) GetSkillProficiencyChoices() *DndApiClassProficiencyChoice {
	for i, dndApiClassProficiencyChoice := range dndApiFullClass.ProficiencyChoices {
		if strings.HasPrefix(dndApiClassProficiencyChoice.From.Options[0].Item.Index, "skill-") {
			return &dndApiFullClass.ProficiencyChoices[i] // Use index to point to actual object, not the temporary copy of the loop
		}
	}

	return nil
}

func (dndApiClassWithLevels *DndApiClassWithLevels) GetClassLevelByLevel(level int) (*DndApiClassLevel, error) {
	classLevelList := dndApiClassWithLevels.ClassLevelList
	for i, classLevel := range classLevelList {
		if classLevel.Level == level {
			return &classLevelList[i], nil // Use index to point to actual object, not the temporary copy of the loop
		}
	}

	err := fmt.Errorf("could not find class level struct for level %d", level)
	return nil, err
}
