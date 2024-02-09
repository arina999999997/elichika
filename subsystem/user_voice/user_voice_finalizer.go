package user_voice

import (
	"elichika/userdata"
	"elichika/utils"
)

func userVoiceFinalizer(session *userdata.Session) {
	for _, userVoice := range session.UserModel.UserVoiceByVoiceId.Map {
		affected, err := session.Db.Table("u_voice").Where("user_id = ? AND navi_voice_master_id = ?",
			session.UserId, userVoice.NaviVoiceMasterId).AllCols().Update(*userVoice)
		utils.CheckErr(err)
		if affected == 0 {
			userdata.GenericDatabaseInsert(session, "u_voice", *userVoice)
		}
	}
}
func init() {
	userdata.AddFinalizer(userVoiceFinalizer)
}
