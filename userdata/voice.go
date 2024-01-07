package userdata

import (
	"elichika/model"
	"elichika/utils"
)

func (session *Session) UpdateVoice(naviVoiceMasterId int, isNew bool) {
	userVoice := model.UserVoice{}
	exist, err := session.Db.Table("u_voice").Where("user_id = ? AND navi_voice_master_id = ?",
		session.UserStatus.UserId, naviVoiceMasterId).Get(&userVoice)
	utils.CheckErr(err)
	if exist {
		if userVoice.IsNew == isNew {
			return
		}
		userVoice.IsNew = isNew
	} else {
		userVoice = model.UserVoice{
			UserId:            session.UserStatus.UserId,
			NaviVoiceMasterId: naviVoiceMasterId,
			IsNew:             isNew,
		}
	}
	session.UserModel.UserVoiceByVoiceId.PushBack(userVoice)
}
func voiceFinalizer(session *Session) {
	for _, userVoice := range session.UserModel.UserVoiceByVoiceId.Objects {
		affected, err := session.Db.Table("u_voice").Where("user_id = ? AND navi_voice_master_id = ?",
			session.UserStatus.UserId, userVoice.NaviVoiceMasterId).AllCols().Update(userVoice)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_voice").Insert(userVoice)
			utils.CheckErr(err)
		}
	}
}
func init() {
	addFinalizer(voiceFinalizer)
	addGenericTableFieldPopulator("u_voice", "UserVoiceByVoiceId")
}
