package infrastructure

type DndApiRaceWithSubRaces struct {
	Index               string                        `json:"index"`
	Name                string                        `json:"name"`
	AbilityBonusList    []DndApiAbilityBonus          `json:"ability_bonuses"`
	AbilityBonusOptions *DndApiRaceAbilityScoreChoice `json:"ability_bonus_options"`
	SubRaceReferences   *[]DndApiReference            `json:"subrace_references"`
	SubRaceList         []DndApiSubRace               `json:"subraces"`
}

func NewDndApiRaceWithSubRaces(
	index string,
	name string,
	AbilityBonusList []DndApiAbilityBonus,
	AbilityBonusOptions *DndApiRaceAbilityScoreChoice,
	SubRaceReferences *[]DndApiReference,
	SubRaceList []DndApiSubRace,
) DndApiRaceWithSubRaces {
	return DndApiRaceWithSubRaces{
		Index:               index,
		Name:                name,
		AbilityBonusList:    AbilityBonusList,
		AbilityBonusOptions: AbilityBonusOptions,
		SubRaceReferences:   SubRaceReferences,
		SubRaceList:         SubRaceList,
	}
}

func (dndApiRaceWithSubRaces DndApiRaceWithSubRaces) GetDeepCopy() DndApiRaceWithSubRaces {
	var deepCopiedAbilityBonusOptions *DndApiRaceAbilityScoreChoice
	if dndApiRaceWithSubRaces.AbilityBonusOptions != nil {
		value := dndApiRaceWithSubRaces.AbilityBonusOptions.GetDeepCopy()
		deepCopiedAbilityBonusOptions = &value
	}

	var deepCopiedSubRaceReferences *[]DndApiReference
	if dndApiRaceWithSubRaces.SubRaceReferences != nil {
		value := make([]DndApiReference, len(*dndApiRaceWithSubRaces.SubRaceReferences))
		for i, subRaceReference := range *dndApiRaceWithSubRaces.SubRaceReferences {
			value[i] = subRaceReference.GetDeepCopy()
		}
		deepCopiedSubRaceReferences = &value
	}

	return NewDndApiRaceWithSubRaces(
		dndApiRaceWithSubRaces.Index,
		dndApiRaceWithSubRaces.Name,
		dndApiRaceWithSubRaces.AbilityBonusList,
		deepCopiedAbilityBonusOptions,
		deepCopiedSubRaceReferences,
		dndApiRaceWithSubRaces.SubRaceList,
	)
}
