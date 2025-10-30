package domain

type SkillProficiencyList struct {
	Acrobatics     SkillProficiency
	AnimalHandling SkillProficiency
	Arcana         SkillProficiency
	Athletics      SkillProficiency
	Deception      SkillProficiency
	History        SkillProficiency
	Insight        SkillProficiency
	Intimidation   SkillProficiency
	Investigation  SkillProficiency
	Medicine       SkillProficiency
	Nature         SkillProficiency
	Perception     SkillProficiency
	Performance    SkillProficiency
	Persuasion     SkillProficiency
	Religion       SkillProficiency
	SleightOfHand  SkillProficiency
	Stealth        SkillProficiency
	Survival       SkillProficiency
}

func NewSkillProficiencyList(abilityScoreList *AbilityScoreList, skillProficiencies []SkillProficiencyName, proficiencyBonus int) SkillProficiencyList {
	acrobatics := NewSkillProficiency(ACROBATICS, false, 0, &abilityScoreList.Dexterity, proficiencyBonus)
	animalHandling := NewSkillProficiency(ANIMAL_HANDLING, false, 0, &abilityScoreList.Wisdom, proficiencyBonus)
	arcana := NewSkillProficiency(ARCANA, false, 0, &abilityScoreList.Intelligence, proficiencyBonus)
	athletics := NewSkillProficiency(ATHLETICS, false, 0, &abilityScoreList.Strength, proficiencyBonus)
	deception := NewSkillProficiency(DECEPTION, false, 0, &abilityScoreList.Charisma, proficiencyBonus)
	history := NewSkillProficiency(HISTORY, false, 0, &abilityScoreList.Intelligence, proficiencyBonus)
	insight := NewSkillProficiency(INSIGHT, false, 0, &abilityScoreList.Wisdom, proficiencyBonus)
	intimidation := NewSkillProficiency(INTIMIDATION, false, 0, &abilityScoreList.Charisma, proficiencyBonus)
	investigation := NewSkillProficiency(INVESTIGATION, false, 0, &abilityScoreList.Intelligence, proficiencyBonus)
	medicine := NewSkillProficiency(MEDICINE, false, 0, &abilityScoreList.Wisdom, proficiencyBonus)
	nature := NewSkillProficiency(NATURE, false, 0, &abilityScoreList.Intelligence, proficiencyBonus)
	perception := NewSkillProficiency(PERCEPTION, false, 0, &abilityScoreList.Wisdom, proficiencyBonus)
	performance := NewSkillProficiency(PERFORMANCE, false, 0, &abilityScoreList.Charisma, proficiencyBonus)
	persuasion := NewSkillProficiency(PERSUASION, false, 0, &abilityScoreList.Charisma, proficiencyBonus)
	religion := NewSkillProficiency(RELIGION, false, 0, &abilityScoreList.Intelligence, proficiencyBonus)
	sleightOfHand := NewSkillProficiency(SLEIGHT_OF_HAND, false, 0, &abilityScoreList.Dexterity, proficiencyBonus)
	stealth := NewSkillProficiency(STEALTH, false, 0, &abilityScoreList.Dexterity, proficiencyBonus)
	survival := NewSkillProficiency(SURVIVAL, false, 0, &abilityScoreList.Wisdom, proficiencyBonus)

	for _, skillProficiency := range skillProficiencies {
		switch skillProficiency {
		case ACROBATICS:
			acrobatics.MakeProficient()
			acrobatics.CalculateModifier(proficiencyBonus)
		case ANIMAL_HANDLING:
			animalHandling.MakeProficient()
			animalHandling.CalculateModifier(proficiencyBonus)
		case ARCANA:
			arcana.MakeProficient()
			arcana.CalculateModifier(proficiencyBonus)
		case ATHLETICS:
			athletics.MakeProficient()
			athletics.CalculateModifier(proficiencyBonus)
		case DECEPTION:
			deception.MakeProficient()
			deception.CalculateModifier(proficiencyBonus)
		case HISTORY:
			history.MakeProficient()
			history.CalculateModifier(proficiencyBonus)
		case INSIGHT:
			insight.MakeProficient()
			insight.CalculateModifier(proficiencyBonus)
		case INTIMIDATION:
			intimidation.MakeProficient()
			intimidation.CalculateModifier(proficiencyBonus)
		case INVESTIGATION:
			investigation.MakeProficient()
			investigation.CalculateModifier(proficiencyBonus)
		case MEDICINE:
			medicine.MakeProficient()
			medicine.CalculateModifier(proficiencyBonus)
		case NATURE:
			nature.MakeProficient()
			nature.CalculateModifier(proficiencyBonus)
		case PERCEPTION:
			perception.MakeProficient()
			perception.CalculateModifier(proficiencyBonus)
		case PERFORMANCE:
			performance.MakeProficient()
			performance.CalculateModifier(proficiencyBonus)
		case PERSUASION:
			persuasion.MakeProficient()
			persuasion.CalculateModifier(proficiencyBonus)
		case RELIGION:
			religion.MakeProficient()
			religion.CalculateModifier(proficiencyBonus)
		case SLEIGHT_OF_HAND:
			sleightOfHand.MakeProficient()
			sleightOfHand.CalculateModifier(proficiencyBonus)
		case STEALTH:
			stealth.MakeProficient()
			stealth.CalculateModifier(proficiencyBonus)
		case SURVIVAL:
			survival.MakeProficient()
			survival.CalculateModifier(proficiencyBonus)
		}
	}

	return SkillProficiencyList{
		Acrobatics:     acrobatics,
		AnimalHandling: animalHandling,
		Arcana:         arcana,
		Athletics:      athletics,
		Deception:      deception,
		History:        history,
		Insight:        insight,
		Intimidation:   intimidation,
		Investigation:  investigation,
		Medicine:       medicine,
		Nature:         nature,
		Perception:     perception,
		Performance:    performance,
		Persuasion:     persuasion,
		Religion:       religion,
		SleightOfHand:  sleightOfHand,
		Stealth:        stealth,
		Survival:       survival,
	}
}

