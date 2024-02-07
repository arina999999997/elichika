package user_member

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func GetMember(session *userdata.Session, memberMasterId int32) client.UserMember {
	ptr, exist := session.UserModel.UserMemberByMemberId.Get(memberMasterId)
	if exist {
		return *ptr
	}
	member := client.UserMember{}
	exist, err := session.Db.Table("u_member").
		Where("user_id = ? AND member_master_id = ?", session.UserId, memberMasterId).Get(&member)
	utils.CheckErr(err)
	if !exist {
		// always inserted at login if not exist
		panic("member not found")
	}
	return member
}
