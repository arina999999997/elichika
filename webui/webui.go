package webui

import (
	"elichika/gamedata"
	"elichika/generic"
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

// TODO: it's possible to replace serial.llas.bushimo.jp and use that button to redirect here from inside the game
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

	userIdString := form.Value["user_id"][0]
	userId := -1
	userId, err = strconv.Atoi(userIdString)
	if err != nil {
		exist, err := userdata.Engine.Table("u_info").OrderBy("last_login_at DESC").Limit(1).Cols("user_id").Get(&userId)
		utils.CheckErr(err)
		if !exist {
			ctx.Set("has_user_id", false)
			ctx.Redirect(http.StatusFound, commonPrefix+"Error: there is no user in the database, start playing first")
			return
		}
	}
	ctx.Set("user_id", userId)
	ctx.Next()
}

func Birthday(ctx *gin.Context) {
	if !ctx.MustGet("has_user_id").(bool) {
		return
	}

	userId := ctx.MustGet("user_id").(int)
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	if session == nil {
		ctx.Redirect(http.StatusFound, commonPrefix+fmt.Sprint("Error: user ", userId, " doesn't exist"))
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
	session.UserStatus.BirthDate = generic.NewNullable(int32(year*10000 + month*100 + day))
	session.UserStatus.BirthDay = generic.NewNullable(int32(day))
	session.UserStatus.BirthMonth = generic.NewNullable(int32(month))
	session.Finalize()
	ctx.Redirect(http.StatusFound, commonPrefix+fmt.Sprintf("Success: update birthday for user %d to %d/%d/%d", userId, year, month, day))
}

func Accessory(ctx *gin.Context) {
	if !ctx.MustGet("has_user_id").(bool) {
		return
	}
	userId := ctx.MustGet("user_id").(int)
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	if session == nil {
		ctx.Redirect(http.StatusFound, commonPrefix+fmt.Sprint("Error: user ", userId, " doesn't exist"))
		return
	}
	session.UserStatus.LastLoginAt = time.Now().Unix()
	form, _ := ctx.MultipartForm()
	specificAccessoryString := form.Value["accessory_id"][0]
	accessoryIds := []int32{}
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	if specificAccessoryString != "" {
		value, _ := strconv.Atoi(specificAccessoryString)
		accessoryIds = append(accessoryIds, int32(value))
	} else {
		getRarity := make(map[int32]bool)
		getRarity[30] = len(form.Value["ur_accessories"]) > 0
		getRarity[20] = len(form.Value["sr_accessories"]) > 0
		getRarity[10] = len(form.Value["r_accessories"]) > 0
		for _, accessory := range gamedata.Accessory {
			if getRarity[accessory.RarityType] {
				accessoryIds = append(accessoryIds, accessory.Id)
			}
		}
	}
	if len(accessoryIds) == 0 {
		ctx.Redirect(http.StatusFound, commonPrefix+"Error: no accessory found, add a specific Id or choose at least one rarity")
		return
	}
	amount, _ := strconv.Atoi(form.Value["accessory_amount"][0])
	index := time.Now().UnixNano()
	total := 0
	for _, accessoryMasterId := range accessoryIds {
		masterAccessory, exist := gamedata.Accessory[accessoryMasterId]
		if !exist {
			ctx.Redirect(http.StatusFound, commonPrefix+fmt.Sprint("Error: invalid accessory id ", accessoryMasterId))
			return
		}
		for i := 1; i <= amount; i++ {
			total++
			accessory := session.GetUserAccessory(index + int64(total))
			accessory.AccessoryMasterId = masterAccessory.Id
			accessory.Level = 1
			accessory.Exp = 0
			accessory.Grade = 0
			accessory.Attribute = masterAccessory.Attribute
			if masterAccessory.Grade[0].PassiveSkill1MasterId != nil {
				accessory.PassiveSkill1Id = generic.NewNullable(*masterAccessory.Grade[0].PassiveSkill1MasterId)
			}
			if masterAccessory.Grade[0].PassiveSkill2MasterId != nil {
				accessory.PassiveSkill2Id = generic.NewNullable(*masterAccessory.Grade[0].PassiveSkill2MasterId)
			}
			session.UpdateUserAccessory(accessory)
		}
	}
	session.Finalize()
	ctx.Redirect(http.StatusFound, commonPrefix+fmt.Sprint("Success: Added ", total, " accessories"))
}
