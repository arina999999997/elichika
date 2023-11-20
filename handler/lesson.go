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
		ExecuteLessonIDs   []int `json:"execute_lesson_ids"`
		ConsumedContentIDs []int `json:"consumed_content_ids"`
		SelectedDeckID     int   `json:"selected_deck_id"`
		IsThreeTimes       bool  `json:"is_three_times"`
	}
	req := ExecuteLessonReq{}
	if err := json.Unmarshal([]byte(reqBody), &req); err != nil {
		panic(err)
	}

	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	deckBytes, _ := json.Marshal(session.GetUserLessonDeck(req.SelectedDeckID))
	deckInfo := string(deckBytes)
	var actionList []model.LessonMenuAction

	gjson.Parse(deckInfo).ForEach(func(key, value gjson.Result) bool {
		if strings.Contains(key.String(), "card_master_id") {
			actionList = append(actionList, model.LessonMenuAction{
				CardMasterID:                  value.Int(),
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

	session.UserStatus.MainLessonDeckID = req.SelectedDeckID
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
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	signBody := session.Finalize(GetData("resultLesson.json"), "user_model_diff")
	signBody, _ = sjson.Set(signBody, "selected_deck_id", session.UserStatus.MainLessonDeckID)
	resp := SignResp(ctx, signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func SkillEditResult(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")

	req := gjson.Parse(reqBody).Array()[0]

	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	skillList := req.Get("selected_skill_ids")
	skillList.ForEach(func(key, cardId gjson.Result) bool {
		if key.Int()%2 == 0 {
			userCard := session.GetUserCard(int(cardId.Int()))
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
		DeckID        int   `json:"deck_id"`
		CardMasterIDs []int `json:"card_master_ids"`
	}
	req := SaveDeckReq{}
	if err := json.Unmarshal([]byte(reqBody), &req); err != nil {
		panic(err)
	}

	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	userLessonDeck := session.GetUserLessonDeck(req.DeckID)
	deckByte, _ := json.Marshal(userLessonDeck)
	deckInfo := string(deckByte)
	for i := 0; i < len(req.CardMasterIDs); i += 2 {
		deckInfo, _ = sjson.Set(deckInfo, fmt.Sprintf("card_master_id_%d", req.CardMasterIDs[i]), req.CardMasterIDs[i+1])
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
		DeckID   int    `json:"deck_id"`
		DeckName string `json:"deck_name"`
	}
	req := ChangeDeckNameReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	lessonDeck := session.GetUserLessonDeck(req.DeckID)
	lessonDeck.Name = req.DeckName
	session.UpdateLessonDeck(lessonDeck)
	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
