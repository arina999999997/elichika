package request

type DrawGachaRequest struct {
	GachaDrawMasterId int32 `json:"gacha_draw_master_id"`
	ButtonDrawCount   int32 `json:"button_draw_count"`
}
