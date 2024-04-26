package user

import (
	"elichika/client"
	"elichika/enum"
	"elichika/locale"
	"elichika/router"
	"elichika/subsystem/user_present"
	"elichika/userdata"
	"elichika/webui/webui_utils"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AccessoryHandler(ctx *gin.Context) {
	session := ctx.MustGet("session").(*userdata.Session)

	accessoryIds := []int32{}
	getRarity := make(map[int32]bool)
	getRarity[enum.AccessoryRarityRare] = webui_utils.GetFormBool(ctx, "r")
	getRarity[enum.AccessoryRaritySRare] = webui_utils.GetFormBool(ctx, "sr")
	getRarity[enum.AccessoryRarityURare] = webui_utils.GetFormBool(ctx, "ur")
	for _, accessory := range session.Gamedata.Accessory {
		if getRarity[accessory.RarityType] {
			accessoryIds = append(accessoryIds, accessory.Id)
		}
	}
	if len(accessoryIds) == 0 {
		webui_utils.CommonResponse(ctx, "No accessory to add, choose at least 1 rarity", "")
		return
	}
	amount := webui_utils.GetFormInt32(ctx, "amount")
	for _, accessoryMasterId := range accessoryIds {
		user_present.AddPresentWithDuration(session, client.PresentItem{
			Content: client.Content{
				ContentType:   enum.ContentTypeAccessory,
				ContentId:     accessoryMasterId,
				ContentAmount: amount,
			},
			PresentRouteType: enum.PresentRouteTypeAdminPresent,
		}, 3600)
	}
	session.Finalize()
	webui_utils.CommonResponse(ctx, "Added accessories to the present box, they will expire in 1 hour so be sure to get them", "")
}

type ResourceHelper struct {
	ResourceName string // used as the main label of the feature
	Contents     []*client.Content
	MaxAllowed   int32
}

var resourceHelpers = []ResourceHelper{}

func AddResourceHelper(resourceName string, contents []*client.Content, maxAllowed int32) {
	helper := ResourceHelper{
		ResourceName: resourceName,
		Contents:     contents,
		MaxAllowed:   maxAllowed,
	}
	resourceHelpers = append(resourceHelpers, helper)
}

func (rh *ResourceHelper) ToHTML(id int) string {
	html := fmt.Sprintf(`<div>
	<form id="form_%d" method="POST" enctype="multipart/form-data">
		<label>%s</label>
		<input type="number" name="amount" min="1" max="%d">
		<input type="hidden" name="id" value="%d">
		<input type="button" onclick="submit_form('form_%d', './resource_helper')" value="Add">
	</form>
</div>`, id, rh.ResourceName, rh.MaxAllowed, id, id)
	return html
}

func resourceHelperForm(ctx *gin.Context) {
	form :=
		`<div><label>Add some resource to make your playthrough easier / faster.</label></div>
<div><label>You can (or at least will be able to) get all the (relevant) resource from the game, but some will take forever (like signed boards).</label></div>
<div><label>To use, just click on the add button after setting the amount</label></div>
<div><label>Don't be too greedy or your account might become unplayable due to overflow.</label></div>
<div><label>If you want to get things done, you can try the account builder instead!</label></div>
<br>
`

	for id, helper := range resourceHelpers {
		form += helper.ToHTML(id)
		form += "\n"
	}

	form += `
	<div><form id="accessory">Add accessories of selected rarities
	<input type="checkbox" name="r"> R 
	<input type="checkbox" name="sr"> SR 
	<input type="checkbox" name="ur"> UR 
	amount of each accessory: <input type="number" name="amount" min="1" max="100">
	<input type="button" onclick="submit_form('accessory', './accessory')" value="Add">
	</form>
	</div>
	`
	ctx.HTML(http.StatusOK, "logged_in_user.html", gin.H{
		"body": form,
	})
}

func resourceHelperHandler(ctx *gin.Context) {
	id := webui_utils.GetFormInt32(ctx, "id")
	amount := webui_utils.GetFormInt32(ctx, "amount")
	session := ctx.MustGet("session").(*userdata.Session)
	for _, item := range resourceHelpers[id].Contents {
		user_present.AddPresentWithDuration(session, client.PresentItem{
			Content:          item.Amount(amount),
			PresentRouteType: enum.PresentRouteTypeAdminPresent,
		}, 3600)
	}
	session.Finalize()
	webui_utils.CommonResponse(ctx, "Added resources to your present box, they will expire in 1 hour so be sure to claim them", "")
}

func init() {
	addFeature("Resource helper", "resource_helper")
	router.AddHandler("/webui/user", "GET", "/resource_helper", resourceHelperForm)
	router.AddHandler("/webui/user", "POST", "/resource_helper", resourceHelperHandler)
	router.AddHandler("/webui/user", "POST", "/accessory", AccessoryHandler)
	gamedata := locale.Locales["en"].Gamedata

	AddResourceHelper("Add training materials (including memorials and signed boards)", gamedata.ContentsByContentType[enum.ContentTypeTrainingMaterial], 1000000)
	AddResourceHelper("Add grade up items", gamedata.ContentsByContentType[enum.ContentTypeGradeUpper], 100000)
	AddResourceHelper("Add training tickets", gamedata.ContentsByContentType[enum.ContentTypeRecoveryAp], 10000)
	AddResourceHelper("Add LP candies", gamedata.ContentsByContentType[enum.ContentTypeRecoveryLp], 10000)
	AddResourceHelper("Add exchange currencies", gamedata.ContentsByContentType[enum.ContentTypeExchangeEventPoint], 10000000)
	AddResourceHelper("Add accessory practice items", gamedata.ContentsByContentType[enum.ContentTypeAccessoryLevelUp], 1000000)
	AddResourceHelper("Add accessory rarity up items", gamedata.ContentsByContentType[enum.ContentTypeAccessoryLevelUp], 1000000)
	AddResourceHelper("Add skip ticket", gamedata.ContentsByContentType[enum.ContentTypeLiveSkipTicket], 1000000)
	AddResourceHelper("Add memory key (open old event story)", gamedata.ContentsByContentType[enum.ContentTypeStoryEventUnlock], 10000)
	AddResourceHelper("Add water bottle (recover DLP PP)", gamedata.ContentsByContentType[enum.ContentTypeRecoveryTowerCardUsedCount], 10000)
}
