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

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	user_live_partner.SetLivePartnerCard(session, req.LivePartnerCategoryId, req.CardMasterId)

	session.Finalize()
	common.JsonResponse(ctx, response.EmptyResponse{})
}

func init() {
	router.AddHandler("/livePartners/setLivePartner", setLivePartner)
}
