package infrastructure

type DndApiRace struct {
	Index               string                        `json:"index"`
	Name                string                        `json:"name"`
	AbilityBonusList    []DndApiAbilityBonus          `json:"ability_bonuses"`
	AbilityBonusOptions *DndApiRaceAbilityScoreChoice `json:"ability_bonus_options"`
	SubRaceReferences   *[]DndApiReference            `json:"subraces"`
}
