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

func activateTrainingTreeCell(ctx *gin.Context) {
	req := request.ActivateTrainingTreeCellRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	user_training_tree.ActivateTrainingTreeCells(session, req.CardMasterId, req.CellMasterIds.Slice)

	session.Finalize()
	common.JsonResponse(ctx, &response.ActivateTrainingTreeCellResponse{
		UserCardTrainingTreeCellList: user_training_tree.GetUserTrainingTree(session, req.CardMasterId),
		UserModelDiff:                &session.UserModel,
	})
}

func init() {
	router.AddHandler("/", "POST", "/trainingTree/activateTrainingTreeCell", activateTrainingTreeCell)
}
