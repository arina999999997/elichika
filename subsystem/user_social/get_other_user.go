package user_social

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func GetOtherUser(session *userdata.Session, otherUserId int32) client.OtherUser {
	otherUserStatus := GetOtherUserStatus(session, otherUserId)
	userFriendStatus := GetUserFriendStatus(session, otherUserId)
	otherUser := client.OtherUser{
		UserId:                otherUserId,
		Name:                  otherUserStatus.Name,
		Rank:                  otherUserStatus.Rank,
		LastPlayedAt:          otherUserStatus.LastLoginAt, // this might not be correct
		RecommendCardMasterId: otherUserStatus.RecommendCardMasterId,
		EmblemId:              otherUserStatus.EmblemId,
		IsNew:                 userFriendStatus.IsNew,
		IntroductionMessage:   otherUserStatus.Message,
		FriendApprovedAt:      userFriendStatus.FriendApprovedAt,
		RequestStatus:         userFriendStatus.RequestStatus,
		IsRequestPending:      userFriendStatus.IsRequestPending,
	}
	recommendCard := client.UserCard{}
	exist, err := session.Db.Table("u_card").
		Where("user_id = ? AND card_master_id = ?", otherUserId, otherUser.RecommendCardMasterId).Get(&recommendCard)
	utils.CheckErrMustExist(err, exist)
	otherUser.RecommendCardLevel = recommendCard.Level
	otherUser.IsRecommendCardImageAwaken = recommendCard.IsAwakeningImage
	otherUser.IsRecommendCardAllTrainingActivated = recommendCard.IsAllTrainingActivated
	return otherUser
}
