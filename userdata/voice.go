package userdata

import (
	"elichika/model"
	"elichika/utils"
)

func (session *Session) UpdateVoice(naviVoiceMasterID int, isNew bool) {
	userVoice := model.UserVoice{}
	exists, err := session.Db.Table("u_voice").Where("user_id = ? AND navi_voice_master_id = ?",
		session.UserStatus.UserID, naviVoiceMasterID).Get(&userVoice)
	utils.CheckErr(err)
	if exists {
		if userVoice.IsNew == isNew {
			return
		}
		userVoice.IsNew = isNew
		_, err = session.Db.Table("u_voice").Where("user_id = ? AND navi_voice_master_id = ?",
			session.UserStatus.UserID, naviVoiceMasterID).AllCols().Update(userVoice)
		utils.CheckErr(err)
	} else {
		userVoice = model.UserVoice{
			UserID:            session.UserStatus.UserID,
			NaviVoiceMasterID: naviVoiceMasterID,
			IsNew:             isNew,
		}
		_, err = session.Db.Table("u_voice").Insert(userVoice)
		utils.CheckErr(err)
	}
	session.UserModel.UserVoiceByVoiceID.PushBack(userVoice)
}

func init() {
	addGenericTableFieldPopulator("u_voice", "UserVoiceByVoiceID")
}
