package userdata

// this file contain all the logic for actually changing what an user own
// more precisely, it handle things found in m_content_route_guide
// things like player bonds / exp are handled elsewhere

import (
	"elichika/client"
	"elichika/enum"
	"elichika/utils"

	"fmt"
	"reflect"
)

type UserResource struct {
	UserId  int            `xorm:"'user_id'"`
	Content client.Content `xorm:"extends"`
}

var (
	resourceHandler = map[int32]func(*Session, int32, int32, int32){}
	// new system
	userModelField = map[int32]string{}
)

func (session *Session) AddSnsCoin(coin int32) {
	session.UserStatus.FreeSnsCoin += coin
}

func (session *Session) RemoveSnsCoin(coin int32) {
	session.UserStatus.FreeSnsCoin -= coin
}

func (session *Session) AddGameMoney(money int32) {
	session.UserStatus.GameMoney += money
}

func (session *Session) RemoveGameMoney(money int32) {
	session.UserStatus.GameMoney -= money
}

func (session *Session) AddCardExp(exp int32) {
	session.UserStatus.CardExp += exp
}

func (session *Session) RemoveCardExp(exp int32) {
	session.UserStatus.CardExp -= exp
}

func (session *Session) GetUserResource(contentType, contentId int32) UserResource {
	_, exist := session.UserResourceDiffs[contentType]
	if !exist {
		session.UserResourceDiffs[contentType] = make(map[int32]UserResource)
	}
	resource, exist := session.UserResourceDiffs[contentType][contentId]
	if exist {
		return resource
	}
	// load from db
	exist, err := session.Db.Table("u_resource").Where("user_id = ? AND content_type = ? AND content_id = ?",
		session.UserId, contentType, contentId).Get(&resource)
	utils.CheckErr(err)
	if !exist {
		resource = UserResource{
			UserId: session.UserId,
			Content: client.Content{
				ContentType:   contentType,
				ContentId:     contentId,
				ContentAmount: 10000000,
			},
		}
	}
	return resource
}

func (session *Session) UpdateUserResource(resource UserResource) {
	_, exist := session.UserResourceDiffs[resource.Content.ContentType]
	if !exist {
		session.UserResourceDiffs[resource.Content.ContentType] = make(map[int32]UserResource)
	}
	session.UserResourceDiffs[resource.Content.ContentType][resource.Content.ContentId] = resource
}

func (session *Session) AddResource(resource client.Content) {
	handler, exist := resourceHandler[resource.ContentType]
	if !exist {
		fmt.Println("TODO: Add handler for content type ", resource.ContentType)
		return
	}
	handler(session, resource.ContentType, resource.ContentId, resource.ContentAmount)
}

func (session *Session) RemoveResource(resource client.Content) {
	handler, exist := resourceHandler[resource.ContentType]
	if !exist {
		fmt.Println("TODO: Add handler for content type ", resource.ContentType)
		return
	}
	handler(session, resource.ContentType, resource.ContentId, -resource.ContentAmount)
}

