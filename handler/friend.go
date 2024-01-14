package handler

import (
	"elichika/client/response"
	"elichika/enum"
	// "elichika/userdata"

	"github.com/gin-gonic/gin"
)

// TODO(friend): Implement friend system
func FetchFriendList(ctx *gin.Context) {
	// there's no request body

	// userId := ctx.GetInt("user_id")
	// session := userdata.GetSession(ctx, userId)
	// defer session.Close()

	JsonResponse(ctx, response.FriendListResponse{
		SuccessType: enum.FriendSuccessTypeNoProblem, // not sure why this value is necessary but ok
	})
}
