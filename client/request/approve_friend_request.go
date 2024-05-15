package request

import (
	"elichika/generic"
)

type ApproveFriendRequest struct {
	UserIds generic.Array[int32] `json:"user_ids"`
	IsMass  bool                 `json:"is_mass"`
}
