package user_social

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func GetOtherUserProfileInfomation(session *userdata.Session, otherUserId int32) client.ProfileInfomation {
	otherUserStatus := client.UserStatus{}
	exist, err := session.Db.Table("u_status").Where("user_id = ?", otherUserId).Get(&otherUserStatus)
	utils.CheckErrMustExist(err, exist)
	otherUser := GetOtherUser(session, otherUserId)
	profileInfomation := client.ProfileInfomation{
		BasicInfo:                 otherUser,
		MemberGuildMemberMasterId: otherUserStatus.MemberGuildMemberMasterId,
	}
	err = session.Db.Table("u_member").Where("user_id = ?", otherUserId).OrderBy("member_master_id").Find(&profileInfomation.LoveMembers.Slice)
	utils.CheckErr(err)
	for _, member := range profileInfomation.LoveMembers.Slice {
		profileInfomation.TotalLovePoint += member.LovePoint
	}
	return profileInfomation
}
