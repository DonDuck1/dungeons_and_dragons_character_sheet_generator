package infrastructure

type DndApiClass struct {
	Index              string                         `json:"index"`
	Name               string                         `json:"name"`
	HitDie             int                            `json:"hit_die"`
	ProficiencyChoices []DndApiClassProficiencyChoice `json:"proficiency_choices"`
	ClassLevelsUrl     string                         `json:"class_levels"`
	Spellcasting       *DndApiClassSpellcasting       `json:"spellcasting"`
}