func (skillProficiencyList *SkillProficiencyList) UpdateSkillProficiencies(proficiencyBonus int) {
	skillProficiencyList.Acrobatics.CalculateModifier(proficiencyBonus)
	skillProficiencyList.AnimalHandling.CalculateModifier(proficiencyBonus)
	skillProficiencyList.Arcana.CalculateModifier(proficiencyBonus)
	skillProficiencyList.Athletics.CalculateModifier(proficiencyBonus)
	skillProficiencyList.Deception.CalculateModifier(proficiencyBonus)
	skillProficiencyList.History.CalculateModifier(proficiencyBonus)
	skillProficiencyList.Insight.CalculateModifier(proficiencyBonus)
	skillProficiencyList.Intimidation.CalculateModifier(proficiencyBonus)
	skillProficiencyList.Investigation.CalculateModifier(proficiencyBonus)
	skillProficiencyList.Medicine.CalculateModifier(proficiencyBonus)
	skillProficiencyList.Nature.CalculateModifier(proficiencyBonus)
	skillProficiencyList.Perception.CalculateModifier(proficiencyBonus)
	skillProficiencyList.Performance.CalculateModifier(proficiencyBonus)
	skillProficiencyList.Persuasion.CalculateModifier(proficiencyBonus)
	skillProficiencyList.Religion.CalculateModifier(proficiencyBonus)
	skillProficiencyList.SleightOfHand.CalculateModifier(proficiencyBonus)
	skillProficiencyList.Stealth.CalculateModifier(proficiencyBonus)
	skillProficiencyList.Survival.CalculateModifier(proficiencyBonus)
}

func appendSkillProficiencyToSliceOnlyIfProficient(skillProficiency SkillProficiency, proficientSkillProficiencies *[]SkillProficiency) {
	if skillProficiency.Proficient {
		*proficientSkillProficiencies = append(*proficientSkillProficiencies, skillProficiency)
	}
}

func (skillProficiencyList *SkillProficiencyList) GetSkillProficienciesThatAreProficient() *[]SkillProficiency {
	proficientSkillProficiencies := []SkillProficiency{}

	appendSkillProficiencyToSliceOnlyIfProficient(skillProficiencyList.Acrobatics, &proficientSkillProficiencies)
	appendSkillProficiencyToSliceOnlyIfProficient(skillProficiencyList.AnimalHandling, &proficientSkillProficiencies)
	appendSkillProficiencyToSliceOnlyIfProficient(skillProficiencyList.Arcana, &proficientSkillProficiencies)
	appendSkillProficiencyToSliceOnlyIfProficient(skillProficiencyList.Athletics, &proficientSkillProficiencies)
	appendSkillProficiencyToSliceOnlyIfProficient(skillProficiencyList.Deception, &proficientSkillProficiencies)
	appendSkillProficiencyToSliceOnlyIfProficient(skillProficiencyList.History, &proficientSkillProficiencies)
	appendSkillProficiencyToSliceOnlyIfProficient(skillProficiencyList.Insight, &proficientSkillProficiencies)
	appendSkillProficiencyToSliceOnlyIfProficient(skillProficiencyList.Intimidation, &proficientSkillProficiencies)
	appendSkillProficiencyToSliceOnlyIfProficient(skillProficiencyList.Investigation, &proficientSkillProficiencies)
	appendSkillProficiencyToSliceOnlyIfProficient(skillProficiencyList.Medicine, &proficientSkillProficiencies)
	appendSkillProficiencyToSliceOnlyIfProficient(skillProficiencyList.Nature, &proficientSkillProficiencies)
	appendSkillProficiencyToSliceOnlyIfProficient(skillProficiencyList.Perception, &proficientSkillProficiencies)
	appendSkillProficiencyToSliceOnlyIfProficient(skillProficiencyList.Performance, &proficientSkillProficiencies)
	appendSkillProficiencyToSliceOnlyIfProficient(skillProficiencyList.Persuasion, &proficientSkillProficiencies)
	appendSkillProficiencyToSliceOnlyIfProficient(skillProficiencyList.Religion, &proficientSkillProficiencies)
	appendSkillProficiencyToSliceOnlyIfProficient(skillProficiencyList.SleightOfHand, &proficientSkillProficiencies)
	appendSkillProficiencyToSliceOnlyIfProficient(skillProficiencyList.Stealth, &proficientSkillProficiencies)
	appendSkillProficiencyToSliceOnlyIfProficient(skillProficiencyList.Survival, &proficientSkillProficiencies)

	return &proficientSkillProficiencies
}
