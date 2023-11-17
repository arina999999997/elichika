package userdata

// this file contain all the logic for actually changing what an user own
// more precisely, it handle things found in m_content_route_guide
// things like player bonds / exp are handled elsewhere

import (
	"elichika/enum"
	"elichika/model"
	"elichika/utils"

	"fmt"
	"reflect"
	"sort"

	"github.com/tidwall/sjson"
)

type UserResource struct {
	UserID   int           `xorm:"pk 'user_id'"`
	Resource model.Content `xorm:"extends"`
	// ContentType   int   `xorm:"pk 'content_type'"`
	// ContentID     int   `xorm:"pk 'content_id'"`
	// ContentAmount int64 `xorm:"'content_amount'"`
}

var (
	ResourceHandler   map[int]func(*Session, int, int, int64)
	// old system
	ResourceFinalizer map[int]func(*Session, int, map[int]UserResource) (string, string)
	ResourceGroupKey  map[int]string
	ResourceItemKey   map[int]string

	// new system
	userModelField map[int]string
)

func (session *Session) AddSnsCoin(coin int64) {
	session.AddResource(model.Content{
		ContentType:   enum.ContentTypeSnsCoin,
		ContentID:     0,
		ContentAmount: coin,
	})
}

func (session *Session) RemoveSnsCoin(coin int64) {
	session.RemoveResource(model.Content{
		ContentType:   enum.ContentTypeSnsCoin,
		ContentID:     0,
		ContentAmount: coin,
	})
}

func (session *Session) AddGameMoney(money int64) {
	session.AddResource(model.Content{
		ContentType:   enum.ContentTypeGameMoney,
		ContentID:     1200,
		ContentAmount: money,
	})
}

func (session *Session) RemoveGameMoney(money int64) {
	session.RemoveResource(model.Content{
		ContentType:   enum.ContentTypeGameMoney,
		ContentID:     1200,
		ContentAmount: money,
	})
}

func (session *Session) AddCardExp(exp int64) {
	session.AddResource(model.Content{
		ContentType:   enum.ContentTypeCardExp,
		ContentID:     1100,
		ContentAmount: exp,
	})
}

func (session *Session) RemoveCardExp(exp int64) {
	session.RemoveResource(model.Content{
		ContentType:   enum.ContentTypeCardExp,
		ContentID:     1100,
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
			UserID: session.UserStatus.UserID,
			Resource: model.Content{
				ContentType:   contentType,
				ContentID:     contentID,
				ContentAmount: 100000,
			},
		}
	}
	return resource
}

