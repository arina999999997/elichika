package live_partners

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_live"
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
	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	partnerUserIds := []int32{} // TODO(friend): Fill this with some users
	partnerUserIds = append(partnerUserIds, userId)

	for _, partnerId := range partnerUserIds {
		resp.PartnerSelectState.LivePartners.Append(user_live.GetLivePartner(session, partnerId))
	}
	resp.PartnerSelectState.FriendCount = int32(resp.PartnerSelectState.LivePartners.Size())
	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/livePartners/fetch", fetch)
}
