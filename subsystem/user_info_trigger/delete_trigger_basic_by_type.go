package user_info_trigger

import (
	"elichika/userdata"
	"elichika/utils"

	"xorm.io/xorm"
)

func DeleteTriggerBasicByType(session *userdata.Session, infoTriggerType int32) {
	ids := []int64{}
	err := session.Db.Table("u_info_trigger_basic").Where("user_id = ? AND info_trigger_type = ?",
		session.UserId, infoTriggerType).Cols("trigger_id").Find(&ids)
	utils.CheckErr(err)
	for _, id := range ids {
		session.UserModel.UserInfoTriggerBasicByTriggerId.SetNull(id)
	}
}

func CleanUpTriggerBasicByType(userdata_db *xorm.Session, infoTriggerType int32) {
	_, err := userdata_db.Table("u_info_trigger_basic").Where("info_trigger_type = ?", infoTriggerType).Cols("trigger_id").Delete()
	utils.CheckErr(err)
}
