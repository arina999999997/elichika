package user

import (
	"elichika/client"
	"elichika/enum"
	"elichika/router"
	"elichika/subsystem/user_present"
	"elichika/userdata"
	"elichika/webui/webui_utils"

	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func addPresentHandler(ctx *gin.Context) {
	session := ctx.MustGet("session").(*userdata.Session)

	session.UserStatus.LastLoginAt = time.Now().Unix()
	contentType, existA := webui_utils.SafeGetFormInt32(ctx, "content_type")
	contentId, existB := webui_utils.SafeGetFormInt32(ctx, "content_id")
	if !(existA && existB) {
		webui_utils.CommonResponse(ctx, "Content Type and or Content Id is wrong!", "")
	}
	contentAmount, existC := webui_utils.SafeGetFormInt32(ctx, "content_amount")

	if !existC {
		contentAmount = 1
	}

	user_present.AddPresentWithDuration(session, client.PresentItem{
		PresentRouteType: enum.PresentRouteTypeAdminPresent,
		Content: client.Content{
			ContentType:   int32(contentType),
			ContentId:     int32(contentId),
			ContentAmount: int32(contentAmount),
		}}, 3600)

	session.Finalize()
	webui_utils.CommonResponse(ctx, fmt.Sprintf("Added item (%d %d) %d times to present box. It will expire in 1 hour, so be sure to collect it.", contentType, contentId, contentAmount), "")
}

func addPresentForm(ctx *gin.Context) {
	form := `
<div><label>This feature requires understanding of the item id system of the game, DO NOT use it if you don't know what you are doing.</div></label>
<div><label>Even if you know what you're doing, it's still a good idea to make a backup.</div></label>
<div><label>Finally, if you added the wrong item, then you can wait for it to expire and you should have no problem with the present box.</div></label>
<form id="add_present">
	<label>Content type:<label>
	<select name="content_type" id="content_type">
	<option value="1"> SnsCoin (1, 0x01) </option>
	<option value="3"> Card (3, 0x03) </option>
	<option value="4"> CardExp (4, 0x04) </option>
	<option value="5"> GachaPoint (5, 0x05) </option>
	<option value="6"> LessonEnhancingItem (6, 0x06) </option>
	<option value="7"> Suit (7, 0x07) </option>
	<option value="8"> Voice (8, 0x08) </option>
	<option value="9"> GachaTicket (9, 0x09) </option>
	<option value="10"> GameMoney (10, 0x0a) </option>
	<option value="12"> TrainingMaterial (12, 0x0c) </option>
	<option value="13"> GradeUpper (13, 0x0d) </option>
	<option value="14"> GiftBox (14, 0x0e) </option>
	<option value="15"> Emblem (15, 0x0f) </option>
	<option value="16"> RecoveryAp (16, 0x10) </option>
	<option value="17"> RecoveryLp (17, 0x11) </option>
	<option value="19"> StorySide (19, 0x13) </option>
	<option value="20"> StoryMember (20, 0x14) </option>
	<option value="21"> ExchangeEventPoint (21, 0x15) </option>
	<option value="23"> Accessory (23, 0x17) </option>
	<option value="24"> AccessoryLevelUp (24, 0x18) </option>
	<option value="25"> AccessoryRarityUp (25, 0x19) </option>
	<option value="26"> CustomBackground (26, 0x1a) </option>
	<option value="27"> EventMarathonBooster (27, 0x1b) </option>
	<option value="28"> LiveSkipTicket (28, 0x1c) </option>
	<option value="29"> EventMiningBooster (29, 0x1d) </option>
	<option value="30"> StoryEventUnlock (30, 0x1e) </option>
	<option value="31"> RecoveryTowerCardUsedCount (31, 0x1f) </option>
	<option value="32"> SubscriptionCoin (32, 0x20) </option>
	<option value="33"> MemberGuildSupport (33, 0x21) </option>
</select>
<label>Content id:<label>
<input type="number" name="content_id" min="0" max="2147483647">
<label>Content amount:<label>
<input type="number" name="content_amount" min="0" max="2147483647">
<input type="button" value="Add" onclick="submit_form('add_present', './add_present')">
</form>
`

	ctx.HTML(http.StatusOK, "logged_in_user.html", gin.H{
		"body": form,
	})
}

func init() {
	addFeature("(Advanced) Add item using id", "add_present")
	router.AddHandler("/webui/user", "GET", "/add_present", addPresentForm)
	router.AddHandler("/webui/user", "POST", "/add_present", addPresentHandler)
}
