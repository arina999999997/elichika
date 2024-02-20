package client

type LessonDropItem struct {
	ContentType    int32 `json:"content_type" enum:"ContentType"`
	ContentId      int32 `json:"content_id"`
	ContentAmount  int32 `json:"content_amount"`
	DropRarity     int32 `json:"drop_rarity" enum:"LessonDropRarityType"`
	IsSubscription bool  `json:"is_subscription"`
}
