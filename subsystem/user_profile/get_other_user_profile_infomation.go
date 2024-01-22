package user_profile

import (
	"elichika/client"
	"elichika/subsystem/user_card"
	"elichika/userdata"
	"elichika/utils"
)

func GetOtherUserProfileInfomation(session *userdata.Session, otherUserId int32) client.ProfileInfomation {
	// TODO(friend): Actually calculate the friend links
	otherUserStatus := client.UserStatus{}
	exist, err := session.Db.Table("u_status").Where("user_id = ?", otherUserId).Get(&otherUserStatus)
	utils.CheckErrMustExist(err, exist)
	otherUserCard := user_card.GetOtherUserCard(session, otherUserId, otherUserStatus.RecommendCardMasterId)
	otherUser := client.OtherUser{
		UserId:                              otherUserId,
		Name:                                otherUserStatus.Name,
		Rank:                                otherUserStatus.Rank,
		LastPlayedAt:                        otherUserStatus.LastLoginAt, // this might not be correct
		RecommendCardMasterId:               otherUserStatus.RecommendCardMasterId,
		RecommendCardLevel:                  otherUserCard.Level,
		IsRecommendCardImageAwaken:          otherUserCard.IsAwakeningImage,
		IsRecommendCardAllTrainingActivated: otherUserCard.IsAllTrainingActivated,
		EmblemId:                            otherUserStatus.EmblemId,
		// IsNew: otherUserStatus.IsNew,
		IntroductionMessage: otherUserStatus.Message,
		// FriendApprovedAt: otherUserStatus.FriendApprovedAt,
		// RequestStatus: otherUserStatus.RequestStatus,
		// IsRequestPending: otherUserStatus.IsRequestPending,
	}
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
