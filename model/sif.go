package model

import (
	"elichika/generic"
)

// this is saved only for archival reason
// maybe one day we can have some sort of cross game stuff with sif, but maybe the code for handling that isn't even there anymore
type UserSchoolIdolFestivalIdRewardMission struct {
	SchoolIdolFestivalIdRewardMissionMasterId int  `xorm:"pk 'school_idol_festival_id_reward_mission_master_id'" json:"school_idol_festival_id_reward_mission_master_id"`
	IsCleared                                 bool `xorm:"'is_cleared'" json:"is_cleared"`
	IsNew                                     bool `xorm:"'is_new'" json:"is_new"`
	Count                                     int  `xorm:"'count'" json:"count"`
}

func (usifidrm *UserSchoolIdolFestivalIdRewardMission) Id() int64 {
	return int64(usifidrm.SchoolIdolFestivalIdRewardMissionMasterId)
}

func init() {
	TableNameToInterface["u_school_idol_festival_id_reward_mission"] = generic.UserIdWrapper[UserSchoolIdolFestivalIdRewardMission]{}
}
