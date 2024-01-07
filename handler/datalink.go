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
		IsSchoolIdolFestivalIdLinked bool `json:"is_school_idol_festival_id_linked"`
		IsTakeOverIdLinked           bool `json:"is_take_over_id_linked"`
	}
	links := DataLinks{
		IsPlatformLinked:             false,
		IsSchoolIdolFestivalIdLinked: false,
		IsTakeOverIdLinked:           true,
	}
	signBodyBytes, _ := json.Marshal(links)
	signBody := string(signBodyBytes)
	resp := SignResp(ctx, string(signBody), config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
