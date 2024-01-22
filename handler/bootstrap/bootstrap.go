package bootstrap

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/handler/login_bonus"
	"elichika/router"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func FetchBootstrap(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FetchBootstrapRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	session.UserStatus.BootstrapSifidCheckAt = session.Time.UnixMilli()
	session.UserStatus.DeviceToken = req.DeviceToken

	resp := response.FetchBootstrapResponse{
		UserModelDiff: &session.UserModel,
		FetchBootstrapNewBadgeResponse: client.BootstrapNewBadge{
			DailyTheaterTodayId: generic.NewNullable(int32(1001243)),
		},
		FetchBootstrapNoticeResponse: client.BootstrapNotice{
			FetchedAt:           1688014785,
			ReviewSuperNoticeAt: 2019600000,
		},
		ChallengeBeginnerCompletedIds: generic.List[int32]{
			Slice: []int32{1, 2, 3, 4, 5, 6},
		},
	}
	for _, fetchType := range req.BootstrapFetchTypes {
		switch fetchType {
		case enum.BootstrapFetchTypeBanner:
			continue
		case enum.BootstrapFetchTypeNewBadge:
			continue
		case enum.BootstrapFetchTypePickupInfo:
			continue
		case enum.BootstrapFetchTypeExpiredItem:
			continue
		case enum.BootstrapFetchTypeGachaPointExchange:
			continue
		case enum.BootstrapFetchTypeLoginBonus:
			resp.FetchBootstrapLoginBonusResponse = login_bonus.GetBootstrapLoginBonus(ctx, session)
		case enum.BootstrapFetchTypeNotice:
			continue
		case enum.BootstrapFetchTypeSchoolIdolFestivalIdReward:
			continue
		default:
			panic("unexpected type")
		}
	}
	status := session.GetSubsriptionStatus(13001)
	session.UserModel.UserSubscriptionStatusById.Set(status.SubscriptionMasterId, status)
	session.Finalize()
	common.JsonResponse(ctx, resp)
}

func GetClearedPlatformAchievement(ctx *gin.Context) {
	common.JsonResponse(ctx, &response.GetClearedPlatformAchievementResponse{})
}

func init() {

	router.AddHandler("/bootstrap/fetchBootstrap", FetchBootstrap)
	router.AddHandler("/bootstrap/getClearedPlatformAchievement", GetClearedPlatformAchievement)

}
