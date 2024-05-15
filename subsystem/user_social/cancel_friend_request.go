package user_social

import (
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/userdata"
)

// cancel a friend request
// return ok, error key (if not ok)
func CancelFriendRequest(session *userdata.Session, otherUserIds []int32, isMass bool) (*response.FriendListResponse, *response.FriendRecoverableExceptionResponse) {
	for _, userId := range otherUserIds {
		good, errorKey := CancelFriendRequestShared(session, userId)
		if !good {
			if isMass {
				return nil, &response.FriendRecoverableExceptionResponse{
					ErrorKey:       enum.FriendFailureTypeCancelTargetNotExistInMass,
					FriendViewList: generic.NewNullable(GetFriendViewList(session)),
				}
			} else {
				return nil, &response.FriendRecoverableExceptionResponse{
					ErrorKey:       errorKey,
					FriendViewList: generic.NewNullable(GetFriendViewList(session)),
				}
			}
		}
	}
	return &response.FriendListResponse{
		SuccessType:    enum.FriendSuccessTypeNoProblem,
		FriendViewList: GetFriendViewList(session),
	}, nil
}
