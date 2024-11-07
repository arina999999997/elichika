package admin

import (
	"elichika/gamedata"
	"elichika/router"
	"elichika/subsystem/event"
	"elichika/utils"
	"elichika/webui/object_form"
	"elichika/webui/webui_utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

type EventSelectForm struct {
	// Event *int32 `json:"event" of_type:"select" of_options:"Secret Party!\n30001\nAutumn Rain Club\n30027\nCutie Wonderland\n30035" of_label:"Event"`
	Event *int32 `json:"event" of_type:"select" of_options_external:"events" of_label:"Event"`
}

func init() {
	// TODO(locale): Make other language available too
	object_form.SetExternalOptions("events", gamedata.GamedataByLocale["en"].EventAvailable.GetEventToIdList())
}

func EventSelector(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")

	starts := `<head><meta name="viewport" content="width=device-width, initial-scale=1"/></head>
	<div>Select active event</div>
	<div>Note that once you click the change event button, the current event (if any) will be ended, the reward will be given out, and the selected event will be started instantly.</div>
	<div>This happens even if you have selected the same event, so if you are here by mistake, just go back.</div>
	`

	eventForm := EventSelectForm{
		Event: new(int32),
	}
	*eventForm.Event = 30001
	ctx.HTML(http.StatusOK, "logged_in_admin.html", gin.H{
		"body": starts + object_form.GenerateWebForm(&eventForm, "config_form", ` onclick="submit_form('config_form', './change_event')"`, "", "Select and start new event"),
	})
}

func ChangeEvent(ctx *gin.Context) {
	eventForm := EventSelectForm{}
	err := object_form.ParseForm(ctx, &eventForm)
	utils.CheckErr(err)
	eventId := *eventForm.Event
	event.ChangeEvent(eventId)
	webui_utils.CommonResponse(ctx, "Event changed!", "")
}

// instead of selecting an event, schedule an event that would be selected by the event auto scheduler
// also add the event auto scheduler into the scheduled task if necessary
func EventScheduler(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")

	starts := `<head><meta name="viewport" content="width=device-width, initial-scale=1"/></head>
	<div>Schedule next event</div>
	<div>Only one event can be scheduled at a time, the scheduled one are displayed</div>
	<div>If no event is scheduled, then None will be displayed, and the server will automatically choose an event.</div>
	`
	eventForm := EventSelectForm{
		Event: new(int32),
	}
	*eventForm.Event = 0
	ctx.HTML(http.StatusOK, "logged_in_admin.html", gin.H{
		"body": starts + object_form.GenerateWebForm(&eventForm, "config_form", ` onclick="submit_form('config_form', './schedule_event')"`, "", "Select and schedule event"),
	})
}

func ScheduleEvent(ctx *gin.Context) {
	eventForm := EventSelectForm{}
	err := object_form.ParseForm(ctx, &eventForm)
	utils.CheckErr(err)
	eventId := *eventForm.Event
	event.ScheduleEvent(eventId)
	webui_utils.CommonResponse(ctx, "Event scheduled!", "")
}

func init() {
	router.AddHandler("/webui/admin", "GET", "/event_selector", EventSelector)
	router.AddHandler("/webui/admin", "POST", "/change_event", ChangeEvent)
	router.AddHandler("/webui/admin", "GET", "/event_scheduler", EventScheduler)
	router.AddHandler("/webui/admin", "POST", "/schedule_event", ScheduleEvent)
}