func init() {
	// reference for type can be found in m_content_setting
	resourceHandler[enum.ContentTypeSnsCoin] = userStatusResourceHandler
	resourceHandler[enum.ContentTypeCardExp] = userStatusResourceHandler
	resourceHandler[enum.ContentTypeGameMoney] = userStatusResourceHandler
	resourceHandler[enum.ContentTypeSubscriptionCoin] = userStatusResourceHandler

	resourceHandler[enum.ContentTypeSuit] = suitResourceHandler

	resourceHandler[enum.ContentTypeGachaTicket] = genericResourceHandler
	userModelField[enum.ContentTypeGachaTicket] = "UserGachaTicketByTicketId"

	resourceHandler[enum.ContentTypeGachaPoint] = genericResourceHandler // gacha point (quartz)
	userModelField[enum.ContentTypeGachaPoint] = "UserGachaPointByPointId"

	resourceHandler[enum.ContentTypeLessonEnhancingItem] = genericResourceHandler // light bulbs
	userModelField[enum.ContentTypeLessonEnhancingItem] = "UserLessonEnhancingItemByItemId"

	resourceHandler[enum.ContentTypeTrainingMaterial] = genericResourceHandler // training items (macarons, memorials)
	userModelField[enum.ContentTypeTrainingMaterial] = "UserTrainingMaterialByItemId"

	resourceHandler[enum.ContentTypeGradeUpper] = genericResourceHandler // card grade up items
	userModelField[enum.ContentTypeGradeUpper] = "UserGradeUpItemByItemId"

	resourceHandler[enum.ContentTypeRecoveryLp] = genericResourceHandler // lp candies
	userModelField[enum.ContentTypeRecoveryLp] = "UserRecoveryLpById"

	resourceHandler[enum.ContentTypeRecoveryAp] = genericResourceHandler // training ticket
	userModelField[enum.ContentTypeRecoveryAp] = "UserRecoveryApById"

	resourceHandler[enum.ContentTypeAccessoryLevelUp] = genericResourceHandler // accessory stickers
	userModelField[enum.ContentTypeAccessoryLevelUp] = "UserAccessoryLevelUpItemById"

	resourceHandler[enum.ContentTypeAccessoryRarityUp] = genericResourceHandler // accessory rarity up items
	userModelField[enum.ContentTypeAccessoryRarityUp] = "UserAccessoryRarityUpItemById"

	// generics exchange point (SBL / DLP)
	// also include channel exchanges
	resourceHandler[enum.ContentTypeExchangeEventPoint] = genericResourceHandler
	userModelField[enum.ContentTypeExchangeEventPoint] = "UserExchangeEventPointById"

	resourceHandler[enum.ContentTypeLiveSkipTicket] = genericResourceHandler // skip tickets
	userModelField[enum.ContentTypeLiveSkipTicket] = "UserLiveSkipTicketById"

	resourceHandler[enum.ContentTypeStoryEventUnlock] = genericResourceHandler // event story unlock key
	userModelField[enum.ContentTypeStoryEventUnlock] = "UserStoryEventUnlockItemById"

	resourceHandler[enum.ContentTypeEventMarathonBooster] = genericResourceHandler // marathon boosters
	userModelField[enum.ContentTypeEventMarathonBooster] = "UserEventMarathonBoosterById"

	resourceHandler[enum.ContentTypeRecoveryTowerCardUsedCount] = genericResourceHandler // dlp water bottle
	userModelField[enum.ContentTypeRecoveryTowerCardUsedCount] = "UserRecoveryTowerCardUsedCountItemByRecoveryTowerCardUsedCountItemMasterId"

	resourceHandler[enum.ContentTypeStoryMember] = memberStoryHandler
	resourceHandler[enum.ContentTypeVoice] = voiceHandler
}

func genericResourceHandler(session *Session, contentType, contentId, contentAmount int32) {
	resource := session.GetUserResource(contentType, contentId)
	resource.Content.ContentAmount += contentAmount
	session.UpdateUserResource(resource)
}

func suitResourceHandler(session *Session, _, suitMasterId, _ int32) {
	session.InsertUserSuit(suitMasterId)
}

func memberStoryHandler(session *Session, _, memberStoryId, _ int32) {
	session.InsertMemberStory(memberStoryId)
}

func voiceHandler(session *Session, _, naviVoiceMasterId, _ int32) {
	session.UpdateVoice(naviVoiceMasterId, false)
}

// these resources amount are stored in the user status
func userStatusResourceHandler(session *Session, resourceContentType, resourceContentId, amount int32) {
	switch resourceContentType {
	case enum.ContentTypeSnsCoin: // star gems
		session.AddSnsCoin(amount)
	case enum.ContentTypeCardExp: // card exp
		session.AddCardExp(amount)
	case enum.ContentTypeGameMoney: // game money (gold)
		session.AddGameMoney(amount)
	case enum.ContentTypeSubscriptionCoin: // subscription coin (purple coin)
		session.UserStatus.SubscriptionCoin += amount
	default:
		fmt.Println("TODO: handle user status content type:", resourceContentType)
	}
}

