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
)

type UserResource struct {
	UserId  int           `xorm:"'user_id'"`
	Content model.Content `xorm:"extends"`
}

var (
	resourceHandler map[int]func(*Session, int, int, int64)
	// new system
	userModelField map[int]string
)

func (session *Session) AddSnsCoin(coin int64) {
	session.AddResource(model.Content{
		ContentType:   enum.ContentTypeSnsCoin,
		ContentId:     0,
		ContentAmount: coin,
	})
}

func (session *Session) RemoveSnsCoin(coin int64) {
	session.RemoveResource(model.Content{
		ContentType:   enum.ContentTypeSnsCoin,
		ContentId:     0,
		ContentAmount: coin,
	})
}

func (session *Session) AddGameMoney(money int64) {
	session.AddResource(model.Content{
		ContentType:   enum.ContentTypeGameMoney,
		ContentId:     1200,
		ContentAmount: money,
	})
}

func (session *Session) RemoveGameMoney(money int64) {
	session.RemoveResource(model.Content{
		ContentType:   enum.ContentTypeGameMoney,
		ContentId:     1200,
		ContentAmount: money,
	})
}

func (session *Session) AddCardExp(exp int64) {
	session.AddResource(model.Content{
		ContentType:   enum.ContentTypeCardExp,
		ContentId:     1100,
		ContentAmount: exp,
	})
}

func (session *Session) RemoveCardExp(exp int64) {
	session.RemoveResource(model.Content{
		ContentType:   enum.ContentTypeCardExp,
		ContentId:     1100,
		ContentAmount: exp,
	})
}

