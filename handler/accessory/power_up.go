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

	// limit break (grade up) is processed first, then exp is processed later
	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	resp := user_accessory.LevelUpAccessory(session, req.UserAccessoryId, req.PowerUpAccessoryIds, req.AccessoryLevelUpItems)

	session.Finalize()
	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/accessory/powerUp", powerUp)
}
