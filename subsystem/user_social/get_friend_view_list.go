package user_social

import (
	"elichika/client"
	"elichika/enum"
	"elichika/userdata"
	"elichika/userdata/database"
	"elichika/utils"
)

// fetch the friend list, but also mark everyone as not new
func GetFriendViewList(session *userdata.Session) client.FriendViewList {
	view := client.FriendViewList{}
	friends := []database.UserFriendStatus{}
	err := session.Db.Table("u_friend_status").Where("user_id = ?", session.UserId).Find(&friends)
	utils.CheckErr(err)
	for _, friendStatus := range friends {
		if friendStatus.RequestStatus == enum.FriendRequestStatusFriend {
			view.FriendList.Append(GetOtherUser(session, friendStatus.OtherUserId))
		} else if friendStatus.IsRequestPending {
			view.FriendApprovalPendingList.Append(GetOtherUser(session, friendStatus.OtherUserId))
		} else {
			view.FriendApplicationList.Append(GetOtherUser(session, friendStatus.OtherUserId))
		}
		if friendStatus.IsNew {
			friendStatus.IsNew = false
			UpdateUserFriendStatus(session, friendStatus)
		}
	}
	recommend := GetRecommendedUserIds(session)
	for _, otherUserId := range recommend {
		view.RecomendPlayerList.Append(GetOtherUser(session, otherUserId))
	}
	return view
}
