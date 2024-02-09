package training_tree

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_training_tree"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func levelUpCard(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.LevelUpCardRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	user_training_tree.LevelUpCard(session, req.CardMasterId, req.AdditionalLevel)

	session.Finalize()
	common.JsonResponse(ctx, response.LevelUpCardResponse{
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/trainingTree/levelUpCard", levelUpCard)
}