func (session *Session) UpdateUserResource(resource UserResource) {
	_, exists := session.UserResourceDiffs[resource.Resource.ContentType]
	if !exists {
		session.UserResourceDiffs[resource.Resource.ContentType] = make(map[int]UserResource)
	}
	session.UserResourceDiffs[resource.Resource.ContentType][resource.Resource.ContentID] = resource
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
	ResourceHandler[enum.ContentTypeSnsCoin] = UserStatusResourceHandler
	ResourceHandler[enum.ContentTypeCardExp] = UserStatusResourceHandler
	ResourceHandler[enum.ContentTypeGameMoney] = UserStatusResourceHandler
	ResourceHandler[enum.ContentTypeSubscriptionCoin] = UserStatusResourceHandler

	ResourceHandler[enum.ContentTypeSuit] = SuitResourceHandler

	ResourceFinalizer = make(map[int]func(*Session, int, map[int]UserResource) (string, string))
	ResourceGroupKey = make(map[int]string)
	ResourceItemKey = make(map[int]string)
	userModelField = make(map[int]string)

	ResourceHandler[enum.ContentTypeGachaPoint] = GenericResourceHandler // gacha point (quartz)
	ResourceFinalizer[enum.ContentTypeGachaPoint] = GenericResourceFinalizer
	ResourceGroupKey[enum.ContentTypeGachaPoint] = "user_gacha_point_by_point_id"
	ResourceItemKey[enum.ContentTypeGachaPoint] = "point_master_id"
	userModelField[enum.ContentTypeGachaPoint] = "UserGachaPointByPointID"

	ResourceHandler[enum.ContentTypeLessonEnhancingItem] = GenericResourceHandler // light bulbs
	ResourceFinalizer[enum.ContentTypeLessonEnhancingItem] = GenericResourceFinalizer
	ResourceGroupKey[enum.ContentTypeLessonEnhancingItem] = "user_lesson_enhancing_item_by_item_id"
	ResourceItemKey[enum.ContentTypeLessonEnhancingItem] = "enhancing_item_id"
	userModelField[enum.ContentTypeLessonEnhancingItem] = "UserLessonEnhancingItemByItemID"

	ResourceHandler[enum.ContentTypeTrainingMaterial] = GenericResourceHandler // training items (macarons, memorials)
	ResourceFinalizer[enum.ContentTypeTrainingMaterial] = GenericResourceFinalizer
	ResourceGroupKey[enum.ContentTypeTrainingMaterial] = "user_training_material_by_item_id"
	ResourceItemKey[enum.ContentTypeTrainingMaterial] = "training_material_master_id"
	userModelField[enum.ContentTypeTrainingMaterial] = "UserTrainingMaterialByItemID"

	ResourceHandler[enum.ContentTypeCardExchange] = GenericResourceHandler // card grade up items
	ResourceFinalizer[enum.ContentTypeCardExchange] = GenericResourceFinalizer
	ResourceGroupKey[enum.ContentTypeCardExchange] = "user_grade_up_item_by_item_id"
	ResourceItemKey[enum.ContentTypeCardExchange] = "item_master_id"
	userModelField[enum.ContentTypeCardExchange] = "UserGradeUpItemByItemID"

	ResourceHandler[enum.ContentTypeSheetRecoveryAP] = GenericResourceHandler // training ticket
	ResourceFinalizer[enum.ContentTypeSheetRecoveryAP] = GenericResourceFinalizer
	ResourceGroupKey[enum.ContentTypeSheetRecoveryAP] = "user_recovery_ap_by_id"
	ResourceItemKey[enum.ContentTypeSheetRecoveryAP] = "recovery_ap_master_id"
	userModelField[enum.ContentTypeSheetRecoveryAP] = "UserRecoveryApByID"

	ResourceHandler[enum.ContentTypeRecoveryLP] = GenericResourceHandler // lp candies
	ResourceFinalizer[enum.ContentTypeRecoveryLP] = GenericResourceFinalizer
	ResourceGroupKey[enum.ContentTypeRecoveryLP] = "user_recovery_lp_by_id"
	ResourceItemKey[enum.ContentTypeRecoveryLP] = "recovery_lp_master_id"
	userModelField[enum.ContentTypeRecoveryLP] = "UserRecoveryLpByID"

	// generics exchange point (SBL / DLP)
	// also include channel exchanges
	ResourceHandler[enum.ContentTypeExchangeEventPoint] = GenericResourceHandler
	ResourceFinalizer[enum.ContentTypeExchangeEventPoint] = GenericResourceFinalizer
	ResourceGroupKey[enum.ContentTypeExchangeEventPoint] = "user_exchange_event_point_by_id"
	ResourceItemKey[enum.ContentTypeExchangeEventPoint] = "" // no need
	userModelField[enum.ContentTypeExchangeEventPoint] = "UserExchangeEventPointByID"

	ResourceHandler[enum.ContentTypeAccessoryLevelUpItem] = GenericResourceHandler // accessory stickers
	ResourceFinalizer[enum.ContentTypeAccessoryLevelUpItem] = GenericResourceFinalizer
	ResourceGroupKey[enum.ContentTypeAccessoryLevelUpItem] = "user_accessory_level_up_item_by_id"
	ResourceItemKey[enum.ContentTypeAccessoryLevelUpItem] = "accessory_level_up_item_master_id"
	userModelField[enum.ContentTypeAccessoryLevelUpItem] = "UserAccessoryLevelUpItemByID"

	ResourceHandler[enum.ContentTypeAccessoryRarityUpItem] = GenericResourceHandler // accessory rarity up items
	ResourceFinalizer[enum.ContentTypeAccessoryRarityUpItem] = GenericResourceFinalizer
	ResourceGroupKey[enum.ContentTypeAccessoryRarityUpItem] = "user_accessory_rarity_up_item_by_id"
	ResourceItemKey[enum.ContentTypeAccessoryRarityUpItem] = "accessory_rarity_up_item_master_id"
	userModelField[enum.ContentTypeAccessoryRarityUpItem] = "UserAccessoryRarityUpItemByID"

	ResourceHandler[enum.ContentTypeSkipTicket] = GenericResourceHandler // skip tickets
	ResourceFinalizer[enum.ContentTypeSkipTicket] = GenericResourceFinalizer
	ResourceGroupKey[enum.ContentTypeSkipTicket] = "user_live_skip_ticket_by_id"
	ResourceItemKey[enum.ContentTypeSkipTicket] = "ticket_master_id"
	userModelField[enum.ContentTypeSkipTicket] = "UserLiveSkipTicketByID"

	ResourceHandler[enum.ContentTypeStoryEventUnLock] = GenericResourceHandler // event story unlock key
	ResourceFinalizer[enum.ContentTypeStoryEventUnLock] = GenericResourceFinalizer
	ResourceGroupKey[enum.ContentTypeStoryEventUnLock] = "user_story_event_unlock_item_by_id"
	ResourceItemKey[enum.ContentTypeStoryEventUnLock] = "story_event_unlock_item_master_id"
	userModelField[enum.ContentTypeStoryEventUnLock] = "UserStoryEventUnlockItemByID"

	ResourceHandler[enum.ContentTypeStoryMember] = memberStoryHandler
	ResourceHandler[enum.ContentTypeVoice] = voiceHandler
}

