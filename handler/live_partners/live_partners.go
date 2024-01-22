package live_partners

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func FetchLivePartners(ctx *gin.Context) {
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
	partnerUserIds = append(partnerUserIds, int32(userId))

	for _, partnerId := range partnerUserIds {

		partner := client.LivePartner{}
		userdata.FetchDBProfile(partnerId, &partner)

		partner.IsFriend = true
		partnerCards := userdata.FetchPartnerCards(partnerId) // client.UserCard
		if len(partnerCards) == 0 {
			continue
		}
		for _, card := range partnerCards {
			for i := 1; i <= 7; i++ {
				if (card.LivePartnerCategories & (1 << i)) != 0 {
					partner.CardByCategory.Set(int32(i), session.GetOtherUserCard(partnerId, card.CardMasterId))
				}
			}
		}
		resp.PartnerSelectState.LivePartners.Append(partner)
	}
	resp.PartnerSelectState.FriendCount = int32(resp.PartnerSelectState.LivePartners.Size())
	common.JsonResponse(ctx, &resp)
}

func SetLivePartner(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SetLivePartnerCardRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	// set the bit on the correct card
	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	newCard := session.GetUserCard(req.CardMasterId)
	newCard.LivePartnerCategories |= (1 << req.LivePartnerCategoryId)
	session.UpdateUserCard(newCard)

	// TODO(refactor): Get rid of this bit hacking stuff
	// remove the bit on the other cards
	partnerCards := userdata.FetchPartnerCards(userId)
	for _, card := range partnerCards {
		if card.CardMasterId == req.CardMasterId {
			continue
		}
		if (card.LivePartnerCategories & (1 << req.LivePartnerCategoryId)) != 0 {
			card.LivePartnerCategories ^= (1 << req.LivePartnerCategoryId)
			session.UpdateUserCard(card)
		}
	}

	session.Finalize()
	common.JsonResponse(ctx, response.EmptyResponse{})
}

func init() {
	router.AddHandler("/livePartners/fetch", FetchLivePartners)
	router.AddHandler("/livePartners/setLivePartner", SetLivePartner)
}
