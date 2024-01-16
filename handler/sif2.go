package handler

import (
	"elichika/client/request"
	"elichika/client/response"
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
		JsonResponse(ctx, &response.Sif2DataLinkResponse{
			PassWord: "Kashikoi Kawaii Elichika",
		})
	} else {
		JsonResponse(ctx, &response.EmptyResponse{})
	}
}