func genericResourceByResourceIdFinalizer(session *Session) {
	// TODO: user client.Content instead of UserResource
	rModel := reflect.ValueOf(&session.UserModel)
	for contentType, resourceDiffByContentId := range session.UserResourceDiffs {
		rDictionary := rModel.Elem().FieldByName(userModelField[contentType])
		if !rDictionary.IsValid() {
			fmt.Println("Invalid field: ", contentType, "->", userModelField[contentType])
			continue
		}
		rDictionaryPtrType := reflect.PointerTo(rDictionary.Type())
		rDictionarySet, ok := rDictionaryPtrType.MethodByName("Set")
		if !ok {
			panic(fmt.Sprintln("Type ", rDictionaryPtrType, " must have method Set"))
		}
		rElementPtrType := rDictionary.FieldByName("Map").Type().Elem()
		rElementType := rElementPtrType.Elem()
		rElementFromContent, ok := rElementPtrType.MethodByName("FromContent")
		if !ok {
			panic(fmt.Sprintln("Type ", rElementPtrType, " must have method FromContent"))
		}

		for _, resource := range resourceDiffByContentId {
			// update or insert the resource
			affected, err := session.Db.Table("u_resource").Where("user_id = ? AND content_type = ? AND content_id = ?",
				session.UserId, resource.Content.ContentType, resource.Content.ContentId).AllCols().Update(resource)
			utils.CheckErr(err)
			if affected == 0 { // doesn't exist, insert
				// fmt.Println("Inserted: ", resource)
				_, err = session.Db.Table("u_resource").Insert(resource)
				utils.CheckErr(err)
			}

			obj := reflect.New(rElementType)
			rElementFromContent.Func.Call([]reflect.Value{obj, reflect.ValueOf(resource.Content)})
			rDictionarySet.Func.Call([]reflect.Value{rDictionary.Addr(), reflect.ValueOf(resource.Content.ContentId), reflect.Indirect(obj)})
		}
	}
}

func genericResourceByResourceIdPopulator(session *Session) {
	rModel := reflect.ValueOf(&session.UserModel)
	contents := []client.Content{}
	// order by is not necessary
	err := session.Db.Table("u_resource").Where("user_id = ?", session.UserId).Find(&contents)
	utils.CheckErr(err)
	contentByType := map[int32][]client.Content{}
	for _, content := range contents {
		contentByType[content.ContentType] = append(contentByType[content.ContentType], content)
	}
	for contentType, fieldName := range userModelField {
		contents, exist := contentByType[contentType]
		if !exist {
			continue
		}
		rDictionary := rModel.Elem().FieldByName(fieldName)
		if !rDictionary.IsValid() {
			fmt.Println("Invalid field: ", contentType, "->", fieldName)
			continue
		}
		rDictionaryPtrType := reflect.PointerTo(rDictionary.Type())
		rDictionarySet, ok := rDictionaryPtrType.MethodByName("Set")
		if !ok {
			panic(fmt.Sprintln("Type ", rDictionaryPtrType, " must have method Set"))
		}
		rElementPtrType := rDictionary.FieldByName("Map").Type().Elem()
		rElementType := rElementPtrType.Elem()
		rElementFromContent, ok := rElementPtrType.MethodByName("FromContent")
		if !ok {
			panic(fmt.Sprintln("Type ", rElementPtrType, " must have method FromContent"))
		}
		for _, resource := range contents {
			obj := reflect.New(rElementType)
			rElementFromContent.Func.Call([]reflect.Value{obj, reflect.ValueOf(resource)})
			rDictionarySet.Func.Call([]reflect.Value{rDictionary.Addr(), reflect.ValueOf(resource.ContentId), reflect.Indirect(obj)})
		}
	}
}

// this is used for loading exported account data
func (session *Session) populateGenericResourceDiffFromUserModel() {
	rModel := reflect.ValueOf(&session.UserModel)
	for contentType, fieldName := range userModelField {
		rDictionary := rModel.Elem().FieldByName(fieldName)
		if !rDictionary.IsValid() {
			fmt.Println("Invalid field: ", contentType, "->", fieldName)
			continue
		}
		rDictionaryPtrType := reflect.PointerTo(rDictionary.Type())
		rDictionaryToContents, ok := rDictionaryPtrType.MethodByName("ToContents")
		if !ok {
			panic(fmt.Sprintln("Type ", rDictionaryPtrType, " must have method ToContents"))
		}
		contents := rDictionaryToContents.Func.Call([]reflect.Value{rDictionary.Addr()})[0].Interface().([]any)
		for _, content := range contents {
			session.UpdateUserResource(UserResource{
				UserId:  session.UserId,
				Content: content.(client.Content),
			})
		}
	}

}

func init() {
	addFinalizer(genericResourceByResourceIdFinalizer)
	addPopulator(genericResourceByResourceIdPopulator)
}
