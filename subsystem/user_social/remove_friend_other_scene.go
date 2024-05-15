package user_social

import (
	"elichika/client/response"
	"elichika/enum"
	"elichika/userdata"
)

func RemoveFriendOtherScene(session *userdata.Session, otherUserId int32) (*response.FriendActionResponse, *response.FriendActionRecoverableExceptionResponse) {
	RemoveConnection(session, otherUserId)
	return &response.FriendActionResponse{
		SuccessType:  enum.FriendSuccessTypeNoProblem,
		TargetPlayer: GetNullableOtherUser(session, otherUserId),
	}, nil
}
