package handler

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
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

	userId := ctx.GetInt("user_id")
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
			resp.FetchBootstrapLoginBonusResponse = GetBootstrapLoginBonus(ctx, session)
		case enum.BootstrapFetchTypeNotice:
			continue
		case enum.BootstrapFetchTypeSchoolIdolFestivalIdReward:
			continue
		default:
			panic("unexpected type")
		}
	}
	status := session.GetSubsriptionStatus()
	session.UserModel.UserSubscriptionStatusById.Set(status.SubscriptionMasterId, status)
	session.Finalize("{}", "dummy")
	JsonResponse(ctx, resp)
}

func GetClearedPlatformAchievement(ctx *gin.Context) {
	JsonResponse(ctx, &response.GetClearedPlatformAchievementResponse{})
}
