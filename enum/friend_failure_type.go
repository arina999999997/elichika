package enum

const (
	FriendFailureTypeRequestMaxFriend            int32 = 0x00000001 // sending request from other scene and self has max friend
	FriendFailureTypeRequestMaxTargetFriend      int32 = 0x00000002 // sending request from other scene and other has max friend
	FriendFailureTypeRequestMaxApplication       int32 = 0x00000003 // sending request and self has too many out going request
	FriendFailureTypeRequestMaxTargetApplication int32 = 0x00000004 // sending request and other has too many in coming request
	FriendFailureTypeRequestAlreadyFriend        int32 = 0x00000005 // already friend
	FriendFailureTypeApproveMaxFriend            int32 = 0x00000006 // approving request and self has max friend
	FriendFailureTypeApproveNotExist             int32 = 0x00000007 // friend request cancelled by other side
	FriendFailureTypeApproveExpired              int32 = 0x00000008 // friend request expired
	FriendFailureTypeApproveMaxTargetFriend      int32 = 0x00000009 // approving request and other has max friend
	FriendFailureTypeCancelAlreadyFriend         int32 = 0x0000000a // canceling request but the other has already accepted it
	FriendFailureTypeSearchNotExist              int32 = 0x0000000b // other not exist when searching id
	FriendFailureTypeCancelTargetNotExistInMass  int32 = 0x0000000c // canceling multiple outgoing request but there's some problem
	FriendFailureTypeApproveTargetNotExistInMass int32 = 0x0000000d // accepting multiple incoming request but there's some problem
	FriendFailureTypeUserAccountDeleted          int32 = 0x0000000e // same as FriendFailureTypeSearchNotExist
)
