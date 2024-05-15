package user_social

import (
	"elichika/client/response"
	"elichika/enum"
	"elichika/userdata"
)

// refuse friend requests
func RefuseFriendRequest(session *userdata.Session, otherUserIds []int32, isMass bool) (*response.FriendListResponse, *response.FriendRecoverableExceptionResponse) {
	for _, userId := range otherUserIds {
		RemoveConnection(session, userId)
	}
	return &response.FriendListResponse{
		SuccessType:    enum.FriendSuccessTypeNoProblem,
		FriendViewList: GetFriendViewList(session),
	}, nil
}
