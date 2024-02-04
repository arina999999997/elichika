package webui


import (
	"elichika/client"
	"elichika/enum"
	"elichika/subsystem/user_present"
	"elichika/userdata"
	// "elichika/utils"

	"fmt"
	"net/http"
	"strconv"
	// "strings"
	"time"

	"github.com/gin-gonic/gin"
)

func AddPresent(ctx *gin.Context) {
	if !ctx.MustGet("has_user_id").(bool) {
		return
	}

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	if session == nil {
		ctx.Redirect(http.StatusFound, commonPrefix+fmt.Sprint("Error: user ", userId, " doesn't exist"))
		return
	}
	session.UserStatus.LastLoginAt = time.Now().Unix()
	form, _ := ctx.MultipartForm()

	contentTypeString := form.Value["content_type"][0]
	contentIdString := form.Value["content_id"][0]
	if contentIdString == "" {
		ctx.Redirect(http.StatusFound, commonPrefix+"Error: no content id given")
		return
	}

	contentAmountString := form.Value["content_amount"][0]
	if contentAmountString == "" {
		contentAmountString = "1"
	}

	contentType, _ := strconv.ParseInt(contentTypeString, 16, 64)
	contentId, _ := strconv.Atoi(contentIdString)
	contentAmount, _ := strconv.Atoi(contentAmountString)
	
	user_present.AddPresent(session, client.PresentItem{
		PresentRouteType: enum.PresentRouteTypeAdminPresent,
		Content: client.Content{
			ContentType: int32(contentType),
			ContentId: int32(contentId),
			ContentAmount: int32(contentAmount),
		}})
	
	session.Finalize()
	ctx.Redirect(http.StatusFound, commonPrefix+fmt.Sprintf("Success: Added item (%d %d %d) for user %d", contentType, contentId, contentAmount, userId))
}