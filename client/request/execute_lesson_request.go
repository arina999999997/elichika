package request

import (
	"elichika/generic"
)

type ExecuteLessonRequest struct {
	ExecuteLessonIds   generic.Array[int32] `json:"execute_lesson_ids"`
	ConsumedContentIds generic.Array[int32] `json:"consumed_content_ids"`
	SelectedDeckId     int32                `json:"selected_deck_id"`
	IsThreeTimes       bool                 `json:"is_three_times"`
}
