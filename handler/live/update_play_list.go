package live

import (
	"elichika/config"
	"elichika/handler"
	"elichika/model"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func LiveUpdatePlayList(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := model.LiveUpdatePlayListReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	session.UpdateUserPlayList(model.UserPlayListItem{
		UserId:         userId,
		UserPlayListId: req.GroupNum + req.LiveMasterId*10,
		GroupNum:       req.GroupNum,
		LiveId:         req.LiveMasterId,
		IsNull:         !req.IsSet,
	})

	signBody := session.Finalize("{}", "user_model_diff")
	signBody, _ = sjson.Set(signBody, "is_success", true)
	resp := handler.SignResp(ctx, string(signBody), config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
