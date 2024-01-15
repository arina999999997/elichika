package request

type ChangeNameLessonDeckRequest struct {
	DeckId   int32  `json:"deck_id"`
	DeckName string `json:"deck_name"`
}