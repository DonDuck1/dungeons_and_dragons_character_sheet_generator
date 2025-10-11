package domain

type Class struct {
	Name                                            ClassName
	Level                                           int
	SkillProficiencies                              []SkillProficiencyName
	UnarmoredArmorClassAbilityScoreModifierNameList []AbilityScoreName
	ClassSpellcastingInfo                           *ClassSpellcastingInfo
	ClassWarlockCastingInfo                         *ClassWarlockCastingInfo
}

func NewClass(
	name ClassName,
	level int,
	skillProficiencies []SkillProficiencyName,
	unarmoredArmorClassAbilityScoreModifierNameList []AbilityScoreName,
	classSpellcastingInfo *ClassSpellcastingInfo,
	classWarlockCastingInfo *ClassWarlockCastingInfo,
) Class {
	if level < 1 {
		level = 1
	} else if level > 20 {
		level = 20
	}

	return Class{
		Name:               name,
		Level:              level,
		SkillProficiencies: skillProficiencies,
		UnarmoredArmorClassAbilityScoreModifierNameList: unarmoredArmorClassAbilityScoreModifierNameList,
		ClassSpellcastingInfo:                           classSpellcastingInfo,
		ClassWarlockCastingInfo:                         classWarlockCastingInfo,
	}
}
