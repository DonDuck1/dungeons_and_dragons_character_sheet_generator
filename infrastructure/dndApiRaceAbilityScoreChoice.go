package infrastructure

type DndApiRaceAbilityScoreChoice struct {
	Choose int                                    `json:"choose"`
	From   DndApiRaceAbilityScoreChoiceOptionList `json:"from"`
}

func NewDndApiRaceAbilityScoreChoice(choose int, from DndApiRaceAbilityScoreChoiceOptionList) DndApiRaceAbilityScoreChoice {
	return DndApiRaceAbilityScoreChoice{Choose: choose, From: from}
}

func (dndApiRaceAbilityScoreChoice DndApiRaceAbilityScoreChoice) GetDeepCopy() DndApiRaceAbilityScoreChoice {
	return NewDndApiRaceAbilityScoreChoice(dndApiRaceAbilityScoreChoice.Choose, dndApiRaceAbilityScoreChoice.From)
}
