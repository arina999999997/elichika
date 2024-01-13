package client

type UserInfoTriggerGachaPointExchangeRow struct {
	TriggerId          int64         `json:"trigger_id"`
	GachaMasterId      int32         `json:"gacha_master_id"`
	GachaTitle         LocalizedText `json:"gacha_title"`
	Point1MasterId     int32         `json:"point_1_master_id"`
	Point1BeforeAmount int32         `json:"point_1_before_amount"`
	Point1AfterAmount  int32         `json:"point_1_after_amount"`
	Point2MasterId     int32         `json:"point_2_master_id"`
	Point2BeforeAmount int32         `json:"point_2_before_amount"`
	Point2AfterAmount  int32         `json:"point_2_after_amount"`
}
