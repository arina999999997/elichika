package handler

import (
	"elichika/client"
	"elichika/config"
	"elichika/enum"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func LiveMvStart(ctx *gin.Context) {
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	signBody := session.Finalize("{}", "user_model_diff")
	signBody, _ = sjson.Set(signBody, "uniq_id", session.Time.UnixNano())
	resp := SignResp(ctx, signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func LiveMvSaveDeck(ctx *gin.Context) {
	// TODO: actually save this in db
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()

	type LiveSaveDeckReq struct {
		LiveMasterId        int32   `json:"live_master_id"`
		LiveMvDeckType      int32   `json:"live_mv_deck_type"`
		MemberMasterIdByPos []int32 `json:"member_master_id_by_pos"`
		SuitMasterIdByPos   []int32 `json:"suit_master_id_by_pos"`
		ViewStatusByPos     []int32 `json:"view_status_by_pos"`
	}

	req := LiveSaveDeckReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userLiveMvDeck := client.UserLiveMvDeck{
		LiveMasterId: req.LiveMasterId,
	}
	deckJsonBytes, err := json.Marshal(userLiveMvDeck)
	utils.CheckErr(err)
	deckJson := string(deckJsonBytes)

	for k, v := range req.MemberMasterIdByPos {
		if k%2 == 0 {
			memberId := req.MemberMasterIdByPos[k+1]
			deckJson, err = sjson.Set(deckJson, fmt.Sprintf("member_master_id_%d", v), memberId)
			utils.CheckErr(err)
		}
	}
	for k, v := range req.SuitMasterIdByPos {
		if k%2 == 0 {
			suitId := req.SuitMasterIdByPos[k+1]
			deckJson, err = sjson.Set(deckJson, fmt.Sprintf("suit_master_id_%d", v), suitId)
			utils.CheckErr(err)
		}
	}
	err = json.Unmarshal([]byte(deckJson), &userLiveMvDeck)
	utils.CheckErr(err)
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	for k := range req.ViewStatusByPos {
		if k%2 == 0 {
			memberId := req.MemberMasterIdByPos[k+1]
			// Rina-chan board toggle
			if memberId == enum.MemberMasterIdRina {
				RinaChan := session.GetMember(enum.MemberMasterIdRina)
				RinaChan.ViewStatus = int32(req.ViewStatusByPos[k+1])
				session.UpdateMember(RinaChan)
			}
		}
	}

	var userLiveMvDeckCustomById []any
	userLiveMvDeckCustomById = append(userLiveMvDeckCustomById, req.LiveMasterId)
	userLiveMvDeckCustomById = append(userLiveMvDeckCustomById, userLiveMvDeck)

	signBody := session.Finalize("{}", "user_model")
	if req.LiveMvDeckType == 1 {
		signBody, _ = sjson.Set(signBody, "user_model.user_live_mv_deck_by_id", userLiveMvDeckCustomById)
	} else {
		signBody, _ = sjson.Set(signBody, "user_model.user_live_mv_deck_custom_by_id", userLiveMvDeckCustomById)
	}

	resp := SignResp(ctx, signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
