package user_voice

import (
	"elichika/client"
	"elichika/utils"
	"elichika/userdata"
)

func UpdateUserVoice(session *userdata.Session, naviVoiceMasterId int32, isNew bool) {
	userVoice := client.UserVoice{}
	exist, err := session.Db.Table("u_voice").Where("user_id = ? AND navi_voice_master_id = ?",
		session.UserId, naviVoiceMasterId).Get(&userVoice)
	utils.CheckErr(err)
	if exist {
		if userVoice.IsNew == isNew {
			return
		}
		userVoice.IsNew = isNew
	} else {
		userVoice = client.UserVoice{
			NaviVoiceMasterId: naviVoiceMasterId,
			IsNew:             isNew,
		}
	}
	session.UserModel.UserVoiceByVoiceId.Set(naviVoiceMasterId, userVoice)
}