func GenericResourceHandler(session *Session, contentType, contentID int, contentAmount int64) {
	resource := session.GetUserResource(contentType, contentID)
	resource.Resource.ContentAmount += contentAmount
	session.UpdateUserResource(resource)
}

func GenericResourceFinalizer(session *Session, contentType int,
	resourceDiffByContentID map[int]UserResource) (string, string) {
	groupKey := ResourceGroupKey[contentType]
	itemKey := ResourceItemKey[contentType]
	result := "[]"
	index := 0

	keys := []int{}
	for key := range resourceDiffByContentID {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, contentID := range keys {
		resource := resourceDiffByContentID[contentID]
		// update or insert the resource
		affected, err := session.Db.Table("u_resource").Where("user_id = ? AND content_type = ? AND content_id = ?",
			session.UserStatus.UserID, contentType, contentID).AllCols().Update(resource)
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
		result, _ = sjson.Set(result, fmt.Sprintf("%d.amount", index), resource.Resource.ContentAmount)
		index++
	}
	return groupKey, result
}

func SuitResourceHandler(session *Session, _, suitMasterID int, _ int64) {
	session.InsertUserSuit(suitMasterID)
}

func memberStoryHandler(session *Session, _, memberStoryID int, _ int64) {
	session.InsertMemberStory(memberStoryID)
}

func voiceHandler(session *Session, _, naviVoiceMasterID int, _ int64) {
	session.UpdateVoice(naviVoiceMasterID, false)
}

// these resources amount are stored in the user status
func UserStatusResourceHandler(session *Session, resourceContentType, resourceContentID int, amount int64) {
	switch resourceContentType {
	case enum.ContentTypeSnsCoin: // star gems
		session.UserStatus.FreeSnsCoin += int(amount)
	case enum.ContentTypeCardExp: // card exp
		session.UserStatus.CardExp += amount
	case enum.ContentTypeGameMoney: // game money (gold)
		session.UserStatus.GameMoney += amount
	case enum.ContentTypeSubscriptionCoin: // subscription coin (purple coin)
		session.UserStatus.SubscriptionCoin += int(amount)
	default:
		fmt.Println("TODO: handle user status content type:", resourceContentType)
	}
}

func genericResourceByResourceIDFinalizer(session *Session) {
	rModel := reflect.ValueOf(&session.UserModel)
	for contentType, resourceDiffByContentID := range session.UserResourceDiffs {
		rObject := rModel.Elem().FieldByName(userModelField[contentType])
		if !rObject.IsValid() {
			fmt.Println("Invalid field: ", contentType, "->", userModelField[contentType])
			continue
		}
		rObjectPtrType := reflect.PointerTo(rObject.Type())
		rObjectPushBack, ok := rObjectPtrType.MethodByName("PushBack")
		if !ok {
			panic(fmt.Sprintln("Type ", rObjectPtrType, " must have method PushBack"))
		}
		rElementType := rObject.FieldByName("Objects").Type().Elem()
		if rElementType == reflect.ValueOf(0).Type() {
			fmt.Println("Not handled: ", contentType, userModelField[contentType])
			continue
		}
		rElementPtrType := reflect.PointerTo(rElementType)
		rElementFromContent, ok := rElementPtrType.MethodByName("FromContent")
		if !ok {
			panic(fmt.Sprintln("Type ", rElementPtrType, " must have method FromContent"))
		}
		// this is to produce a consistent order, to check against the other method
		// TODO: remove once no longer necessary
		keys := []int{}
		for key := range resourceDiffByContentID {
			keys = append(keys, key)
		}
		sort.Ints(keys)
		for _, key := range keys {
			resource := resourceDiffByContentID[key]
			obj := reflect.New(rElementType)
			rElementFromContent.Func.Call([]reflect.Value{obj, reflect.ValueOf(resource.Resource)})
			rObjectPushBack.Func.Call([]reflect.Value{rObject.Addr(), reflect.Indirect(obj)})
		}
	}
}

func init() {
	addFinalizer(genericResourceByResourceIDFinalizer)
}

func genericResourceByResourceIDPopulator(session *Session) {
	rModel := reflect.ValueOf(&session.UserModel)
	contents := []model.Content{}
	// order by is not necessary
	err := session.Db.Table("u_resource").Where("user_id = ?", session.UserStatus.UserID).
		OrderBy("content_type, content_id").Find(&contents)
	utils.CheckErr(err)
	contentByType := make(map[int][]model.Content)
	for _, content := range contents {
		contentByType[content.ContentType] = append(contentByType[content.ContentType], content)
	}
	for contentType, fieldName := range userModelField {
		contents, exists := contentByType[contentType]
		if !exists {
			continue
		}
		rObject := rModel.Elem().FieldByName(fieldName)
		if !rObject.IsValid() {
			fmt.Println("Invalid field: ", contentType, "->", fieldName)
			continue
		}
		rObjectPtrType := reflect.PointerTo(rObject.Type())
		rObjectPushBack, ok := rObjectPtrType.MethodByName("PushBack")
		if !ok {
			panic(fmt.Sprintln("Type ", rObjectPtrType, " must have method PushBack"))
		}
		rElementType := rObject.FieldByName("Objects").Type().Elem()
		if rElementType == reflect.ValueOf(0).Type() {
			fmt.Println("Not handled: ", contentType, fieldName)
			continue
		}
		rElementPtrType := reflect.PointerTo(rElementType)
		rElementFromContent, ok := rElementPtrType.MethodByName("FromContent")
		if !ok {
			panic(fmt.Sprintln("Type ", rElementPtrType, " must have method FromContent"))
		}
		for _, resource := range contents {
			obj := reflect.New(rElementType)
			rElementFromContent.Func.Call([]reflect.Value{obj, reflect.ValueOf(resource)})
			rObjectPushBack.Func.Call([]reflect.Value{rObject.Addr(), reflect.Indirect(obj)})
		}
	}
}

func init() {
	addPopulator(genericResourceByResourceIDPopulator)
}
