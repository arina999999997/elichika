package communication_member

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/gamedata"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_live_difficulty"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func updateUserLiveDifficultyNewFlag(ctx *gin.Context) {
	// mark all the song that this member is featured in as not new
	// only choose from the song user has access to, so no bond song and story locked songs
	// this use the same request as the above, must have been coded around the same time
	req := request.UpdateMemberDetailBadgeRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	liveDifficultyRecords := user_live_difficulty.GetAllLiveDifficulties(session)
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)

	for _, liveDifficultyRecord := range liveDifficultyRecords {
		if !liveDifficultyRecord.IsNew { // no need to update
			continue
		}
		// update if it feature this member
		liveDifficultyMaster := gamedata.LiveDifficulty[liveDifficultyRecord.LiveDifficultyId]
		if liveDifficultyMaster == nil {
			// some song no longer exists but official server still send them
			// it's ok to ignore these for now
			continue
		}
		_, exist := liveDifficultyMaster.Live.LiveMemberMapping[req.MemberMasterId]
		if exist {
			liveDifficultyRecord.IsNew = false
			user_live_difficulty.UpdateLiveDifficulty(session, liveDifficultyRecord)
		}
	}

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/communicationMember/updateUserLiveDifficultyNewFlag", updateUserLiveDifficultyNewFlag)
}
