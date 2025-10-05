package infrastructure

type DndApiRaceAbilityScoreChoice struct {
	Choose int                                    `json:"choose"`
	From   DndApiRaceAbilityScoreChoiceOptionList `json:"from"`
}
