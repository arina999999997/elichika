package handler

import (
	"elichika/config"
	"elichika/protocol/response"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func FetchDailyTheater(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	signBody := session.Finalize(GetData("userModelDiff.json"), "user_model_diff")
	respObj := response.DailyTheaterDetail{
		DailyTheaterID: 1,
		Year:           2021,
		Month:          6,
		Day:            24,
	}
	respObj.Title.DotUnderText = `『誤解から生まれる評判』`
	respObj.DetailText.DotUnderText =
		`<:th_ch0101/>善子ちゃーん、お願いがあるんだけど<:dt_line_end/><:th_ch0106/>善子じゃなくてヨハネね。お願いってなに？<:dt_line_end/><:th_ch0101/>あのね、今度私の生配信にゲストとして出てほしいんだ。やっぱり善子ちゃんが出てると生配信もすっごく盛り上がってるから！<:dt_line_end/><:th_ch0106/>フッ、まあヨハネは誰よりも先にリトルデーモンたちの声を直接聞いていたから当然ね……。いいわ、出てあげる<:dt_line_end/><:th_ch0101/>やったー！　ありがとう！<:dt_line_end/><:th_ch0106/>それで何の生配信をやるの？<:dt_line_end/><:th_ch0101/>えっとね、ゲーム実況をやってみようと思って<:dt_line_end/><:th_ch0106/>え、ゲーム実況……？<:dt_line_end/><:th_ch0101/>うん。この前花陽ちゃんと一緒にホラーゲームの実況やってたでしょ？　そのときの善子ちゃんの叫び声とか、動きとか、盛り上げ方が本当にすごかったから、私も善子ちゃんと一緒にホラーゲームやってみたくって！<:dt_line_end/><:th_ch0106/>あ、あれは盛り上げとかじゃなくて……、っていうかもうホラーゲームは勘弁してよーーー！！<:dt_line_end/>`

	signBody, _ = sjson.Set(signBody, "daily_theater_detail", respObj)
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func DailyTheaterSetLike(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type SetLikeReq struct {
		DailyTheaterID int  `json:"daily_theater_id"`
		IsLike         bool `json:"is_like"`
	}
	req := SetLikeReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	type UserDailyTheater struct {
		DailyTheaterID int  `json:"daily_theater_id"`
		IsLiked        bool `json:"is_liked"`
	}

	// fmt.Println(reqBody)
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	signBody := session.Finalize(GetData("userModel.json"), "user_model")

	response := []any{}
	response = append(response, req.DailyTheaterID)
	response = append(response, UserDailyTheater{
		DailyTheaterID: req.DailyTheaterID,
		IsLiked:        req.IsLike,
	})
	signBody, _ = sjson.Set(signBody, "user_model.user_daily_theater_by_daily_theater_id", response)
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func FetchDailyTheaterArchive(ctx *gin.Context) {
	// this is used to publish new daily theater without having to update the database
	// client have the old items in m_daily_theater_archive_client and m_daily_theater_archive_member_client
	// client's missing 20230629 and 20230630
	respObj := response.FetchDailyTheaterArchiveResponse{}
	respObj.DailyTheaterArchiveMasterRows = []response.DailyTheaterArchiveMasterRow{}
	respObj.DailyTheaterArchiveMemberMasterRows = []response.DailyTheaterArchiveMemberMasterRow{}

	// this isn't the actual thing
	respObj.DailyTheaterArchiveMasterRows = append(respObj.DailyTheaterArchiveMasterRows,
		response.DailyTheaterArchiveMasterRow{
			DailyTheaterID: 1001243,
			Year:           2023,
			Month:          6,
			Day:            29,
			PublishedAt:    1687964400,
		})
	respObj.DailyTheaterArchiveMemberMasterRows = append(respObj.DailyTheaterArchiveMemberMasterRows,
		response.DailyTheaterArchiveMemberMasterRow{
			DailyTheaterID: 1001243,
			MemberMasterID: 101, // Chika
		})
	respObj.DailyTheaterArchiveMemberMasterRows = append(respObj.DailyTheaterArchiveMemberMasterRows,
		response.DailyTheaterArchiveMemberMasterRow{
			DailyTheaterID: 1001243,
			MemberMasterID: 106, // Yoshiko
		})

	signBody, _ := json.Marshal(respObj)

	resp := SignResp(ctx, string(signBody), config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
