//go:build dev

package event_marathon_dev

import (
	"elichika/client"
	"elichika/config"
	"elichika/generic"
	"elichika/router"
	"elichika/utils"
	"elichika/webui/image_form"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	TitleImagesAssetPaths = []string{
		`^[h`,
		`$&0`,
		`Q+|`,
		`.#+`,
		`nm*`,
		`7(M`,
		`pWy`,
		`VG+`,
		`)Ou`,
		`})?`,
		`fwU`,
		`!;1`,
		`pEP`,
		`wZ.`,
		`$>k`,
		`1f?`,
		`Y]t`,
		`^#;`,
		`]){`,
		`fGz`,
		`~%S`,
		`4FM`,
		`<KO`,
		`Z7C`,
		`WTV`,
		`Dxt`,
		`5d3`,
		`AaG`,
		`ew8`,
		`s9Y`,
		`:?&`,
		`lL^`,
		`BXC`,
		`Uaj`,
		`?5}`,
		`}EW`,
		`SwE`,
		`iPV`,
		`.!K`,
		`p(R`,
		`I9Z`,
		`L2k`,
		`\%d`,
		`:\c`,
		`,5m`,
		`ZAv`,
		`\(^`,
		`{%D`,
		`>_U`,
		`N'e`,
		`i9g`,
		`eF>`,
		`}Hv`,
		`Kcz`,
		`cS(`,
		`uBp`,
		`uBp`,
		`H3G`,
		`b[g`,
		`jYA`,
		"2`/",
		`8GA`,
		`Vja`,
		`Jb8`,
		`F5:`,
		`ND{`,
		`gpP`,
		`Jv~`,
		`JwY`,
		`7!9`,
		`DvP`,
		`Z(l`,
		`*v\`,
		`uqN`,
		`zD6`,
		`U~i`,
		`1:z`,
		`R2c`,
		`J.&`,
		`ty'`,
		`D&C`,
		`f)T`,
		`'.4`,
		`24c`,
		`YGp`,
		`cTV`,
		`kX@`,
		`0}+`,
		`5[>`,
		`K9%`,
		`Uz!`,
		`Zb$`,
		`4%P`,
		`6[L`,
		`GB~`,
		`?0=`}
)

func EventMarathonDev03GET(ctx *gin.Context) {

	form := image_form.ImageForm{
		FormId:    "title_image_path_form",
		DataLabel: "Event title image",
		DataId:    "title_image_path",
	}
	for _, assetPath := range TitleImagesAssetPaths {
		form.AddImageByAssetPath("", assetPath)
	}

	ctx.Header("Content-Type", "text/html")
	msg := form.GetHTML()
	ctx.HTML(http.StatusOK, "event_marathon_dev.html", gin.H{
		"body": msg,
	})
}

func EventMarathonDev03POST(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	utils.CheckErr(err)
	fmt.Println(form.Value)
	TopStatus.TitleImagePath = client.TextureStruktur{
		V: generic.NewNullable[string](form.Value["title_image_path"][0]),
	}
	ctx.Header("Location", "/webui/event_marathon_dev/04")
	ctx.String(http.StatusSeeOther, "")
}

func init() {
	if config.DeveloperMode == config.DeveloperModeEventMarathonDev {
		router.AddHandler("/webui/event_marathon_dev", "GET", "/03", EventMarathonDev03GET)
		router.AddHandler("/webui/event_marathon_dev", "POST", "/03", EventMarathonDev03POST)
	}
}
