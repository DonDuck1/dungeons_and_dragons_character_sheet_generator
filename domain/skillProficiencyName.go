package domain

import "fmt"

type SkillProficiencyName string

const (
	ACROBATICS      SkillProficiencyName = "Acrobatics"
	ANIMAL_HANDLING SkillProficiencyName = "Animal Handling"
	ARCANA          SkillProficiencyName = "Arcana"
	ATHLETICS       SkillProficiencyName = "Athletics"
	DECEPTION       SkillProficiencyName = "Deception"
	HISTORY         SkillProficiencyName = "History"
	INSIGHT         SkillProficiencyName = "Insight"
	INTIMIDATION    SkillProficiencyName = "Intimidation"
	INVESTIGATION   SkillProficiencyName = "Investigation"
	MEDICINE        SkillProficiencyName = "Medicine"
	NATURE          SkillProficiencyName = "Nature"
	PERCEPTION      SkillProficiencyName = "Perception"
	PERFORMANCE     SkillProficiencyName = "Performance"
	PERSUASION      SkillProficiencyName = "Persuasion"
	RELIGION        SkillProficiencyName = "Religion"
	SLEIGHT_OF_HAND SkillProficiencyName = "Sleight of Hand"
	STEALTH         SkillProficiencyName = "Stealth"
	SURVIVAL        SkillProficiencyName = "Survival"
)

const (
	NO_SKILL_PROFICIENCY_WITH_INDEX string = "no skill proficiency with index '%s' found"
)

func SkillProficiencyNameFromApiIndex(index string) (SkillProficiencyName, error) {
	switch index {
	case "skill-acrobatics":
		return ACROBATICS, nil
	case "skill-animal-handling":
		return ANIMAL_HANDLING, nil
	case "skill-arcana":
		return ARCANA, nil
	case "skill-athletics":
		return ATHLETICS, nil
	case "skill-deception":
		return DECEPTION, nil
	case "skill-history":
		return HISTORY, nil
	case "skill-insight":
		return INSIGHT, nil
	case "skill-intimidation":
		return INTIMIDATION, nil
	case "skill-investigation":
		return INVESTIGATION, nil
	case "skill-medicine":
		return MEDICINE, nil
	case "skill-nature":
		return NATURE, nil
	case "skill-perception":
		return PERCEPTION, nil
	case "skill-performance":
		return PERFORMANCE, nil
	case "skill-persuasion":
		return PERSUASION, nil
	case "skill-religion":
		return RELIGION, nil
	case "skill-sleight-of-hand":
		return SLEIGHT_OF_HAND, nil
	case "skill-stealth":
		return STEALTH, nil
	case "skill-survival":
		return SURVIVAL, nil
	default:
		return SURVIVAL, fmt.Errorf(NO_SKILL_PROFICIENCY_WITH_INDEX, index)
	}
}
