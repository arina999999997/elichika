package webui

import (
	"elichika/account"
	"elichika/userdata"
	"elichika/utils"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ImportAccount(ctx *gin.Context) {
	userId := ctx.GetInt("user_id")
	{
		session := userdata.GetSession(ctx, userId)
		defer session.Close()
		if session != nil {
			ctx.Redirect(http.StatusFound, commonPrefix+fmt.Sprint("Error: User ", userId, " already exists, select a different user Id or delete that account first"))
			return
		}
	}

	form, _ := ctx.MultipartForm()
	for _, fileHeader := range form.File["account_data"] {
		file, err := fileHeader.Open()
		utils.CheckErr(err)
		bytes := make([]byte, fileHeader.Size)
		length, err := file.Read(bytes)
		utils.CheckErr(err)
		if int64(length) != fileHeader.Size {
			panic("error reading file")
		}
		ctx.Redirect(http.StatusFound, commonPrefix+account.ImportUser(ctx, string(bytes), userId))
	}
}

func ExportAccount(ctx *gin.Context) {
	if !ctx.MustGet("has_user_id").(bool) {
		return
	}
	userId := ctx.GetInt("user_id")
	{
		session := userdata.GetSession(ctx, userId)
		defer session.Close()
		if session == nil {
			ctx.Redirect(http.StatusFound, commonPrefix+fmt.Sprint("Error: User ", userId, " doesn't exists"))
			return
		}
	}
	content := account.ExportUser(ctx)

	ctx.Header("Content-Disposition", fmt.Sprint("attachment; filename=login_", userId, ".json"))
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Accept-Length", fmt.Sprint(len(content)))
	ctx.Writer.Write([]byte(content))
}
