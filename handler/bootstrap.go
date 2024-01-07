package handler

import (
	"elichika/client"
	"elichika/config"
	"elichika/enum"
	"elichika/protocol/request"
	"elichika/protocol/response"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"net/http"

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

	respObj := response.FetchBootstrapResponse{
		UserModelDiff: &session.UserModel,
		UserInfoTrigger: response.UserInfoTrigger{
			UserInfoTriggerGachaPointExchangeRows:           []any{},
			UserInfoTriggerExpiredGiftBoxRows:               []any{},
			UserInfoTriggerEventMarathonShowResultRows:      []any{},
			UserInfoTriggerEventMiningShowResultRows:        []any{},
			UserInfoTriggerEventCoopShowResultRows:          []any{},
			UserInfoTriggerSubscriptionTrialEndRows:         []any{},
			UserInfoTriggerSubscriptionEndRows:              []any{},
			UserInfoTriggerMemberGuildRankingShowResultRows: []any{},
		},
		BillingStateInfo: response.BillingStateInfo{
			Age:                        42,
			CurrentMonthPurcharsePrice: 0,
		},
		FetchBootstrapBannerResponse: response.BootstrapBanner{
			Banners: []response.Banner{},
		},
		FetchBootstrapNewBadgeResponse: response.BootstrapNewBadge{
			IsNewMainStory:                     false,
			UnreceivedPresentBox:               0,
			IsUnreceivedPresentBoxSubscription: false,
			NoticeNewArrivalsIds:               []int{},
			IsUpdateFriend:                     false,
			UnreceivedMission:                  0,
			UnreceivedChallengeBeginner:        0,
			DailyTheaterTodayId:                1001243,
		},
		FetchBootstrapPickupInfoResponse: response.BootstrapPickupInfo{
			ActiveEvent: nil,
			LiveCampaignInfo: response.LiveCampaignInfo{
				LiveCampaignEndAt:          nil,
				LiveDailyCampaignEndAt:     nil,
				LiveExtraCampaignEndAt:     nil,
				LiveChallengeCampaignEndAt: nil,
			},
			IsLessonCampaign: false,
			AppealGachas:     []client.TextureStruktur{},
			IsShopSale:       false,
			IsSnsCoinSale:    false,
		},
		FetchBootstrapExpiredItemResponse: response.BootstrapExpiredItem{
			ExpiredItems: []client.Content{},
		},
		FetchBootstrapLoginBonusResponse: client.BootstrapLoginBonus{},
		FetchBootstrapNoticeResponse: response.BootstrapNotice{
			SuperNotices:        []any{},
			FetchedAt:           1688014785,
			ReviewSuperNoticeAt: 2019600000,
			ForceViewNoticeIds:  []any{},
		},
		FetchBootstrapSubscriptionResponse: response.BootstrapSubscription{
			ContinueRewards: []any{},
		},
		MissionBeginnerMasterId:       nil,
		ShowChallengeBeginnerButton:   false,
		ChallengeBeginnerCompletedIds: []int{1, 2, 3, 4, 5, 6},
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
			respObj.FetchBootstrapLoginBonusResponse = GetBootstrapLoginBonus(ctx, session)
		case enum.BootstrapFetchTypeNotice:
			continue
		case enum.BootstrapFetchTypeSchoolIdolFestivalIdReward:
			continue
		default:
			panic("unexpected type")
		}
	}

	session.UserModel.UserSubscriptionStatusById.PushBack(session.GetSubsriptionStatus())
	session.Finalize("{}", "dummy")

	respBytes, _ := json.Marshal(respObj)
	resp := SignResp(ctx, string(respBytes), config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func GetClearedPlatformAchievement(ctx *gin.Context) {
	signBody := GetData("getClearedPlatformAchievement.json")
	resp := SignResp(ctx, signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
