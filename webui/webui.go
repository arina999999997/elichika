package webui

import (
	"elichika/gamedata"
	"elichika/locale"
	"elichika/userdata"
	"elichika/utils"

	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// TODO: it's possible to replace serial.llas.bushimo.jp and use that button to redirect here from inside the game, maybe I'll do it one day
const commonPrefix = "/webui/response.html?response="

func Common(ctx *gin.Context) {
	var err error
	ctx.Set("has_user_id", true)
	form, err := ctx.MultipartForm()
	utils.CheckErr(err)
	lang := "en"
	if len(form.Value["client"]) > 0 {
		lang = form.Value["client"][0]
	}
	// some request don't actually need these, but a session expect them so we provide some stuff anyway
	ctx.Set("locale", locale.Locales[lang])
	ctx.Set("gamedata", locale.Locales[lang].Gamedata)
	ctx.Set("dictionary", locale.Locales[lang].Dictionary)

	userIDString := form.Value["user_id"][0]
	userID := -1
	userID, err = strconv.Atoi(userIDString)
	if err != nil {
		exist, err := userdata.Engine.Table("u_info").OrderBy("last_login_at DESC").Limit(1).Cols("user_id").Get(&userID)
		utils.CheckErr(err)
		if !exist {
			ctx.Set("has_user_id", false)
			ctx.Redirect(http.StatusFound, commonPrefix+"Error: there is no user in the database, start playing first")
			return
		}
	}
	ctx.Set("user_id", userID)
	ctx.Next()
}
func Birthday(ctx *gin.Context) {
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
	session.UserStatus.LastLoginAt = time.Now().Unix()
	form, _ := ctx.MultipartForm()

	birthdayString := form.Value["birthday"][0]
	if birthdayString == "" {
		ctx.Redirect(http.StatusFound, commonPrefix+"Error: no birthday given")
		return
	}
	tokens := strings.Split(birthdayString, "-")
	year, _ := strconv.Atoi(tokens[0])
	month, _ := strconv.Atoi(tokens[1])
	day, _ := strconv.Atoi(tokens[2])
	session.UserStatus.BirthDate = year*10000 + month*100 + day
	session.UserStatus.BirthDay = day
	session.UserStatus.BirthMonth = month
	session.Finalize("{}", "dummy")
	ctx.Redirect(http.StatusFound, commonPrefix+fmt.Sprintf("Success: update birthday for user %d to %d/%d/%d", userID, year, month, day))
}

func Accessory(ctx *gin.Context) {
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
	session.UserStatus.LastLoginAt = time.Now().Unix()
	form, _ := ctx.MultipartForm()
	specificAccessoryString := form.Value["accessory_id"][0]
	accessoryIDs := []int{}
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	if specificAccessoryString != "" {
		value, _ := strconv.Atoi(specificAccessoryString)
		accessoryIDs = append(accessoryIDs, value)
	} else {
		getRarity := make(map[int]bool)
		getRarity[30] = len(form.Value["ur_accessories"]) > 0
		getRarity[20] = len(form.Value["sr_accessories"]) > 0
		getRarity[10] = len(form.Value["r_accessories"]) > 0
		for _, accessory := range gamedata.Accessory {
			if getRarity[accessory.RarityType] {
				accessoryIDs = append(accessoryIDs, accessory.ID)
			}
		}
	}
	if len(accessoryIDs) == 0 {
		ctx.Redirect(http.StatusFound, commonPrefix+"Error: no accessory found, add a specific ID or choose at least one rarity")
		return
	}
	amount, _ := strconv.Atoi(form.Value["accessory_amount"][0])
	index := time.Now().UnixNano()
	total := 0
	for _, accessoryMasterID := range accessoryIDs {
		masterAccessory, exist := gamedata.Accessory[accessoryMasterID]
		if !exist {
			ctx.Redirect(http.StatusFound, commonPrefix+fmt.Sprint("Error: invalid accessory id ", accessoryMasterID))
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
	session.Finalize("{}", "dummy")
	ctx.Redirect(http.StatusFound, commonPrefix+fmt.Sprint("Success: Added ", total, " accessories"))
}
