package friend

import (
	"elichika/client/response"
	"elichika/enum"
	"elichika/handler/common"
	"elichika/router"
	// "elichika/userdata"

	"github.com/gin-gonic/gin"
)

// TODO(friend): Implement friend system
func fetchFriendList(ctx *gin.Context) {
	// there's no request body

	// userId := int32(ctx.GetInt("user_id"))
	// session := userdata.GetSession(ctx, userId)
	// defer session.Close()

	common.JsonResponse(ctx, response.FriendListResponse{
		SuccessType: enum.FriendSuccessTypeNoProblem, // not sure why this value is necessary but ok
	})
}

func init() {
	router.AddHandler("/friend/fetchFriendList", fetchFriendList)
}
