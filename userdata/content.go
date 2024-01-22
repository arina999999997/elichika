package userdata

// this file contain all the logic for actually changing what an user own
// more precisely, it handle things found in m_content_route_guide
// things like player bonds / exp are handled elsewhere

import (
	"elichika/client"
	"elichika/config"
	"elichika/enum"
	"elichika/utils"

	"fmt"
	"reflect"
)

var (
	contentHandler = map[int32]func(*Session, int32, int32, int32){}
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

func (session *Session) GetUserContent(contentType, contentId int32) client.Content {
	_, exist := session.UserContentDiffs[contentType]
	if !exist {
		session.UserContentDiffs[contentType] = make(map[int32]client.Content)
	}
	content, exist := session.UserContentDiffs[contentType][contentId]
	if exist {
		return content
	}
	// load from db
	exist, err := session.Db.Table("u_content").Where("user_id = ? AND content_type = ? AND content_id = ?",
		session.UserId, contentType, contentId).Get(&content)
	utils.CheckErr(err)
	if !exist {
		content = client.Content{
			ContentType:   contentType,
			ContentId:     contentId,
			ContentAmount: *config.Conf.DefaultContentAmount,
		}
	}
	return content
}

func (session *Session) GetUserContentByContent(content client.Content) client.Content {
	return session.GetUserContent(content.ContentType, content.ContentId)
}

func (session *Session) UpdateUserContent(content client.Content) {
	_, exist := session.UserContentDiffs[content.ContentType]
	if !exist {
		session.UserContentDiffs[content.ContentType] = make(map[int32]client.Content)
	}
	session.UserContentDiffs[content.ContentType][content.ContentId] = content
}

func (session *Session) AddContent(content client.Content) {
	handler, exist := contentHandler[content.ContentType]
	if !exist {
		fmt.Println("TODO: Add handler for content type ", content.ContentType)
		return
	}
	handler(session, content.ContentType, content.ContentId, content.ContentAmount)
}

func (session *Session) RemoveContent(content client.Content) {
	handler, exist := contentHandler[content.ContentType]
	if !exist {
		fmt.Println("TODO: Add handler for content type ", content.ContentType)
		return
	}
	handler(session, content.ContentType, content.ContentId, -content.ContentAmount)
}

func init() {
	// reference for type can be found in elichika/enum/content_type
	contentHandler[enum.ContentTypeSnsCoin] = userStatusContentHandler
	contentHandler[enum.ContentTypeCardExp] = userStatusContentHandler
	contentHandler[enum.ContentTypeGameMoney] = userStatusContentHandler
	contentHandler[enum.ContentTypeSubscriptionCoin] = userStatusContentHandler

	contentHandler[enum.ContentTypeSuit] = suitHandler

	contentHandler[enum.ContentTypeGachaTicket] = genericContentHandler
	userModelField[enum.ContentTypeGachaTicket] = "UserGachaTicketByTicketId"

	contentHandler[enum.ContentTypeGachaPoint] = genericContentHandler // gacha point (quartz)
	userModelField[enum.ContentTypeGachaPoint] = "UserGachaPointByPointId"

	contentHandler[enum.ContentTypeLessonEnhancingItem] = genericContentHandler // light bulbs
	userModelField[enum.ContentTypeLessonEnhancingItem] = "UserLessonEnhancingItemByItemId"

	contentHandler[enum.ContentTypeTrainingMaterial] = genericContentHandler // training items (macarons, memorials)
	userModelField[enum.ContentTypeTrainingMaterial] = "UserTrainingMaterialByItemId"

	contentHandler[enum.ContentTypeGradeUpper] = genericContentHandler // card grade up items
	userModelField[enum.ContentTypeGradeUpper] = "UserGradeUpItemByItemId"

	contentHandler[enum.ContentTypeRecoveryLp] = genericContentHandler // lp candies
	userModelField[enum.ContentTypeRecoveryLp] = "UserRecoveryLpById"

	contentHandler[enum.ContentTypeRecoveryAp] = genericContentHandler // training ticket
	userModelField[enum.ContentTypeRecoveryAp] = "UserRecoveryApById"

	contentHandler[enum.ContentTypeAccessoryLevelUp] = genericContentHandler // accessory stickers
	userModelField[enum.ContentTypeAccessoryLevelUp] = "UserAccessoryLevelUpItemById"

	contentHandler[enum.ContentTypeAccessoryRarityUp] = genericContentHandler // accessory rarity up items
	userModelField[enum.ContentTypeAccessoryRarityUp] = "UserAccessoryRarityUpItemById"

	// generics exchange point (SBL / DLP)
	// also include channel exchanges
	contentHandler[enum.ContentTypeExchangeEventPoint] = genericContentHandler
	userModelField[enum.ContentTypeExchangeEventPoint] = "UserExchangeEventPointById"

	contentHandler[enum.ContentTypeLiveSkipTicket] = genericContentHandler // skip tickets
	userModelField[enum.ContentTypeLiveSkipTicket] = "UserLiveSkipTicketById"

	contentHandler[enum.ContentTypeStoryEventUnlock] = genericContentHandler // event story unlock key
	userModelField[enum.ContentTypeStoryEventUnlock] = "UserStoryEventUnlockItemById"

	contentHandler[enum.ContentTypeEventMarathonBooster] = genericContentHandler // marathon boosters
	userModelField[enum.ContentTypeEventMarathonBooster] = "UserEventMarathonBoosterById"

	contentHandler[enum.ContentTypeRecoveryTowerCardUsedCount] = genericContentHandler // dlp water bottle
	userModelField[enum.ContentTypeRecoveryTowerCardUsedCount] = "UserRecoveryTowerCardUsedCountItemByRecoveryTowerCardUsedCountItemMasterId"

	contentHandler[enum.ContentTypeStoryMember] = memberStoryHandler
	contentHandler[enum.ContentTypeVoice] = voiceHandler
}

func genericContentHandler(session *Session, contentType, contentId, contentAmount int32) {
	content := session.GetUserContent(contentType, contentId)
	content.ContentAmount += contentAmount
	session.UpdateUserContent(content)
}

func suitHandler(session *Session, _, suitMasterId, _ int32) {
	session.InsertUserSuit(suitMasterId)
}

func memberStoryHandler(session *Session, _, memberStoryId, _ int32) {
	session.InsertMemberStory(memberStoryId)
}

func voiceHandler(session *Session, _, naviVoiceMasterId, _ int32) {
	session.UpdateVoice(naviVoiceMasterId, false)
}

// these content amount are stored in the user status
func userStatusContentHandler(session *Session, contentType, contentId, amount int32) {
	switch contentType {
	case enum.ContentTypeSnsCoin: // star gems
		session.AddSnsCoin(amount)
	case enum.ContentTypeCardExp: // card exp
		session.AddCardExp(amount)
	case enum.ContentTypeGameMoney: // game money (gold)
		session.AddGameMoney(amount)
	case enum.ContentTypeSubscriptionCoin: // subscription coin (purple coin)
		session.UserStatus.SubscriptionCoin += amount
	default:
		fmt.Println("TODO: handle user status content type:", contentType)
	}
}

func genericContentByContentIdFinalizer(session *Session) {
	rModel := reflect.ValueOf(&session.UserModel)
	for contentType, contentDiffByContentId := range session.UserContentDiffs {
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

		for _, content := range contentDiffByContentId {
			// update or insert the content
			affected, err := session.Db.Table("u_content").Where("user_id = ? AND content_type = ? AND content_id = ?",
				session.UserId, content.ContentType, content.ContentId).AllCols().Update(content)
			utils.CheckErr(err)
			if affected == 0 { // doesn't exist, insert
				genericDatabaseInsert(session, "u_content", content)
			}

			obj := reflect.New(rElementType)
			rElementFromContent.Func.Call([]reflect.Value{obj, reflect.ValueOf(content)})
			rDictionarySet.Func.Call([]reflect.Value{rDictionary.Addr(), reflect.ValueOf(content.ContentId), reflect.Indirect(obj)})
		}
	}
}

func genericContentByContentIdPopulator(session *Session) {
	rModel := reflect.ValueOf(&session.UserModel)
	contents := []client.Content{}
	// order by is not necessary
	err := session.Db.Table("u_content").Where("user_id = ?", session.UserId).Find(&contents)
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
		for _, content := range contents {
			obj := reflect.New(rElementType)
			rElementFromContent.Func.Call([]reflect.Value{obj, reflect.ValueOf(content)})
			rDictionarySet.Func.Call([]reflect.Value{rDictionary.Addr(), reflect.ValueOf(content.ContentId), reflect.Indirect(obj)})
		}
	}
}

// this is used for loading exported account data
func (session *Session) populateGenericContentDiffFromUserModel() {
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
			session.UpdateUserContent(content.(client.Content))
		}
	}

}

func init() {
	addFinalizer(genericContentByContentIdFinalizer)
	addPopulator(genericContentByContentIdPopulator)
}
