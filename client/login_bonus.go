package client

import (
	"elichika/generic"
)

// the struct in client also has plural
type LoginBonusRewards struct {
	Day                int32                   `json:"day"`
	Status             int32                   `json:"status" enum:"LoginBonusReceiveStatus"`
	ContentGrade       generic.Nullable[int32] `xorm:"json 'content_grade'" json:"content_grade" enum:"LoginBonusContentGrade"` // can be 0
	LoginBonusContents []Content               `json:"login_bonus_contents"`
}

// used for extra login bonus event that use an illustration?
type IllustLoginBonus struct {
	LoginBonusId      int32               `json:"login_bonus_id"`
	LoginBonusRewards []LoginBonusRewards `json:"login_bonus_rewards"`
	BackgroundId      int32               `json:"background_id"`
	StartAt           int64               `json:"start_at"`
	EndAt             int64               `json:"end_at"`
}

// this is the normal login bonus
type NaviLoginBonus struct {
	LoginBonusId           int32               `json:"login_bonus_id"`
	LoginBonusRewards      []LoginBonusRewards `json:"login_bonus_rewards"`
	BackgroundId           int32               `json:"background_id"`
	WhiteboardTextureAsset *TextureStruktur    `json:"whiteboard_texture_asset"`
	StartAt                int64               `json:"start_at"`
	EndAt                  int64               `json:"end_at"`
	// these doesn't seem to do anything but they have to be present
	MaxPage     int32 `json:"max_page"`
	CurrentPage int32 `json:"current_page"`
}

type LoginBonusBirthDayMember struct {
	MemberMasterId generic.Nullable[int32] `xorm:"json 'member_master_id'" json:"member_master_id"`
	SuitMasterId   generic.Nullable[int32] `xorm:"json 'suit_master_id'" json:"suit_master_id"`
}

type BootstrapLoginBonus struct {
	Event2DLoginBonuses    []IllustLoginBonus         `json:"event_2d_login_bonuses"`
	LoginBonuses           []NaviLoginBonus           `json:"login_bonuses"`
	Event3DLoginBonus      []NaviLoginBonus           `json:"event_3d_login_bonuses"`
	BeginnerLoginBonuses   []NaviLoginBonus           `json:"beginner_login_bonuses"`
	ComebackLoginBonuses   []IllustLoginBonus         `json:"comeback_login_bonuses"`
	BirthdayLoginBonuses   []NaviLoginBonus           `json:"birthday_login_bonuses"`
	BirthdayMember         []LoginBonusBirthDayMember `json:"birth_day_member"`
	NextLoginBonsReceiveAt int64                      `json:"next_login_bons_receive_at"` // this is correct
}
