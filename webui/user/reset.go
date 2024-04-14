package user

import (
	"elichika/router"
	"elichika/subsystem/reset_progress"
	"elichika/userdata"
	"elichika/utils"
	"elichika/webui/object_form"

	// "fmt"
	"encoding/json"
	"net/http"
	// "strings"
	// "time"

	"github.com/gin-gonic/gin"
)

type ResetRequest struct {
	StoryMain    bool `of_label:"Main story"`
	StorySide    bool `of_label:"Side story (card story)"`
	StoryMember  bool `of_label:"Member story (bond episode)"`
	StoryLinkage bool `of_label:"Linkage story (anime tie-in)"`
	StoryEvent   bool `of_label:"Event story"`
	Tower        bool `of_label:"DLP"`
}

func resetForm(ctx *gin.Context) {
	form := object_form.GenerateWebForm(&ResetRequest{}, "reset_form",
		` onclick="if (confirm('Reset progress?')) submit_form('reset_form', 'reset')"`, "Clear", "Reset progress")

	ctx.HTML(http.StatusOK, "logged_in_user.html", gin.H{
		"body": `<div><label>Choose aspect(s) to reset: </label></div>` + "\n" + form,
	})
}

func ResetHandler(ctx *gin.Context) {
	req := ResetRequest{}
	err := object_form.ParseForm(ctx, &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)
	resp := Response{
		Response: new(string),
	}

	if req.StoryMain {
		*resp.Response += "Main story progress was reset\n"
		reset_progress.RemoveUserProgress(session, "u_story_main")
		reset_progress.RemoveUserProgress(session, "u_story_main_part_digest_movie")
		reset_progress.RemoveUserProgress(session, "u_story_main_selected")
	}
	if req.StorySide {
		*resp.Response += "Side story progress was reset\n"
		reset_progress.MarkIsNew(session, "u_story_side", true)
	}
	if req.StoryMember {
		*resp.Response += "Bond story progress was reset\n"
		reset_progress.MarkIsNew(session, "u_story_member", true)
	}
	if req.StoryLinkage {
		*resp.Response += "Linkage story progress was reset\n"
		reset_progress.RemoveUserProgress(session, "u_story_linkage")
	}
	if req.StoryEvent {
		*resp.Response += "Event story unlock progress was reset, memory keys added\n"
		reset_progress.RemoveUserProgress(session, "u_story_event_history")
	}
	if req.Tower {
		*resp.Response += "DLP progress was reset\n"
		reset_progress.RemoveUserProgress(session, "u_tower")
	}
	session.Finalize()

	if *resp.Response == "" {
		*resp.Response = "There was nothing to reset!\n"
	}

	jsonBytes, err := json.Marshal(resp)
	utils.CheckErr(err)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, string(jsonBytes))
}

func init() {
	addFeature("Reset Progress", "reset")
	router.AddHandler("/webui/user", "GET", "/reset", resetForm)
	router.AddHandler("/webui/user", "POST", "/reset", ResetHandler)
}
