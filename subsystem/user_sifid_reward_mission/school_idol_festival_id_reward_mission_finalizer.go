package user_sifid_reward_mission

import (
	"elichika/userdata"
	"elichika/utils"
)

func schoolIdolFestivalIdRewardMissionFinalizer(session *userdata.Session) {
	for _, userSchoolIdolFestivalIdRewardMissionFinalizer := range session.UserModel.UserSchoolIdolFestivalIdRewardMissionById.Map {
		affected, err := session.Db.Table("u_school_idol_festival_id_reward_mission").
			Where("user_id = ? AND school_idol_festival_id_reward_mission_master_id = ?",
				session.UserId, userSchoolIdolFestivalIdRewardMissionFinalizer.SchoolIdolFestivalIdRewardMissionMasterId).
			AllCols().Update(*userSchoolIdolFestivalIdRewardMissionFinalizer)
		utils.CheckErr(err)
		if affected == 0 {
			userdata.GenericDatabaseInsert(session, "u_school_idol_festival_id_reward_mission", *userSchoolIdolFestivalIdRewardMissionFinalizer)
		}
	}
}

func init() {
	userdata.AddFinalizer(schoolIdolFestivalIdRewardMissionFinalizer)
}
