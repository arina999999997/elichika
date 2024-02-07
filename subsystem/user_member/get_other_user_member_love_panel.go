package user_member

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func GetOtherUserMemberLovePanel(session *userdata.Session, userId, memberId int32) client.MemberLovePanel {
	result := client.MemberLovePanel{}
	exist, err := session.Db.Table("u_member_love_panel").
		Where("user_id = ? AND member_id = ?", session.UserId, memberId).
		Get(&result)
	utils.CheckErr(err)
	if !exist {
		return client.MemberLovePanel{
			MemberId: memberId,
		}
	} else {
		return result
	}
}
