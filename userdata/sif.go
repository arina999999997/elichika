package userdata

import (
	"elichika/utils"
)

func schoolIdolFestivalIdRewardMissionFinalizer(session *Session) {
	for _, userSchoolIdolFestivalIdRewardMissionFinalizer := range session.UserModel.UserSchoolIdolFestivalIdRewardMissionById.Objects {
		affected, err := session.Db.Table("u_school_idol_festival_id_reward_mission").
			Where("user_id = ? AND school_idol_festival_id_reward_mission_master_id = ?",
				session.UserId, userSchoolIdolFestivalIdRewardMissionFinalizer.SchoolIdolFestivalIdRewardMissionMasterId).
			AllCols().Update(userSchoolIdolFestivalIdRewardMissionFinalizer)
		utils.CheckErr(err)
		if affected == 0 {
			genericDatabaseInsert(session, "u_school_idol_festival_id_reward_mission", userSchoolIdolFestivalIdRewardMissionFinalizer)
		}
	}
}

func init() {
	addFinalizer(schoolIdolFestivalIdRewardMissionFinalizer)
	addGenericTableFieldPopulator("u_school_idol_festival_id_reward_mission", "UserSchoolIdolFestivalIdRewardMissionById")
}
