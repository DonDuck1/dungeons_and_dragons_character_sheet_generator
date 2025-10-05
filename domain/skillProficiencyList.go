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
	acrobatics := NewSkillProficiency(ACROBATICS, false, &abilityScoreList.Dexterity, proficiencyBonus)
	animalHandling := NewSkillProficiency(ANIMAL_HANDLING, false, &abilityScoreList.Wisdom, proficiencyBonus)
	arcana := NewSkillProficiency(ARCANA, false, &abilityScoreList.Intelligence, proficiencyBonus)
	athletics := NewSkillProficiency(ATHLETICS, false, &abilityScoreList.Strength, proficiencyBonus)
	deception := NewSkillProficiency(DECEPTION, false, &abilityScoreList.Charisma, proficiencyBonus)
	history := NewSkillProficiency(HISTORY, false, &abilityScoreList.Intelligence, proficiencyBonus)
	insight := NewSkillProficiency(INSIGHT, false, &abilityScoreList.Wisdom, proficiencyBonus)
	intimidation := NewSkillProficiency(INTIMIDATION, false, &abilityScoreList.Charisma, proficiencyBonus)
	investigation := NewSkillProficiency(INVESTIGATION, false, &abilityScoreList.Intelligence, proficiencyBonus)
	medicine := NewSkillProficiency(MEDICINE, false, &abilityScoreList.Wisdom, proficiencyBonus)
	nature := NewSkillProficiency(NATURE, false, &abilityScoreList.Intelligence, proficiencyBonus)
	perception := NewSkillProficiency(PERCEPTION, false, &abilityScoreList.Wisdom, proficiencyBonus)
	performance := NewSkillProficiency(PERFORMANCE, false, &abilityScoreList.Charisma, proficiencyBonus)
	persuasion := NewSkillProficiency(PERSUASION, false, &abilityScoreList.Charisma, proficiencyBonus)
	religion := NewSkillProficiency(RELIGION, false, &abilityScoreList.Intelligence, proficiencyBonus)
	sleightOfHand := NewSkillProficiency(SLEIGHT_OF_HAND, false, &abilityScoreList.Dexterity, proficiencyBonus)
	stealth := NewSkillProficiency(STEALTH, false, &abilityScoreList.Dexterity, proficiencyBonus)
	survival := NewSkillProficiency(SURVIVAL, false, &abilityScoreList.Wisdom, proficiencyBonus)

	for _, skillProficiency := range skillProficiencies {
		switch skillProficiency {
		case ACROBATICS:
			acrobatics.Proficient = true
			acrobatics.CalculateModifier(proficiencyBonus)
		case ANIMAL_HANDLING:
			animalHandling.Proficient = true
			animalHandling.CalculateModifier(proficiencyBonus)
		case ARCANA:
			arcana.Proficient = true
			arcana.CalculateModifier(proficiencyBonus)
		case ATHLETICS:
			athletics.Proficient = true
			athletics.CalculateModifier(proficiencyBonus)
		case DECEPTION:
			deception.Proficient = true
			deception.CalculateModifier(proficiencyBonus)
		case HISTORY:
			history.Proficient = true
			history.CalculateModifier(proficiencyBonus)
		case INSIGHT:
			insight.Proficient = true
			insight.CalculateModifier(proficiencyBonus)
		case INTIMIDATION:
			intimidation.Proficient = true
			intimidation.CalculateModifier(proficiencyBonus)
		case INVESTIGATION:
			investigation.Proficient = true
			investigation.CalculateModifier(proficiencyBonus)
		case MEDICINE:
			medicine.Proficient = true
			medicine.CalculateModifier(proficiencyBonus)
		case NATURE:
			nature.Proficient = true
			nature.CalculateModifier(proficiencyBonus)
		case PERCEPTION:
			perception.Proficient = true
			perception.CalculateModifier(proficiencyBonus)
		case PERFORMANCE:
			performance.Proficient = true
			performance.CalculateModifier(proficiencyBonus)
		case PERSUASION:
			persuasion.Proficient = true
			persuasion.CalculateModifier(proficiencyBonus)
		case RELIGION:
			religion.Proficient = true
			religion.CalculateModifier(proficiencyBonus)
		case SLEIGHT_OF_HAND:
			sleightOfHand.Proficient = true
			sleightOfHand.CalculateModifier(proficiencyBonus)
		case STEALTH:
			stealth.Proficient = true
			stealth.CalculateModifier(proficiencyBonus)
		case SURVIVAL:
			survival.Proficient = true
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

func (skillProficiencyList *SkillProficiencyList) GetSkillProficienciesThatAreProficient() *[]SkillProficiency {
	proficientSkillProficiencies := []SkillProficiency{}

	if skillProficiencyList.Acrobatics.Proficient {
		proficientSkillProficiencies = append(proficientSkillProficiencies, skillProficiencyList.Acrobatics)
	}
	if skillProficiencyList.AnimalHandling.Proficient {
		proficientSkillProficiencies = append(proficientSkillProficiencies, skillProficiencyList.AnimalHandling)
	}
	if skillProficiencyList.Arcana.Proficient {
		proficientSkillProficiencies = append(proficientSkillProficiencies, skillProficiencyList.Arcana)
	}
	if skillProficiencyList.Athletics.Proficient {
		proficientSkillProficiencies = append(proficientSkillProficiencies, skillProficiencyList.Athletics)
	}
	if skillProficiencyList.Deception.Proficient {
		proficientSkillProficiencies = append(proficientSkillProficiencies, skillProficiencyList.Deception)
	}
	if skillProficiencyList.History.Proficient {
		proficientSkillProficiencies = append(proficientSkillProficiencies, skillProficiencyList.History)
	}
	if skillProficiencyList.Insight.Proficient {
		proficientSkillProficiencies = append(proficientSkillProficiencies, skillProficiencyList.Insight)
	}
	if skillProficiencyList.Intimidation.Proficient {
		proficientSkillProficiencies = append(proficientSkillProficiencies, skillProficiencyList.Intimidation)
	}
	if skillProficiencyList.Investigation.Proficient {
		proficientSkillProficiencies = append(proficientSkillProficiencies, skillProficiencyList.Investigation)
	}
	if skillProficiencyList.Medicine.Proficient {
		proficientSkillProficiencies = append(proficientSkillProficiencies, skillProficiencyList.Medicine)
	}
	if skillProficiencyList.Nature.Proficient {
		proficientSkillProficiencies = append(proficientSkillProficiencies, skillProficiencyList.Nature)
	}
	if skillProficiencyList.Perception.Proficient {
		proficientSkillProficiencies = append(proficientSkillProficiencies, skillProficiencyList.Perception)
	}
	if skillProficiencyList.Performance.Proficient {
		proficientSkillProficiencies = append(proficientSkillProficiencies, skillProficiencyList.Performance)
	}
	if skillProficiencyList.Persuasion.Proficient {
		proficientSkillProficiencies = append(proficientSkillProficiencies, skillProficiencyList.Persuasion)
	}
	if skillProficiencyList.Religion.Proficient {
		proficientSkillProficiencies = append(proficientSkillProficiencies, skillProficiencyList.Religion)
	}
	if skillProficiencyList.SleightOfHand.Proficient {
		proficientSkillProficiencies = append(proficientSkillProficiencies, skillProficiencyList.SleightOfHand)
	}
	if skillProficiencyList.Stealth.Proficient {
		proficientSkillProficiencies = append(proficientSkillProficiencies, skillProficiencyList.Stealth)
	}
	if skillProficiencyList.Survival.Proficient {
		proficientSkillProficiencies = append(proficientSkillProficiencies, skillProficiencyList.Survival)
	}

	return &proficientSkillProficiencies
}
