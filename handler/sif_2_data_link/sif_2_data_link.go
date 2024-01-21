package sif_2_data_link

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func Sif2DataLink(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.Sif2DataLinkRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
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
	router.AddHandler("/sif2DataLink/dataLink", Sif2DataLink)
}
