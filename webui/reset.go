package webui

import (
	"elichika/userdata"
	// "elichika/utils"

	"fmt"
	"net/http"
	// "strconv"
	// "strings"
	// "time"

	"github.com/gin-gonic/gin"
)

func ResetProgress(ctx *gin.Context) {
	if !ctx.MustGet("has_user_id").(bool) {
		return
	}
	userID := ctx.MustGet("user_id").(int)
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	if session == nil {
		ctx.Redirect(http.StatusFound, commonPrefix+fmt.Sprint("Error: user ", userID, " doesn't exist"))
		return
	}
	fmt.Println(ctx.Request.URL.Path)
	switch ctx.Request.URL.Path {
	case "/webui/reset_story_main":
		session.RemoveUserProgress("u_story_main")
		session.RemoveUserProgress("u_story_main_part_digest_movie")
		session.RemoveUserProgress("u_story_main_selected")
	case "/webui/reset_story_side":
		session.MarkIsNew("u_story_side", true)
	case "/webui/reset_story_member":
		session.MarkIsNew("u_story_member", true)
	case "/webui/reset_story_linkage":
		session.RemoveUserProgress("u_story_linkage")
	case "/webui/reset_story_event":
		session.RemoveUserProgress("u_story_event_history")
	case "/webui/reset_dlp":
		session.RemoveUserProgress("u_tower")
	}
	session.Finalize("{}", "dummy")
	ctx.Redirect(http.StatusFound, commonPrefix+"reseted progress, relogin to see the change")
}
