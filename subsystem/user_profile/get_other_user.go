package user_profile

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func GetOtherUser(session *userdata.Session, otherUserId int32) client.OtherUser {
	otherUserStatus := client.UserStatus{}
	userdata.FetchDBProfile(otherUserId, &otherUserStatus)
	otherUser := client.OtherUser{
		UserId:                otherUserId,
		Name:                  otherUserStatus.Name,
		Rank:                  otherUserStatus.Rank,
		LastPlayedAt:          otherUserStatus.LastLoginAt,
		RecommendCardMasterId: otherUserStatus.RecommendCardMasterId,
		EmblemId:              otherUserStatus.EmblemId,
		// IsNew: otherUserStatus.IsNew,
		IntroductionMessage: otherUserStatus.Message,
		// FriendApprovedAt: otherUserStatus.FriendApprovedAt,
		// RequestStatus: otherUserStatus.RequestStatus,
		// IsRequestPending: otherUserStatus.IsRequestPending,
	}
	recommendCard := client.UserCard{}
	exist, err := session.Db.Table("u_card").
		Where("user_id = ? AND card_master_id = ?", otherUserId, otherUser.RecommendCardMasterId).Get(&recommendCard)
	utils.CheckErr(err)
	if !exist {
		panic("other user card doesn't exist")
	}

	otherUser.RecommendCardLevel = recommendCard.Level
	otherUser.IsRecommendCardImageAwaken = recommendCard.IsAwakeningImage
	otherUser.IsRecommendCardAllTrainingActivated = recommendCard.IsAllTrainingActivated

	// TODO(friend): not implemented
	// otherUser.FriendApprovedAt = new(int64)
	// *otherUser.FriendApprovedAt = 0
	// otherUser.RequestStatus = 3
	// otherUser.IsRequestPending = false
	return otherUser
}
