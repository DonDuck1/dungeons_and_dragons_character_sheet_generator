package domain

type Class struct {
	name                    ClassName
	skillProficiencies      []SkillProficiencyName
	classSpellcastingInfo   *ClassSpellcastingInfo
	classWarlockCastingInfo *ClassWarlockCastingInfo
}
