package client

import (
	"elichika/generic"
)

// this is the normal login bonus
type NaviLoginBonus struct {
	LoginBonusId           int32                           `json:"login_bonus_id"`
	LoginBonusRewards      generic.List[LoginBonusRewards] `json:"login_bonus_rewards"`
	BackgroundId           int32                           `json:"background_id"`
	WhiteboardTextureAsset *TextureStruktur                `json:"whiteboard_texture_asset"`
	StartAt                int64                           `json:"start_at"`
	EndAt                  int64                           `json:"end_at"`
	// these doesn't seem to do anything but they have to be present
	MaxPage     int32 `json:"max_page"`
	CurrentPage int32 `json:"current_page"`
}
