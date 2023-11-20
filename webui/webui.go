package webui

import (
	"elichika/gamedata"
	"elichika/locale"
	"elichika/userdata"
	"elichika/utils"

	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	// "github.com/tidwall/gjson"
	// "github.com/tidwall/sjson"
)

const CommonPrefix = "/webui/response.html?"

func BuildPrefix(ctx *gin.Context) string {
	var res string
	res = CommonPrefix
	for key, value := range *ctx.MustGet("params").(*map[string]string) {
		res += key + "=" + value + "&"
	}
	return res + "response="
}

func Common(ctx *gin.Context) {
	var err error
	ctx.Set("is_good", true)
	params := make(map[string]string)
	temp, _ := url.ParseQuery(ctx.GetString("reqBody"))
	for key, value := range temp {
		if len(value) > 0 {
			if value[0] == "" {
				continue
			}
			params[key] = value[0]
		}
	}
	ctx.Set("params", &params)

	lang := params["client"]
	ctx.Set("locale", locale.Locales[lang])
	ctx.Set("gamedata", locale.Locales[lang].Gamedata)
	ctx.Set("dictionary", locale.Locales[lang].Dictionary)

	userIDString := params["user_id"]
	userID := 0
	userID, err = strconv.Atoi(userIDString)
	if err != nil {
		exists, err := userdata.Engine.Table("u_info").OrderBy("last_login_at DESC").Limit(1).Cols("user_id").Get(&userID)
		utils.CheckErr(err)
		if !exists {
			ctx.Set("is_good", false)
			ctx.Redirect(http.StatusFound, BuildPrefix(ctx)+"Error: there is no user in the database, start playing first")
			return
		}
	}

	ctx.Set("user_id", userID)
	params["user_id"] = fmt.Sprint(userID)

	ctx.Next()
}
func Birthday(ctx *gin.Context) {
	if !ctx.MustGet("is_good").(bool) {
		return
	}

	userID := ctx.MustGet("user_id").(int)
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	if session == nil {
		ctx.Redirect(http.StatusFound, BuildPrefix(ctx)+fmt.Sprint("Error: user ", userID, " doesn't exists"))
		return
	}
	// var err error
	params := *ctx.MustGet("params").(*map[string]string)
	birthdayString, exists := params["birthday"]
	if !exists {
		ctx.Redirect(http.StatusFound, BuildPrefix(ctx)+"Error: no birthday given")
		return
	}
	tokens := strings.Split(birthdayString, "-")
	year, _ := strconv.Atoi(tokens[0])
	month, _ := strconv.Atoi(tokens[1])
	day, _ := strconv.Atoi(tokens[2])
	session.UserStatus.BirthDate = year*10000 + month*100 + day
	session.UserStatus.BirthDay = day
	session.UserStatus.BirthMonth = month
	session.Finalize("", "")
	ctx.Redirect(http.StatusFound, BuildPrefix(ctx)+fmt.Sprintf("Success: update birthday for user %d to %d/%d/%d", userID, year, month, day))
}

func Accessory(ctx *gin.Context) {
	if !ctx.MustGet("is_good").(bool) {
		return
	}
	userID := ctx.MustGet("user_id").(int)
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	if session == nil {
		ctx.Redirect(http.StatusFound, BuildPrefix(ctx)+fmt.Sprint("Error: user ", userID, " doesn't exists"))
		return
	}
	params := *ctx.MustGet("params").(*map[string]string)
	specificAccessoryString, exists := params["accessory_id"]
	accessoryIDs := []int{}
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	if exists {
		value, _ := strconv.Atoi(specificAccessoryString)
		accessoryIDs = append(accessoryIDs, value)
	} else {
		getRarity := make(map[int]bool)
		_, getRarity[30] = params["ur_accessories"]
		_, getRarity[20] = params["sr_accessories"]
		_, getRarity[10] = params["r_accessories"]
		for _, accessory := range gamedata.Accessory {
			if getRarity[accessory.RarityType] {
				accessoryIDs = append(accessoryIDs, accessory.ID)
			}
		}
	}
	if len(accessoryIDs) == 0 {
		ctx.Redirect(http.StatusFound, BuildPrefix(ctx)+"Error: no accessory found, add a specific ID or choose at least one rarity")
		return
	}
	amount, _ := strconv.Atoi(params["accessory_amount"])
	index := time.Now().UnixNano()
	total := 0
	for _, accessoryMasterID := range accessoryIDs {
		masterAccessory, exists := gamedata.Accessory[accessoryMasterID]
		if !exists {
			ctx.Redirect(http.StatusFound, BuildPrefix(ctx)+fmt.Sprint("Error: invalid accessory id ", accessoryMasterID))
			return
		}
		for i := 1; i <= amount; i++ {
			total++
			accessory := session.GetUserAccessory(index + int64(total))
			accessory.AccessoryMasterID = masterAccessory.ID
			accessory.Level = 1
			accessory.Exp = 0
			accessory.Grade = 0
			accessory.Attribute = masterAccessory.Attribute
			accessory.PassiveSkill1ID = *masterAccessory.Grade[0].PassiveSkill1MasterID
			accessory.PassiveSkill2ID = masterAccessory.Grade[0].PassiveSkill2MasterID
			session.UpdateUserAccessory(accessory)
		}
	}
	session.Finalize("", "")
	ctx.Redirect(http.StatusFound, BuildPrefix(ctx)+fmt.Sprint("Success: Added ", total, " accessories"))
}
