package domain

type Class struct {
	Name                                            ClassName
	HitDie                                          int
	Level                                           int
	SkillProficiencies                              []SkillProficiencyName
	UnarmoredArmorClassAbilityScoreModifierNameList []AbilityScoreName
	ClassSpellcastingInfo                           *ClassSpellcastingInfo
	ClassWarlockCastingInfo                         *ClassWarlockCastingInfo
}

func NewClass(
	name ClassName,
	hitDie int,
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
		HitDie:             hitDie,
		Level:              level,
		SkillProficiencies: skillProficiencies,
		UnarmoredArmorClassAbilityScoreModifierNameList: unarmoredArmorClassAbilityScoreModifierNameList,
		ClassSpellcastingInfo:                           classSpellcastingInfo,
		ClassWarlockCastingInfo:                         classWarlockCastingInfo,
	}
}

func (class Class) GetMaxHitPointsFromClass(constitutionModifier int) int {
	maxHitPoints := class.HitDie + constitutionModifier

	for i := class.Level - 1; i > 0; i-- {
		maxHitPoints += (class.HitDie / 2) + 1 + constitutionModifier
	}

	return maxHitPoints
}
