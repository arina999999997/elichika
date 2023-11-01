package handler

import (
	"elichika/config"
	"elichika/locale"

	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func GetPackUrl(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")
	// fmt.Println(reqBody)

	var packNames []string
	gjson.Parse(reqBody).ForEach(func(key, value gjson.Result) bool {
		if value.IsObject() {
			if err := json.Unmarshal([]byte(value.Get("pack_names").String()), &packNames); err != nil {
				panic(err)
			}
			return false
		}
		return true
	})

	var packUrls []string
	cdnMasterVersion := "2d61e7b4e89961c7"
	if ctx.MustGet("locale").(*locale.Locale).MasterVersion == config.MasterVersionJp {
		cdnMasterVersion = "b66ec2295e9a00aa"
	}
	for _, pack := range packNames {

		packUrls = append(packUrls, config.Conf.CdnServer+"/"+cdnMasterVersion+"/"+pack)
	}

	packBody, _ := sjson.Set("{}", "url_list", packUrls)
	resp := SignResp(ctx, packBody, config.SessionKey)
	// fmt.Println("Response:", resp)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
