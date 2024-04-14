package sif_2_data_link

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func dataLink(ctx *gin.Context) {
	req := request.Sif2DataLinkRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	if req.IsPermission {
		common.JsonResponse(ctx, &response.Sif2DataLinkResponse{
			PassWord: "Kashikoi Kawaii Elichika",
		})
	} else {
		common.JsonResponse(ctx, &response.EmptyResponse{})
	}
}

func init() {
	router.AddHandler("/", "POST", "/sif2DataLink/dataLink", dataLink)
}
