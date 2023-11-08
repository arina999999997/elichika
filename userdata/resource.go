package userdata

// this file contain all the logic for actually changing what an user own
// more precisely, it handle things found in m_content_route_guide
// things like player bonds / exp are handled elsewhere

import (
	"elichika/model"
	"elichika/utils"

	"fmt"

	"github.com/tidwall/sjson"
)

type UserResource struct {
	UserID        int   `xorm:"pk 'user_id'"`
	ContentType   int   `xorm:"pk 'content_type'"`
	ContentID     int   `xorm:"pk 'content_id'"`
	ContentAmount int64 `xorm:"'content_amount'"`
}

var (
	ResourceHandler   map[int]func(*Session, int, int, int64)
	ResourceFinalizer map[int]func(*Session, int, map[int]UserResource) (string, string)
	ResourceGroupKey  map[int]string
	ResourceItemKey   map[int]string
)

func (session *Session) AddGameMoney(money int64) {
	session.AddResource(model.Content{
		ContentType:   10,
		ContentID:     0,
		ContentAmount: money,
	})
}

func (session *Session) RemoveGameMoney(money int64) {
	session.RemoveResource(model.Content{
		ContentType:   10,
		ContentID:     0,
		ContentAmount: money,
	})
}

func (session *Session) AddCardExp(exp int64) {
	session.AddResource(model.Content{
		ContentType:   4,
		ContentID:     0,
		ContentAmount: exp,
	})
}

func (session *Session) RemoveCardExp(exp int64) {
	session.RemoveResource(model.Content{
		ContentType:   4,
		ContentID:     0,
		ContentAmount: exp,
	})
}
func (session *Session) GetUserResource(contentType, contentID int) UserResource {
	_, exists := session.UserResourceDiffs[contentType]
	if !exists {
		session.UserResourceDiffs[contentType] = make(map[int]UserResource)
	}
	resource, exists := session.UserResourceDiffs[contentType][contentID]
	if exists {
		return resource
	}
	// load from db
	exists, err := session.Db.Table("u_resource").Where("user_id = ? AND content_type = ? AND content_id = ?",
		session.UserStatus.UserID, contentType, contentID).Get(&resource)
	utils.CheckErr(err)
	if !exists {
		resource = UserResource{
			UserID:        session.UserStatus.UserID,
			ContentType:   contentType,
			ContentID:     contentID,
			ContentAmount: 100000,
		}
	}
	return resource
}

func (session *Session) UpdateUserResource(resource UserResource) {
	_, exists := session.UserResourceDiffs[resource.ContentType]
	if !exists {
		session.UserResourceDiffs[resource.ContentType] = make(map[int]UserResource)
	}
	session.UserResourceDiffs[resource.ContentType][resource.ContentID] = resource
}

func (session *Session) AddResource(resource model.Content) {
	// fmt.Println(resource)
	handler, exists := ResourceHandler[resource.ContentType]
	if !exists {
		fmt.Println("TODO: Add handler for content type ", resource.ContentType)
		return
	}
	handler(session, resource.ContentType, resource.ContentID, resource.ContentAmount)
}

func (session *Session) RemoveResource(resource model.Content) {
	// fmt.Println(resource)
	handler, exists := ResourceHandler[resource.ContentType]
	if !exists {
		fmt.Println("TODO: Add handler for content type ", resource.ContentType)
		return
	}
	handler(session, resource.ContentType, resource.ContentID, -resource.ContentAmount)
}

// return an array of pair of key / content to set that key to
// this key start from user_model / user_model_diff (so add that to the front)
// the value inside []any are PREMARSHALED json, so use sjson.SetRaw
func (session *Session) FinalizeUserResources() ([]string, []string) {
	keys := []string{}
	values := []string{}
	for contentType, resourceDiffByContentID := range session.UserResourceDiffs {
		key, value := ResourceFinalizer[contentType](session, contentType, resourceDiffByContentID)
		keys = append(keys, key)
		values = append(values, value)
	}
	return keys, values
}

