package infrastructure

type DndApiClassProficiencyChoice struct {
	Choose int                                    `json:"choose"`
	From   DndApiClassProficiencyChoiceOptionList `json:"from"`
}
