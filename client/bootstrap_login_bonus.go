package client

import (
	"elichika/generic"
)

type BootstrapLoginBonus struct {
	Event2DLoginBonuses    generic.List[IllustLoginBonus]         `json:"event_2d_login_bonuses"`
	LoginBonuses           generic.List[NaviLoginBonus]           `json:"login_bonuses"`
	Event3DLoginBonus      generic.List[NaviLoginBonus]           `json:"event_3d_login_bonuses"`
	BeginnerLoginBonuses   generic.List[NaviLoginBonus]           `json:"beginner_login_bonuses"`
	ComebackLoginBonuses   generic.List[IllustLoginBonus]         `json:"comeback_login_bonuses"`
	BirthdayLoginBonuses   generic.List[NaviLoginBonus]           `json:"birthday_login_bonuses"`
	BirthdayMember         generic.List[LoginBonusBirthDayMember] `json:"birth_day_member"`
	NextLoginBonsReceiveAt int64                                  `json:"next_login_bons_receive_at"` // this is correct
}
