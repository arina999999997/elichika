package client

// this is saved only for archival reason
// maybe one day we can have some sort of cross game stuff with sif, but maybe the code for handling that isn't even there anymore
type UserSchoolIdolFestivalIdRewardMission struct {
	SchoolIdolFestivalIdRewardMissionMasterId int32 `xorm:"pk 'school_idol_festival_id_reward_mission_master_id'" json:"school_idol_festival_id_reward_mission_master_id"`
	IsCleared                                 bool  `xorm:"'is_cleared'" json:"is_cleared"`
	IsNew                                     bool  `xorm:"'is_new'" json:"is_new"`
	Count                                     int32 `xorm:"'count'" json:"count"`
}
