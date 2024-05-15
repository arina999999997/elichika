package friend

import (
	"elichika/client/response"
	"elichika/enum"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_social"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

func fetchFriendList(ctx *gin.Context) {
	session := ctx.MustGet("session").(*userdata.Session)

	common.JsonResponse(ctx, response.FriendListResponse{
		SuccessType:    enum.FriendSuccessTypeNoProblem,
		FriendViewList: user_social.GetFriendViewList(session),
	})
}

func init() {
	router.AddHandler("/", "POST", "/friend/fetchFriendList", fetchFriendList)
}
