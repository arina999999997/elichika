package handler

import (
	"elichika/config"
	"elichika/model"
	"elichika/userdata"
	"elichika/utils"
	"fmt"
	"net/http"
	"strings"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func ExecuteLesson(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type ExecuteLessonReq struct {
		ExecuteLessonIds   []int `json:"execute_lesson_ids"`
		ConsumedContentIds []int `json:"consumed_content_ids"`
		SelectedDeckId     int32 `json:"selected_deck_id"`
		IsThreeTimes       bool  `json:"is_three_times"`
	}
	req := ExecuteLessonReq{}
	if err := json.Unmarshal([]byte(reqBody), &req); err != nil {
		panic(err)
	}

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	deckBytes, _ := json.Marshal(session.GetUserLessonDeck(req.SelectedDeckId))
	deckInfo := string(deckBytes)
	var actionList []model.LessonMenuAction

	gjson.Parse(deckInfo).ForEach(func(key, value gjson.Result) bool {
		if strings.Contains(key.String(), "card_master_id") {
			actionList = append(actionList, model.LessonMenuAction{
				CardMasterId:                  value.Int(),
				Position:                      0,
				IsAddedPassiveSkill:           true,
				IsAddedSpecialPassiveSkill:    true,
				IsRankupedPassiveSkill:        true,
				IsRankupedSpecialPassiveSkill: true,
				IsPromotedSkill:               true,
				MaxRarity:                     4,
				UpCount:                       1,
			})
		}
		return true
	})

	session.UserStatus.MainLessonDeckId = int32(req.SelectedDeckId)
	signBody := session.Finalize(GetData("executeLesson.json"), "user_model_diff")
	signBody, _ = sjson.Set(signBody, "lesson_menu_actions.1", actionList)
	signBody, _ = sjson.Set(signBody, "lesson_menu_actions.3", actionList)
	signBody, _ = sjson.Set(signBody, "lesson_menu_actions.5", actionList)
	signBody, _ = sjson.Set(signBody, "lesson_menu_actions.7", actionList)

	resp := SignResp(ctx, signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func ResultLesson(ctx *gin.Context) {
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	signBody := session.Finalize(GetData("resultLesson.json"), "user_model_diff")
	signBody, _ = sjson.Set(signBody, "selected_deck_id", session.UserStatus.MainLessonDeckId)
	resp := SignResp(ctx, signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func SkillEditResult(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")

	req := gjson.Parse(reqBody).Array()[0]

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	skillList := req.Get("selected_skill_ids")
	skillList.ForEach(func(key, cardId gjson.Result) bool {
		if key.Int()%2 == 0 {
			userCard := session.GetUserCard(int32(cardId.Int()))
			skills := skillList.Get(fmt.Sprintf("%d", key.Int()+1))
			cardJsonBytes, _ := json.Marshal(userCard)
			cardJson := string(cardJsonBytes)
			skills.ForEach(func(kk, vv gjson.Result) bool {
				skillIdKey := fmt.Sprintf("additional_passive_skill_%d_id", kk.Int()+1)
				cardJson, _ = sjson.Set(cardJson, skillIdKey, vv.Int())
				return true
			})
			if err := json.Unmarshal([]byte(cardJson), &userCard); err != nil {
				panic(err)
			}
			session.UpdateUserCard(userCard)
		}
		return true
	})
	signBody := session.Finalize(GetData("skillEditResult.json"), "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func SaveDeckLesson(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type SaveDeckReq struct {
		DeckId        int32   `json:"deck_id"`
		CardMasterIds []int32 `json:"card_master_ids"`
	}
	req := SaveDeckReq{}
	if err := json.Unmarshal([]byte(reqBody), &req); err != nil {
		panic(err)
	}

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	userLessonDeck := session.GetUserLessonDeck(req.DeckId)
	deckByte, _ := json.Marshal(userLessonDeck)
	deckInfo := string(deckByte)
	for i := 0; i < len(req.CardMasterIds); i += 2 {
		deckInfo, _ = sjson.Set(deckInfo, fmt.Sprintf("card_master_id_%d", req.CardMasterIds[i]), req.CardMasterIds[i+1])
	}
	if err := json.Unmarshal([]byte(deckInfo), &userLessonDeck); err != nil {
		panic(err)
	}
	session.UpdateLessonDeck(userLessonDeck)
	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func ChangeDeckNameLessonDeck(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type ChangeDeckNameReq struct {
		DeckId   int32  `json:"deck_id"`
		DeckName string `json:"deck_name"`
	}
	req := ChangeDeckNameReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	lessonDeck := session.GetUserLessonDeck(req.DeckId)
	lessonDeck.Name = req.DeckName
	session.UpdateLessonDeck(lessonDeck)
	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
