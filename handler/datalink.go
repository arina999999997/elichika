package handler

import (
	"elichika/config"
	"net/http"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func FetchDataLinks(ctx *gin.Context) {
	type DataLinks struct {
		IsPlatformLinked             bool `json:"is_platform_linked"`
		IsSchoolIdolFestivalIDLinked bool `json:"is_school_idol_festival_id_linked"`
		IsTakeOverIDLinked           bool `json:"is_take_over_id_linked"`
	}
	links := DataLinks{
		IsPlatformLinked:             false,
		IsSchoolIdolFestivalIDLinked: false,
		IsTakeOverIDLinked:           true,
	}
	signBodyBytes, _ := json.Marshal(links)
	signBody := string(signBodyBytes)
	resp := SignResp(ctx, string(signBody), config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
