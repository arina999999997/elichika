package user_bootstrap

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/subsystem/banner"
	"elichika/subsystem/user_login_bonus"
	"elichika/subsystem/pickup_info"
	"elichika/subsystem/user_beginner_challenge"
	"elichika/subsystem/user_expired_item"
	"elichika/subsystem/user_new_badge"
	"elichika/userdata"
)

func FetchBootstrap(session *userdata.Session, req request.FetchBootstrapRequest) response.FetchBootstrapResponse {
	// TODO(bootstrap, authentication): Log the user out if they have differnt device token / name
	session.UserStatus.BootstrapSifidCheckAt = session.Time.UnixMilli()
	session.UserStatus.DeviceToken = req.DeviceToken
	resp := response.FetchBootstrapResponse{
		UserModelDiff: &session.UserModel,
	}

	for _, fetchType := range req.BootstrapFetchTypes.Slice {
		switch fetchType {
		case enum.BootstrapFetchTypeBanner:
			resp.FetchBootstrapBannerResponse = banner.GetBootstrapBannerResponse(session)
		case enum.BootstrapFetchTypeNewBadge:
			resp.FetchBootstrapNewBadgeResponse = generic.NewNullable(user_new_badge.GetBootstrapNewBadgeResponse(session))
		case enum.BootstrapFetchTypePickupInfo:
			resp.FetchBootstrapPickupInfoResponse = pickup_info.GetBootstrapPickupInfo(session)
		case enum.BootstrapFetchTypeExpiredItem:
			resp.FetchBootstrapExpiredItemResponse = user_expired_item.GetBootstrapExpiredItem(session)
		case enum.BootstrapFetchTypeGachaPointExchange:
			// TODO(gacha): Implement gacha point
			continue
		case enum.BootstrapFetchTypeLoginBonus:
			resp.FetchBootstrapLoginBonusResponse = generic.NewNullable(user_login_bonus.GetBootstrapLoginBonus(session))
		case enum.BootstrapFetchTypeNotice:
			// TODO(notice): Implement notice
			continue
		case enum.BootstrapFetchTypeSchoolIdolFestivalIdReward:
			// this is no longer used, the client can't comprehend it unless we change stuff
			continue
		default:
			panic("unexpected type")
		}
	}
	resp.MissionBeginnerMasterId, resp.ShowChallengeBeginnerButton, resp.ChallengeBeginnerCompletedIds =
		user_beginner_challenge.GetUserBeginnerChallengeInfo(session)
	status := session.GetSubsriptionStatus(13001)
	session.UserModel.UserSubscriptionStatusById.Set(status.SubscriptionMasterId, status)
	return resp
}
