package userdata

import (
	"elichika/utils"
)

func schoolIdolFestivalIDRewardMissionFinalizer(session *Session) {
	for _, userSchoolIdolFestivalIDRewardMissionFinalizer := range session.UserModel.UserSchoolIdolFestivalIDRewardMissionByID.Objects {
		affected, err := session.Db.Table("u_school_idol_festival_id_reward_mission").
			Where("user_id = ? AND school_idol_festival_id_reward_mission_master_id = ?",
				session.UserStatus.UserID, userSchoolIdolFestivalIDRewardMissionFinalizer.SchoolIdolFestivalIDRewardMissionMasterID).
			AllCols().Update(userSchoolIdolFestivalIDRewardMissionFinalizer)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_school_idol_festival_id_reward_mission").
				Insert(userSchoolIdolFestivalIDRewardMissionFinalizer)
			utils.CheckErr(err)
		}
	}
}

func init() {
	addFinalizer(schoolIdolFestivalIDRewardMissionFinalizer)
	addGenericTableFieldPopulator("u_school_idol_festival_id_reward_mission", "UserSchoolIdolFestivalIDRewardMissionByID")
}
