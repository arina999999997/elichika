package user_member

import (
	"elichika/userdata"
	"elichika/utils"
)

func userMemberFinalizer(session *userdata.Session) {
	for _, member := range session.UserModel.UserMemberByMemberId.Map {
		affected, err := session.Db.Table("u_member").
			Where("user_id = ? AND member_master_id = ?", session.UserId, member.MemberMasterId).AllCols().Update(*member)
		utils.CheckErr(err)
		if affected == 0 {
			userdata.GenericDatabaseInsert(session, "u_member", *member)
		}
	}
}

func init() {
	userdata.AddFinalizer(userMemberFinalizer)
}
