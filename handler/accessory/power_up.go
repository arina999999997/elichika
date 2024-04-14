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

func powerUp(ctx *gin.Context) {
	req := request.AccessoryPowerUpRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	resp := user_accessory.LevelUpAccessory(session, req.UserAccessoryId, req.PowerUpAccessoryIds, req.AccessoryLevelUpItems)

	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/", "POST", "/accessory/powerUp", powerUp)
}
