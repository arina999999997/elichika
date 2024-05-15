package user_social

import (
	"elichika/enum"
	"elichika/generic"
	"elichika/userdata"
	"elichika/userdata/database"
	"elichika/utils"
)

func GetUserFriendStatus(session *userdata.Session, otherUserId int32) database.UserFriendStatus {
	if otherUserId == session.UserId { // special behaviour to be friend with yourself to use user's own support members
		return database.UserFriendStatus{
			UserId:           session.UserId,
			OtherUserId:      otherUserId,
			FriendApprovedAt: generic.NewNullable(session.UserModel.UserStatus.TutorialEndAt),
			RequestStatus:    enum.FriendRequestStatusFriend,
			IsRequestPending: false,
			IsNew:            false,
		}
	}
	userFriendStatus := database.UserFriendStatus{}
	exist, err := session.Db.Table("u_friend_status").Where("user_id = ? AND other_user_id = ?", session.UserId, otherUserId).
		Get(&userFriendStatus)
	utils.CheckErr(err)
	if exist {
		return userFriendStatus
	} else {
		return database.UserFriendStatus{
			UserId:           session.UserId,
			OtherUserId:      otherUserId,
			RequestStatus:    enum.FriendRequestStatusNone,
			IsRequestPending: false,
			IsNew:            false,
		}
	}

}
