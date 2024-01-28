package live_mv

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_member"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func LiveMvStart(ctx *gin.Context) {
	// we don't really need the request
	// maybe it's once needed or it's only used for gathering data
	// reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	// req := request.StartLiveMvRequest{}
	// err := json.Unmarshal([]byte(reqBody), &req)
	// utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	common.JsonResponse(ctx, &response.StartLiveMvResponse{
		UniqId:        session.Time.UnixNano(),
		UserModelDiff: &session.UserModel,
	})
}

func LiveMvSaveDeck(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SaveLiveMvDeckRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	userLiveMvDeck := client.UserLiveMvDeck{
		LiveMasterId: req.LiveMasterId,
	}

	for pos, memberMasterId := range req.MemberMasterIdByPos.Map {
		reflect.ValueOf(&userLiveMvDeck).Elem().Field(int(pos)).Set(reflect.ValueOf(generic.NewNullable(*memberMasterId)))
	}
	for pos, suitMasterId := range req.SuitMasterIdByPos.Map {
		reflect.ValueOf(&userLiveMvDeck).Elem().Field(12 + int(pos)).Set(reflect.ValueOf(*suitMasterId))
	}
	for pos, viewStatus := range req.ViewStatusByPos.Map {
		memberId := req.MemberMasterIdByPos.GetOnly(pos)
		if *memberId == enum.MemberMasterIdRina {
			RinaChan := user_member.GetMember(session, enum.MemberMasterIdRina)
			RinaChan.ViewStatus = *viewStatus
			user_member.UpdateMember(session, RinaChan)
		}
	}

	if req.LiveMvDeckType == enum.LiveMvDeckTypeOriginal {
		session.UserModel.UserLiveMvDeckById.Set(req.LiveMasterId, userLiveMvDeck)
	} else {
		session.UserModel.UserLiveMvDeckCustomById.Set(req.LiveMasterId, userLiveMvDeck)
	}

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {

	router.AddHandler("/liveMv/saveDeck", LiveMvSaveDeck)
	router.AddHandler("/liveMv/start", LiveMvStart)
}
