package user_social

import (
	"elichika/client/response"
	"elichika/enum"
	"elichika/userdata"
)

// add a friend request from other menus (not friend menu)

func ApplyFriendRequestOtherScene(session *userdata.Session, otherUserId int32) (*response.FriendActionResponse, *response.FriendActionRecoverableExceptionResponse) {
	good, errorKey := ApplyFriendRequestShared(session, otherUserId)
	if good {
		return &response.FriendActionResponse{
			SuccessType:  enum.FriendSuccessTypeNoProblem,
			TargetPlayer: GetNullableOtherUser(session, otherUserId),
		}, nil
	} else {
		return nil, &response.FriendActionRecoverableExceptionResponse{
			ErrorKey:     errorKey,
			TargetPlayer: GetNullableOtherUser(session, otherUserId),
		}
	}
}
