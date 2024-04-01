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
)

func fetchTrainingTree(ctx *gin.Context) {
	req := request.FetchTrainingTreeRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
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
