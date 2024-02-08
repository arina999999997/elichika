package training_tree

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/subsystem/user_training_tree"
	"elichika/router"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func fetchTrainingTree(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FetchTrainingTreeRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	common.JsonResponse(ctx, response.FetchTrainingTreeResponse{
		UserCardTrainingTreeCellList: user_training_tree.GetUserTrainingTree(session, req.CardMasterId),
	})
}

func init() {
	router.AddHandler("/trainingTree/fetchTrainingTree", fetchTrainingTree)
}
