package client

import (
	"elichika/generic"
)

type BootstrapNewBadge struct {
	IsNewMainStory                     bool                    `json:"is_new_main_story"`
	UnreceivedPresentBox               int32                   `json:"unreceived_present_box"`
	IsUnreceivedPresentBoxSubscription bool                    `json:"is_unreceived_present_box_subscription"`
	NoticeNewArrivalsIds               generic.List[int32]     `xorm:"json" json:"notice_new_arrivals_ids"`
	IsUpdateFriend                     bool                    `json:"is_update_friend"`
	UnreceivedMission                  int32                   `json:"unreceived_mission"`
	UnreceivedChallengeBeginner        int32                   `json:"unreceived_challenge_beginner"`
	DailyTheaterTodayId                generic.Nullable[int32] `xorm:"json" json:"daily_theater_today_id"`
}
