package handler

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/generic"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

// TODO(daily_theater): Actually implement this system
func FetchDailyTheater(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FetchDailyTheaterRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	if !req.DailyTheaterId.HasValue { // latest item
		req.DailyTheaterId = generic.NewNullable(int32(1))
	}
	// fetch and track the item

	session.Finalize("", "dummy")

	JsonResponse(ctx, response.FetchDailyTheaterResponse{
		DailyTheaterDetail: client.DailyTheaterDetail{
			DailyTheaterId: req.DailyTheaterId.Value,
			Title: client.LocalizedText{
				DotUnderText: `『誤解から生まれる評判』`,
			},
			DetailText: client.LocalizedText{
				DotUnderText: `<:th_ch0101/>善子ちゃーん、お願いがあるんだけど<:dt_line_end/><:th_ch0106/>善子じゃなくてヨハネね。お願いってなに？<:dt_line_end/><:th_ch0101/>あのね、今度私の生配信にゲストとして出てほしいんだ。やっぱり善子ちゃんが出てると生配信もすっごく盛り上がってるから！<:dt_line_end/><:th_ch0106/>フッ、まあヨハネは誰よりも先にリトルデーモンたちの声を直接聞いていたから当然ね……。いいわ、出てあげる<:dt_line_end/><:th_ch0101/>やったー！　ありがとう！<:dt_line_end/><:th_ch0106/>それで何の生配信をやるの？<:dt_line_end/><:th_ch0101/>えっとね、ゲーム実況をやってみようと思って<:dt_line_end/><:th_ch0106/>え、ゲーム実況……？<:dt_line_end/><:th_ch0101/>うん。この前花陽ちゃんと一緒にホラーゲームの実況やってたでしょ？　そのときの善子ちゃんの叫び声とか、動きとか、盛り上げ方が本当にすごかったから、私も善子ちゃんと一緒にホラーゲームやってみたくって！<:dt_line_end/><:th_ch0106/>あ、あれは盛り上げとかじゃなくて……、っていうかもうホラーゲームは勘弁してよーーー！！<:dt_line_end/>`,
			},
			Year:  2021,
			Month: 6,
			Day:   24,
		},
		UserModelDiff: &session.UserModel,
	})
}

func DailyTheaterSetLike(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.DailyTheaterSetLikeRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	session.UserModel.UserDailyTheaterByDailyTheaterId.Set(
		req.DailyTheaterId,
		client.UserDailyTheater{
			DailyTheaterId: req.DailyTheaterId,
			IsLiked:        req.IsLike,
		})

	session.Finalize("", "dummy")
	JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

// TODO(refactor): Change to use request and response types
func FetchDailyTheaterArchive(ctx *gin.Context) {
	// this is used to publish new daily theater without having to update the database
	// client have the old items in m_daily_theater_archive_client and m_daily_theater_archive_member_client
	// client's missing 20230629 and 20230630

	// There is no request body

	resp := response.FetchDailyTheaterArchiveResponse{}
	// this isn't the actual thing
	resp.DailyTheaterArchiveMasterRows.Append(
		client.DailyTheaterArchiveMasterRow{
			DailyTheaterId: 1001243,
			Year:           2023,
			Month:          6,
			Day:            29,
			PublishedAt:    1687964400,
		})
	resp.DailyTheaterArchiveMemberMasterRows.Append(client.DailyTheaterArchiveMemberMasterRow{
		DailyTheaterId: 1001243,
		MemberMasterId: 101, // Chika
	})
	resp.DailyTheaterArchiveMemberMasterRows.Append(client.DailyTheaterArchiveMemberMasterRow{
		DailyTheaterId: 1001243,
		MemberMasterId: 106, // Yoshiko
	})

	JsonResponse(ctx, &resp)
}
