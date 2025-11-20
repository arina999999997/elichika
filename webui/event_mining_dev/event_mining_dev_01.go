//go:build dev

package event_mining_dev

import (
	"elichika/config"
	"elichika/router"
	"elichika/utils"
	"elichika/webui/form/select_form"

	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func EventMiningDev01GET(ctx *gin.Context) {
	form := select_form.SelectForm{
		FormId:    "event_id_form",
		DataLabel: "Event Id",
		DataId:    "event_id",
	}
	// fetching fromdb would require more dependency which we don't already have
	form.AddOption("31041", "Sweet Dreaming")
	form.AddOption("31040", "Red Hot Baseball Game!")
	form.AddOption("31039", "One and Only Girl's Day")
	form.AddOption("31038", "The Setsubun Chase Is On!")
	form.AddOption("31037", "New Year, New Dreams")
	form.AddOption("31036", "Happiness Snow Time")
	form.AddOption("31035", "Kanata's Treats and Autumn")
	form.AddOption("31034", "Go! Passionate Cheer!")
	form.AddOption("31033", "Joy on a Moonlit Night")
	form.AddOption("31032", "Sunset Summer Challenge!")
	form.AddOption("31031", "Link across the Milky Way!")
	form.AddOption("31030", "Hear the Sound of My Soul!")
	form.AddOption("31029", "A Stylish Job for a Dazzling Future")
	form.AddOption("31028", "Fleeting Night Sakura")
	form.AddOption("31027", "Bonding on White Day")
	form.AddOption("31026", "Flowers for Valentine's")
	form.AddOption("31025", "Whoo! New Year Girls Party!")
	form.AddOption("31024", "White City Lights")
	form.AddOption("31023", "Hospitality at the School Festival")
	form.AddOption("31022", "Wear the Outfit I Picked!")
	form.AddOption("31021", "Moonlit Fairy Tale")
	form.AddOption("31020", "Aloha! Tropical Getaway")
	form.AddOption("31019", "Kasumi's Star Festival Wish")
	form.AddOption("31018", "Rainy Season Blessings")
	form.AddOption("31017", "Kindergarten Live Show!")
	form.AddOption("31016", "Seeking the Cherry Blossom Fairy")
	form.AddOption("31015", "Sharing My Feelings on White Day")
	form.AddOption("31014", "A Valentine's Day Full of Love")
	form.AddOption("31013", "New Year Shrine Visit, Uranohoshi-Style")
	form.AddOption("31012", "The Best Christmas Present")
	form.AddOption("31011", "Strut Down the Runway")
	form.AddOption("31010", "The Rabbit and Goddess of the Moon")
	form.AddOption("31009", "Rebel-ish Makeover")
	form.AddOption("31008", "Grab Victory in the Sports Battle!")
	form.AddOption("31007", "Yohane and Hanayo's Whodunit Caper")
	form.AddOption("31006", "Three Princesses")
	form.AddOption("31005", "Hot Spring Rhapsody")
	form.AddOption("31004", "Cooking with Vegetables!")
	form.AddOption("31003", "Catch the Mischievous Wolf!")
	form.AddOption("31002", "Music Made Together")
	form.AddOption("31001", "Great Battle on the High Seas")

	ctx.Header("Content-Type", "text/html")
	msg := form.GetHTML()
	ctx.HTML(http.StatusOK, "event_mining_dev.html", gin.H{
		"body": msg,
	})
}

func EventMiningDev01POST(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	utils.CheckErr(err)
	fmt.Println(form.Value)
	eventId, err := strconv.Atoi(form.Value["event_id"][0])
	utils.CheckErr(err)
	TopStatus.EventId = int32((eventId))
	ctx.Header("Location", "/webui/event_mining_dev/02")
	ctx.String(http.StatusSeeOther, "")
}

func init() {
	if config.DeveloperMode == config.DeveloperModeEventMiningDev {
		router.AddHandler("/webui/event_mining_dev", "GET", "/", EventMiningDev01GET)
		router.AddHandler("/webui/event_mining_dev", "POST", "/", EventMiningDev01POST)
	}
}
