package live

import (
	"elichika/config"
	"elichika/handler"
	"elichika/model"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func LiveUpdatePlayList(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	// fmt.Println(reqBody)
	req := model.LiveUpdatePlayListReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	mul := 0
	if req.IsSet {
		mul = 1
	}
	session.UpdateUserPlayList(model.UserPlayListItem{
		UserID:         userID,
		UserPlayListID: req.GroupNum + req.LiveMasterID*10,
		GroupNum:       req.GroupNum * mul,
		LiveID:         req.LiveMasterID * mul})

	signBody := session.Finalize(handler.GetData("userModelDiff.json"), "user_model_diff")
	signBody, _ = sjson.Set(signBody, "is_success", true)
	fmt.Println(signBody)
	resp := handler.SignResp(ctx, string(signBody), config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
