package request

import (
	"elichika/generic"
)

type SkillEditResultRequest struct {
	DeckId           int32                                          `json:"deck_id"`
	SelectedSkillIds generic.Dictionary[int32, generic.List[int32]] `json:"selected_skill_ids"`
}
