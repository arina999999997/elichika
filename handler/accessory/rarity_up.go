package accessory

import (
	"elichika/client/request"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_accessory"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func rarityUp(ctx *gin.Context) {
	req := request.AccessoryRarityUpRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	resp := user_accessory.RarityUpAccessory(session, req.UserAccessoryId)

	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/", "POST", "/accessory/rarityUp", rarityUp)
}
