package user_card

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func GetOtherUserProfileUserCard(session *userdata.Session, otherUserId, cardMasterId int32) client.ProfileUserCard {
	res := client.ProfileUserCard{}
	exist, err := session.Db.Table("u_card").Where("user_id = ? AND card_master_id = ?", otherUserId, cardMasterId).
		Cols("card_master_id", "level", "is_awakening_image", "is_all_training_activated").Get(&res)
	utils.CheckErrMustExist(err, exist)
	// xorm has some sort of weird caching going on so we need to fetch to the field directly
	exist, err = session.Db.Table("u_card_play_count_stat").Where("user_id = ? AND card_master_id = ?", otherUserId, cardMasterId).
		Cols("live_join_count", "active_skill_play_count").Get(&res.LiveJoinCount, &res.ActiveSkillPlayCount)
	utils.CheckErrMustExist(err, exist)
	return res
}
