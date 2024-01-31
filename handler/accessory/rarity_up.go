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
	"github.com/tidwall/gjson"
)

func rarityUp(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.AccessoryRarityUpRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	resp := user_accessory.RarityUpAccessory(session, req.UserAccessoryId)

	session.Finalize()
	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/accessory/rarityUp", rarityUp)
}