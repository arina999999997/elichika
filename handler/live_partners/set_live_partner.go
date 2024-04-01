package live_partners

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_live_partner"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func setLivePartner(ctx *gin.Context) {
	req := request.SetLivePartnerCardRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	user_live_partner.SetLivePartnerCard(session, req.LivePartnerCategoryId, req.CardMasterId)

	common.JsonResponse(ctx, response.EmptyResponse{})
}

func init() {
	router.AddHandler("/livePartners/setLivePartner", setLivePartner)
}
