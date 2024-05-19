package user_social

import (
	"elichika/client/response"
	"elichika/userdata"
)

const (
	LivePartnerFriendLimit int32 = 50
)

func GetLivePartners(session *userdata.Session) response.FetchLiveParntersResponse {
	// a set of partners player (i.e. friends and others), then fetch the card for them
	// this set include the current user, so we can use our own cards.
	// note that all card are available, but we need to use the filter functionality in the client to actually get them to show up.
	// official server return all the friend, at least once you have enough friends
	// we will not handle
	resp := response.FetchLiveParntersResponse{}

	partnerUserIds := GetFriendUserIds(session)
	partnerUserIds = append(partnerUserIds, session.UserId)

	resp.PartnerSelectState.FriendCount = int32(len(partnerUserIds))
	if resp.PartnerSelectState.FriendCount < LivePartnerFriendLimit {
		// add random people if not enough friend
		partnerUserIds = append(partnerUserIds, GetRecommendedUserIds(session)...)
	}
	for _, partnerId := range partnerUserIds {
		resp.PartnerSelectState.LivePartners.Append(GetLivePartner(session, partnerId))
	}
	return resp
}
