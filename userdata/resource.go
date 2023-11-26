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
	UserID  int           `xorm:"'user_id'"`
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
	_, exist := session.UserResourceDiffs[contentType]
	if !exist {
		session.UserResourceDiffs[contentType] = make(map[int]UserResource)
	}
	resource, exist := session.UserResourceDiffs[contentType][contentID]
	if exist {
		return resource
	}
	// load from db
	exist, err := session.Db.Table("u_resource").Where("user_id = ? AND content_type = ? AND content_id = ?",
		session.UserStatus.UserID, contentType, contentID).Get(&resource)
	utils.CheckErr(err)
	if !exist {
		resource = UserResource{
			UserID: session.UserStatus.UserID,
			Content: model.Content{
				ContentType:   contentType,
				ContentID:     contentID,
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
	session.UserResourceDiffs[resource.Content.ContentType][resource.Content.ContentID] = resource
}

func (session *Session) AddResource(resource model.Content) {
	handler, exist := resourceHandler[resource.ContentType]
	if !exist {
		fmt.Println("TODO: Add handler for content type ", resource.ContentType)
		return
	}
	handler(session, resource.ContentType, resource.ContentID, resource.ContentAmount)
}

func (session *Session) RemoveResource(resource model.Content) {
	handler, exist := resourceHandler[resource.ContentType]
	if !exist {
		fmt.Println("TODO: Add handler for content type ", resource.ContentType)
		return
	}
	handler(session, resource.ContentType, resource.ContentID, -resource.ContentAmount)
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
	userModelField[enum.ContentTypeGachaPoint] = "UserGachaPointByPointID"

	resourceHandler[enum.ContentTypeGachaTicket] = genericResourceHandler
	userModelField[enum.ContentTypeGachaTicket] = "UserGachaTicketByTicketID"

	resourceHandler[enum.ContentTypeLessonEnhancingItem] = genericResourceHandler // light bulbs
	userModelField[enum.ContentTypeLessonEnhancingItem] = "UserLessonEnhancingItemByItemID"

	resourceHandler[enum.ContentTypeTrainingMaterial] = genericResourceHandler // training items (macarons, memorials)
	userModelField[enum.ContentTypeTrainingMaterial] = "UserTrainingMaterialByItemID"

	resourceHandler[enum.ContentTypeCardExchange] = genericResourceHandler // card grade up items
	userModelField[enum.ContentTypeCardExchange] = "UserGradeUpItemByItemID"

	resourceHandler[enum.ContentTypeSheetRecoveryAP] = genericResourceHandler // training ticket
	userModelField[enum.ContentTypeSheetRecoveryAP] = "UserRecoveryApByID"

	resourceHandler[enum.ContentTypeRecoveryLP] = genericResourceHandler // lp candies
	userModelField[enum.ContentTypeRecoveryLP] = "UserRecoveryLpByID"

	// generics exchange point (SBL / DLP)
	// also include channel exchanges
	resourceHandler[enum.ContentTypeExchangeEventPoint] = genericResourceHandler
	userModelField[enum.ContentTypeExchangeEventPoint] = "UserExchangeEventPointByID"

	resourceHandler[enum.ContentTypeAccessoryLevelUpItem] = genericResourceHandler // accessory stickers
	userModelField[enum.ContentTypeAccessoryLevelUpItem] = "UserAccessoryLevelUpItemByID"

	resourceHandler[enum.ContentTypeAccessoryRarityUpItem] = genericResourceHandler // accessory rarity up items
	userModelField[enum.ContentTypeAccessoryRarityUpItem] = "UserAccessoryRarityUpItemByID"

	resourceHandler[enum.ContentTypeEventMarathonBooster] = genericResourceHandler // marathon boosters
	userModelField[enum.ContentTypeEventMarathonBooster] = "UserEventMarathonBoosterByID"

	resourceHandler[enum.ContentTypeSkipTicket] = genericResourceHandler // skip tickets
	userModelField[enum.ContentTypeSkipTicket] = "UserLiveSkipTicketByID"

	resourceHandler[enum.ContentTypeStoryEventUnlock] = genericResourceHandler // event story unlock key
	userModelField[enum.ContentTypeStoryEventUnlock] = "UserStoryEventUnlockItemByID"

	resourceHandler[enum.ContentTypeTowerRecoveryItem] = genericResourceHandler // dlp water bottle
	userModelField[enum.ContentTypeTowerRecoveryItem] = "UserRecoveryTowerCardUsedCountItemByRecoveryTowerCardUsedCountItemMasterID"

	resourceHandler[enum.ContentTypeStoryMember] = memberStoryHandler
	resourceHandler[enum.ContentTypeVoice] = voiceHandler
}

func genericResourceHandler(session *Session, contentType, contentID int, contentAmount int64) {
	resource := session.GetUserResource(contentType, contentID)
	resource.Content.ContentAmount += contentAmount
	session.UpdateUserResource(resource)
}

func suitResourceHandler(session *Session, _, suitMasterID int, _ int64) {
	session.InsertUserSuit(suitMasterID)
}

func memberStoryHandler(session *Session, _, memberStoryID int, _ int64) {
	session.InsertMemberStory(memberStoryID)
}

func voiceHandler(session *Session, _, naviVoiceMasterID int, _ int64) {
	session.UpdateVoice(naviVoiceMasterID, false)
}

// these resources amount are stored in the user status
func userStatusResourceHandler(session *Session, resourceContentType, resourceContentID int, amount int64) {
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

		for _, resource := range resourceDiffByContentID {
			// update or insert the resource
			affected, err := session.Db.Table("u_resource").Where("user_id = ? AND content_type = ? AND content_id = ?",
				session.UserStatus.UserID, resource.Content.ContentType, resource.Content.ContentID).AllCols().Update(resource)
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

func genericResourceByResourceIDPopulator(session *Session) {
	rModel := reflect.ValueOf(&session.UserModel)
	contents := []model.Content{}
	// order by is not necessary
	err := session.Db.Table("u_resource").Where("user_id = ?", session.UserStatus.UserID).Find(&contents)
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
		contents := rObjectToContents.Func.Call([]reflect.Value{rObject.Addr()})[0].Interface().([](model.Content))
		for _, content := range contents {

			session.UpdateUserResource(UserResource{
				UserID:  session.UserStatus.UserID,
				Content: content,
			})
		}
	}

}

func init() {
	addFinalizer(genericResourceByResourceIDFinalizer)
	addPopulator(genericResourceByResourceIDPopulator)
}
