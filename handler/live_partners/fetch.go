package live_partners

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_live_partner"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

func fetch(ctx *gin.Context) {
	// a set of partners player (i.e. friends and others), then fetch the card for them
	// this set include the current user, so we can use our own cards.
	// currently only have current user
	// note that all card are available, but we need to use the filter functionality in the client to actually get them to show up.

	resp := response.FetchLiveParntersResponse{}

	// there is no request body
	session := ctx.MustGet("session").(*userdata.Session)

	partnerUserIds := []int32{} // TODO(friend): Fill this with some users
	partnerUserIds = append(partnerUserIds, session.UserId)

	for _, partnerId := range partnerUserIds {
		resp.PartnerSelectState.LivePartners.Append(user_live_partner.GetOtherUserLivePartner(session, partnerId))
	}
	resp.PartnerSelectState.FriendCount = int32(resp.PartnerSelectState.LivePartners.Size())
	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/livePartners/fetch", fetch)
}
