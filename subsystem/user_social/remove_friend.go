package user_social

import (
	"elichika/client/response"
	"elichika/enum"
	"elichika/userdata"
)

func RemoveFriend(session *userdata.Session, otherUserId int32) (*response.FriendListResponse, *response.FriendRecoverableExceptionResponse) {
	RemoveConnection(session, otherUserId)
	return &response.FriendListResponse{
		SuccessType:    enum.FriendSuccessTypeNoProblem,
		FriendViewList: GetFriendViewList(session),
	}, nil
}
