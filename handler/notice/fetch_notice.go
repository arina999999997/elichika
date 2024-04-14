package notice

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"

	"github.com/gin-gonic/gin"
)

// TODO(notice): This is not implemented although it could be cool to use it to put guide or stuff
func fetchNotice(ctx *gin.Context) {
	// there is no request body

	resp := response.FetchNoticeResponse{
		NoticeNoCheckAt: 2019600000, // this is used to check if news are already displayed for today
	}
	for i := int32(1); i <= 5; i++ {
		resp.NoticeLists.Set(i, client.NoticeList{
			Category: i,
		})
	}
	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/", "POST", "/notice/fetchNotice", fetchNotice)
}
