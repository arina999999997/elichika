package user_social

import (
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/userdata"
)

// approve a friend request
// return ok, error key (if not ok)
func ApproveFriendRequest(session *userdata.Session, otherUserIds []int32, isMass bool) (*response.FriendListResponse, *response.FriendRecoverableExceptionResponse) {
	resp := response.FriendListResponse{
		SuccessType: enum.FriendSuccessTypeNoProblem,
	}
	for _, userId := range otherUserIds {
		good, errorKey := ApproveFriendRequestShared(session, userId)
		if !good {
			if isMass {
				// if mass then some request can be accepted while others "rejected" due to maxing out from our or their side
				if (errorKey == enum.FriendFailureTypeApproveMaxFriend) || (errorKey == enum.FriendFailureTypeApproveMaxTargetFriend) {
					resp.SuccessType = enum.FriendSuccessTypeMaxFriend
				} else {
					return nil, &response.FriendRecoverableExceptionResponse{
						ErrorKey:       enum.FriendFailureTypeApproveTargetNotExistInMass,
						FriendViewList: generic.NewNullable(GetFriendViewList(session)),
					}
				}
			} else {
				return nil, &response.FriendRecoverableExceptionResponse{
					ErrorKey:       errorKey,
					FriendViewList: generic.NewNullable(GetFriendViewList(session)),
				}
			}
		}
	}
	resp.FriendViewList = GetFriendViewList(session)
	return &resp, nil
}
