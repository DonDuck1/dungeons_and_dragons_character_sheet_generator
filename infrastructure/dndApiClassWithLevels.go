package infrastructure

import (
	"fmt"
	"strings"
)

type DndApiClassWithLevels struct {
	Index              string                         `json:"index"`
	Name               string                         `json:"name"`
	HitDie             int                            `json:"hit_die"`
	ProficiencyChoices []DndApiClassProficiencyChoice `json:"proficiency_choices"`
	ClassLevelsUrl     string                         `json:"class_levels_url"`
	ClassLevelList     []DndApiClassLevel             `json:"class_level_list"`
	Spellcasting       *DndApiClassSpellcasting       `json:"spellcasting"`
}

const (
	NO_CLASS_LEVEL_STRUCT_WITH_LEVEL string = "could not find class level struct for level %d"
)

func NewDndApiClassWithLevels(
	index string,
	name string,
	hitDie int,
	proficiencyChoices []DndApiClassProficiencyChoice,
	classLevelsUrl string,
	classLevelList []DndApiClassLevel,
	spellcasting *DndApiClassSpellcasting,
) DndApiClassWithLevels {
	return DndApiClassWithLevels{
		Index:              index,
		Name:               name,
		HitDie:             hitDie,
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

	err := fmt.Errorf(NO_CLASS_LEVEL_STRUCT_WITH_LEVEL, level)
	return nil, err
}

func (dndApiClassWithLevels DndApiClassWithLevels) GetDeepCopy() DndApiClassWithLevels {
	deepCopiedClassLevelList := make([]DndApiClassLevel, len(dndApiClassWithLevels.ClassLevelList))
	for i, classLevel := range dndApiClassWithLevels.ClassLevelList {
		deepCopiedClassLevelList[i] = classLevel.GetDeepCopy()
	}

	var deepCopiedSpellcasting *DndApiClassSpellcasting
	if dndApiClassWithLevels.Spellcasting != nil {
		value := dndApiClassWithLevels.Spellcasting.GetDeepCopy()
		deepCopiedSpellcasting = &value
	}

	return NewDndApiClassWithLevels(
		dndApiClassWithLevels.Index,
		dndApiClassWithLevels.Name,
		dndApiClassWithLevels.HitDie,
		dndApiClassWithLevels.ProficiencyChoices,
		dndApiClassWithLevels.ClassLevelsUrl,
		deepCopiedClassLevelList,
		deepCopiedSpellcasting,
	)
}
