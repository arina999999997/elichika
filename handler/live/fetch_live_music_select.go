package live

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/gamedata"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/router"
	"elichika/userdata"

	"time"

	"github.com/gin-gonic/gin"
)

func fetchLiveMusicSelect(ctx *gin.Context) {
	// ther is no request body

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	now := time.Now()
	year, month, day := now.Year(), now.Month(), now.Day()
	tomorrow := time.Date(year, month, day+1, 0, 0, 0, 0, now.Location()).Unix()

	weekday := int32(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)

	resp := response.FetchLiveMusicSelectResponse{
		WeekdayState: client.WeekdayState{
			Weekday:       weekday,
			NextWeekdayAt: tomorrow,
		},
		UserModelDiff: &session.UserModel,
	}
	for _, liveDaily := range gamedata.LiveDaily {
		if liveDaily.Weekday != weekday {
			continue
		}
		resp.LiveDailyList.Append(client.LiveDaily{
			LiveDailyMasterId:      liveDaily.Id,
			LiveMasterId:           liveDaily.LiveId,
			EndAt:                  tomorrow,
			RemainingPlayCount:     5, // this is not kept track of
			RemainingRecoveryCount: generic.NewNullable(int32(10)),
		})
	}

	session.Finalize()
	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/live/fetchLiveMusicSelect", fetchLiveMusicSelect)
}