func (session *Session) GetUserResource(contentType, contentId int) UserResource {
	_, exist := session.UserResourceDiffs[contentType]
	if !exist {
		session.UserResourceDiffs[contentType] = make(map[int]UserResource)
	}
	resource, exist := session.UserResourceDiffs[contentType][contentId]
	if exist {
		return resource
	}
	// load from db
	exist, err := session.Db.Table("u_resource").Where("user_id = ? AND content_type = ? AND content_id = ?",
		session.UserStatus.UserId, contentType, contentId).Get(&resource)
	utils.CheckErr(err)
	if !exist {
		resource = UserResource{
			UserId: session.UserStatus.UserId,
			Content: model.Content{
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
		session.UserResourceDiffs[resource.Content.ContentType] = make(map[int]UserResource)
	}
	session.UserResourceDiffs[resource.Content.ContentType][resource.Content.ContentId] = resource
}

func (session *Session) AddResource(resource model.Content) {
	handler, exist := resourceHandler[resource.ContentType]
	if !exist {
		fmt.Println("TODO: Add handler for content type ", resource.ContentType)
		return
	}
	handler(session, resource.ContentType, resource.ContentId, resource.ContentAmount)
}

func (session *Session) RemoveResource(resource model.Content) {
	handler, exist := resourceHandler[resource.ContentType]
	if !exist {
		fmt.Println("TODO: Add handler for content type ", resource.ContentType)
		return
	}
	handler(session, resource.ContentType, resource.ContentId, -resource.ContentAmount)
}

func init() {
	// reference for type can be found in m_content_setting
	resourceHandler = make(map[int]func(*Session, int, int, int64))
	resourceHandler[enum.ContentTypeSnsCoin] = userStatusResourceHandler
	resourceHandler[enum.ContentTypeCardExp] = userStatusResourceHandler
	resourceHandler[enum.ContentTypeGameMoney] = userStatusResourceHandler
	resourceHandler[enum.ContentTypeSubscriptionCoin] = userStatusResourceHandler

	resourceHandler[enum.ContentTypeSuit] = suitResourceHandler

	userModelField = make(map[int]string)

	resourceHandler[enum.ContentTypeGachaPoint] = genericResourceHandler // gacha point (quartz)
	userModelField[enum.ContentTypeGachaPoint] = "UserGachaPointByPointId"

	resourceHandler[enum.ContentTypeGachaTicket] = genericResourceHandler
	userModelField[enum.ContentTypeGachaTicket] = "UserGachaTicketByTicketId"

	resourceHandler[enum.ContentTypeLessonEnhancingItem] = genericResourceHandler // light bulbs
	userModelField[enum.ContentTypeLessonEnhancingItem] = "UserLessonEnhancingItemByItemId"

	resourceHandler[enum.ContentTypeTrainingMaterial] = genericResourceHandler // training items (macarons, memorials)
	userModelField[enum.ContentTypeTrainingMaterial] = "UserTrainingMaterialByItemId"

	resourceHandler[enum.ContentTypeGradeUpper] = genericResourceHandler // card grade up items
	userModelField[enum.ContentTypeGradeUpper] = "UserGradeUpItemByItemId"

	resourceHandler[enum.ContentTypeRecoveryAp] = genericResourceHandler // training ticket
	userModelField[enum.ContentTypeRecoveryAp] = "UserRecoveryApById"

	resourceHandler[enum.ContentTypeRecoveryLp] = genericResourceHandler // lp candies
	userModelField[enum.ContentTypeRecoveryLp] = "UserRecoveryLpById"

	// generics exchange point (SBL / DLP)
	// also include channel exchanges
	resourceHandler[enum.ContentTypeExchangeEventPoint] = genericResourceHandler
	userModelField[enum.ContentTypeExchangeEventPoint] = "UserExchangeEventPointById"

	resourceHandler[enum.ContentTypeAccessoryLevelUp] = genericResourceHandler // accessory stickers
	userModelField[enum.ContentTypeAccessoryLevelUp] = "UserAccessoryLevelUpItemById"

	resourceHandler[enum.ContentTypeAccessoryRarityUp] = genericResourceHandler // accessory rarity up items
	userModelField[enum.ContentTypeAccessoryRarityUp] = "UserAccessoryRarityUpItemById"

	resourceHandler[enum.ContentTypeEventMarathonBooster] = genericResourceHandler // marathon boosters
	userModelField[enum.ContentTypeEventMarathonBooster] = "UserEventMarathonBoosterById"

	resourceHandler[enum.ContentTypeLiveSkipTicket] = genericResourceHandler // skip tickets
	userModelField[enum.ContentTypeLiveSkipTicket] = "UserLiveSkipTicketById"

	resourceHandler[enum.ContentTypeStoryEventUnlock] = genericResourceHandler // event story unlock key
	userModelField[enum.ContentTypeStoryEventUnlock] = "UserStoryEventUnlockItemById"

	resourceHandler[enum.ContentTypeRecoveryTowerCardUsedCount] = genericResourceHandler // dlp water bottle
	userModelField[enum.ContentTypeRecoveryTowerCardUsedCount] = "UserRecoveryTowerCardUsedCountItemByRecoveryTowerCardUsedCountItemMasterId"

	resourceHandler[enum.ContentTypeStoryMember] = memberStoryHandler
	resourceHandler[enum.ContentTypeVoice] = voiceHandler
}

func genericResourceHandler(session *Session, contentType, contentId int, contentAmount int64) {
	resource := session.GetUserResource(contentType, contentId)
	resource.Content.ContentAmount += contentAmount
	session.UpdateUserResource(resource)
}

func suitResourceHandler(session *Session, _, suitMasterId int, _ int64) {
	session.InsertUserSuit(suitMasterId)
}

func memberStoryHandler(session *Session, _, memberStoryId int, _ int64) {
	session.InsertMemberStory(memberStoryId)
}

func voiceHandler(session *Session, _, naviVoiceMasterId int, _ int64) {
	session.UpdateVoice(naviVoiceMasterId, false)
}

// these resources amount are stored in the user status
func userStatusResourceHandler(session *Session, resourceContentType, resourceContentId int, amount int64) {
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

func genericResourceByResourceIdFinalizer(session *Session) {
	rModel := reflect.ValueOf(&session.UserModel)
	for contentType, resourceDiffByContentId := range session.UserResourceDiffs {
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

		for _, resource := range resourceDiffByContentId {
			// update or insert the resource
			affected, err := session.Db.Table("u_resource").Where("user_id = ? AND content_type = ? AND content_id = ?",
				session.UserStatus.UserId, resource.Content.ContentType, resource.Content.ContentId).AllCols().Update(resource)
			utils.CheckErr(err)
			if affected == 0 { // doesn't exist, insert
				// fmt.Println("Inserted: ", resource)
				_, err = session.Db.Table("u_resource").Insert(resource)
				utils.CheckErr(err)
			}

			obj := reflect.New(rElementType)
			rElementFromContent.Func.Call([]reflect.Value{obj, reflect.ValueOf(resource.Content)})
			rObjectPushBack.Func.Call([]reflect.Value{rObject.Addr(), reflect.Indirect(obj)})
		}
	}
}

func genericResourceByResourceIdPopulator(session *Session) {
	rModel := reflect.ValueOf(&session.UserModel)
	contents := []model.Content{}
	// order by is not necessary
	err := session.Db.Table("u_resource").Where("user_id = ?", session.UserStatus.UserId).Find(&contents)
	utils.CheckErr(err)
	contentByType := make(map[int][]model.Content)
	for _, content := range contents {
		contentByType[content.ContentType] = append(contentByType[content.ContentType], content)
	}
	for contentType, fieldName := range userModelField {
		contents, exist := contentByType[contentType]
		if !exist {
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

// this is used for loading exported account data
func (session *Session) populateGenericResourceDiffFromUserModel() {
	rModel := reflect.ValueOf(&session.UserModel)
	for contentType, fieldName := range userModelField {
		rObject := rModel.Elem().FieldByName(fieldName)
		if !rObject.IsValid() {
			fmt.Println("Invalid field: ", contentType, "->", fieldName)
			continue
		}
		rElementType := rObject.FieldByName("Objects").Type().Elem()
		if rElementType == reflect.ValueOf(0).Type() {
			fmt.Println("Not handled: ", contentType, fieldName)
			continue
		}
		rObjectPtrType := reflect.PointerTo(rObject.Type())
		rObjectToContents, ok := rObjectPtrType.MethodByName("ToContents")
		if !ok {
			panic(fmt.Sprintln("Type ", rObjectPtrType, " must have method ToContents"))
		}
		contents := rObjectToContents.Func.Call([]reflect.Value{rObject.Addr()})[0].Interface().([]any)
		for _, content := range contents {
			session.UpdateUserResource(UserResource{
				UserId:  session.UserStatus.UserId,
				Content: content.(model.Content),
			})
		}
	}

}

func init() {
	addFinalizer(genericResourceByResourceIdFinalizer)
	addPopulator(genericResourceByResourceIdPopulator)
}
