package unlock_scene

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/handler/common"
	"elichika/router"
	"elichika/userdata"
	"elichika/utils"
	"elichika/subsystem/user_unlock_scene"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func saveUnlockedScene(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SaveUnlockedSceneRequest1{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	for _, sceneType := range req.UnlockSceneTypes.Slice {
		user_unlock_scene.UnlockScene(session, sceneType, enum.UnlockSceneStatusOpened)
	}

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/unlockScene/saveUnlockedScene", saveUnlockedScene)
}