func init() {
	// reference for type can be found in m_content_setting
	ResourceHandler = make(map[int]func(*Session, int, int, int64))
	ResourceHandler[1] = UserStatusResourceHandler  // sns coins (star gems)
	ResourceHandler[4] = UserStatusResourceHandler  // card exp
	ResourceHandler[10] = UserStatusResourceHandler // game money(gold)
	ResourceHandler[32] = UserStatusResourceHandler // subscription coins (purple coin)

	ResourceHandler[7] = SuitResourceHandler

	ResourceFinalizer = make(map[int]func(*Session, int, map[int]UserResource) (string, string))
	ResourceGroupKey = make(map[int]string)
	ResourceItemKey = make(map[int]string)

	ResourceHandler[5] = GenericResourceHandler // gacha point (quartz)
	ResourceFinalizer[5] = GenericResourceFinalizer
	ResourceGroupKey[5] = "user_gacha_point_by_point_id"
	ResourceItemKey[5] = "point_master_id"

	ResourceHandler[6] = GenericResourceHandler // gacha point (quartz)
	ResourceFinalizer[6] = GenericResourceFinalizer
	ResourceGroupKey[6] = "user_lesson_enhancing_item_by_item_id"
	ResourceItemKey[6] = "enhancing_item_id"

	ResourceHandler[12] = GenericResourceHandler // training items (macarons, memorials)
	ResourceFinalizer[12] = GenericResourceFinalizer
	ResourceGroupKey[12] = "user_training_material_by_item_id"
	ResourceItemKey[12] = "training_material_master_id"

	ResourceHandler[13] = GenericResourceHandler // card grade up items
	ResourceFinalizer[13] = GenericResourceFinalizer
	ResourceGroupKey[13] = "user_grade_up_item_by_item_id"
	ResourceItemKey[13] = "item_master_id"

	ResourceHandler[16] = GenericResourceHandler // training ticket
	ResourceFinalizer[16] = GenericResourceFinalizer
	ResourceGroupKey[16] = "user_recovery_ap_by_id"
	ResourceItemKey[16] = "recovery_ap_master_id"

	ResourceHandler[17] = GenericResourceHandler // candies
	ResourceFinalizer[17] = GenericResourceFinalizer
	ResourceGroupKey[17] = "user_recovery_lp_by_id"
	ResourceItemKey[17] = "recovery_lp_master_id"

	// generics exchange point (SBL / DLP)
	// also include channel exchanges
	ResourceHandler[21] = GenericResourceHandler
	ResourceFinalizer[21] = GenericResourceFinalizer
	ResourceGroupKey[21] = "user_exchange_event_point_by_id"
	ResourceItemKey[21] = "" // no need

	ResourceHandler[24] = GenericResourceHandler // accessory stickers
	ResourceFinalizer[24] = GenericResourceFinalizer
	ResourceGroupKey[24] = "user_accessory_level_up_item_by_id"
	ResourceItemKey[24] = "accessory_level_up_item_master_id"

	ResourceHandler[25] = GenericResourceHandler // accessory rarity up items
	ResourceFinalizer[25] = GenericResourceFinalizer
	ResourceGroupKey[25] = "user_accessory_rarity_up_item_by_id"
	ResourceItemKey[25] = "accessory_rarity_up_item_master_id"

	ResourceHandler[28] = GenericResourceHandler // skip tickets
	ResourceFinalizer[28] = GenericResourceFinalizer
	ResourceGroupKey[28] = "user_live_skip_ticket_by_id"
	ResourceItemKey[28] = "ticket_master_id"

	ResourceHandler[30] = GenericResourceHandler // event story unlock key
	ResourceFinalizer[30] = GenericResourceFinalizer
	ResourceGroupKey[30] = "user_story_event_unlock_item_by_id"
	ResourceItemKey[30] = "story_event_unlock_item_master_id"
}

func GenericResourceHandler(session *Session, contentType, contentID int, contentAmount int64) {
	resource := session.GetUserResource(contentType, contentID)
	resource.ContentAmount += contentAmount
	session.UpdateUserResource(resource)
}

func GenericResourceFinalizer(session *Session, contentType int,
	resourceDiffByContentID map[int]UserResource) (string, string) {
	groupKey := ResourceGroupKey[contentType]
	itemKey := ResourceItemKey[contentType]
	result := "[]"
	index := 0
	for contentID, resource := range resourceDiffByContentID {
		// update or insert the resource
		affected, err := session.Db.Table("u_resource").Where("user_id = ? AND content_type = ? AND content_id = ?",
			session.UserStatus.UserID, contentType, contentID).Update(resource)
		utils.CheckErr(err)
		if affected == 0 { // doesn't exists, insert
			session.Db.Table("u_resource").Insert(resource)
		}

		// set the json
		result, _ = sjson.Set(result, fmt.Sprintf("%d", index), contentID)
		index++
		if itemKey != "" {
			result, _ = sjson.Set(result, fmt.Sprintf("%d.%s", index, itemKey), contentID)
		}
		result, _ = sjson.Set(result, fmt.Sprintf("%d.amount", index), resource.ContentAmount)
		index++
	}
	return groupKey, result
}

func SuitResourceHandler(session *Session, suitContentType, suitMasterID int, amount int64) {
	session.InsertUserSuit(model.UserSuit{
		UserID:       session.UserStatus.UserID,
		SuitMasterID: suitMasterID,
		IsNew:        true})
}

// these resources amount are stored in the user status
func UserStatusResourceHandler(session *Session, resourceContentType, resourceContentID int, amount int64) {
	switch resourceContentType {
	case 1: // star gems
		session.UserStatus.FreeSnsCoin += int(amount)
	case 4: // card exp
		session.UserStatus.CardExp += amount
	case 10: // game money (gold)
		session.UserStatus.GameMoney += amount
	case 32: // subscription coin (purple coin)
		session.UserStatus.SubscriptionCoin += int(amount)
	default:
		fmt.Println("TODO: handle user status content type:", resourceContentType)
	}
}
