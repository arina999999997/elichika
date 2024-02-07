package user_member

import (
	"elichika/userdata"
	"elichika/utils"
)

func communicationMemberDetailBadgeFinalizer(session *userdata.Session) {
	// TODO(member, badge): this is only handled on the read side, new items won't change the badge
	for _, detailBadge := range session.UserModel.UserCommunicationMemberDetailBadgeById.Map {
		affected, err := session.Db.Table("u_communication_member_detail_badge").
			Where("user_id = ? AND member_master_id = ?", session.UserId, detailBadge.MemberMasterId).
			AllCols().Update(*detailBadge)
		utils.CheckErr(err)
		if affected == 0 {
			userdata.GenericDatabaseInsert(session, "u_communication_member_detail_badge", *detailBadge)
		}
	}
}

func init() {
	userdata.AddFinalizer(communicationMemberDetailBadgeFinalizer)
}
