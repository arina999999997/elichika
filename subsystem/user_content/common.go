package user_content

import (
	"elichika/client"
	"elichika/enum"
	"elichika/userdata"
	"elichika/utils"

	"fmt"
	"reflect"
)

var (
	contentHandlerByContentType = map[int32]func(*userdata.Session, client.Content) (bool, any){}

	userModelField = map[int32]string{}
)

func AddContentHandler(contentType int32, handler func(*userdata.Session, client.Content) (bool, any)) {
	_, exist := contentHandlerByContentType[contentType]
	if exist {
		panic("handler for content type already existed")
	}
	contentHandlerByContentType[contentType] = handler
}

func init() {
	// reference for type can be found in elichika/enum/content_type

	contentHandlerByContentType[enum.ContentTypeGachaTicket] = genericContentHandler
	userModelField[enum.ContentTypeGachaTicket] = "UserGachaTicketByTicketId"

	contentHandlerByContentType[enum.ContentTypeGachaPoint] = genericContentHandler // gacha point (quartz)
	userModelField[enum.ContentTypeGachaPoint] = "UserGachaPointByPointId"

	contentHandlerByContentType[enum.ContentTypeLessonEnhancingItem] = genericContentHandler // light bulbs
	userModelField[enum.ContentTypeLessonEnhancingItem] = "UserLessonEnhancingItemByItemId"

	contentHandlerByContentType[enum.ContentTypeTrainingMaterial] = genericContentHandler // training items (macarons, memorials)
	userModelField[enum.ContentTypeTrainingMaterial] = "UserTrainingMaterialByItemId"

	contentHandlerByContentType[enum.ContentTypeGradeUpper] = genericContentHandler // card grade up items
	userModelField[enum.ContentTypeGradeUpper] = "UserGradeUpItemByItemId"

	contentHandlerByContentType[enum.ContentTypeRecoveryLp] = genericContentHandler // lp candies
	userModelField[enum.ContentTypeRecoveryLp] = "UserRecoveryLpById"

	contentHandlerByContentType[enum.ContentTypeRecoveryAp] = genericContentHandler // training ticket
	userModelField[enum.ContentTypeRecoveryAp] = "UserRecoveryApById"

	contentHandlerByContentType[enum.ContentTypeAccessoryLevelUp] = genericContentHandler // accessory stickers
	userModelField[enum.ContentTypeAccessoryLevelUp] = "UserAccessoryLevelUpItemById"

	contentHandlerByContentType[enum.ContentTypeAccessoryRarityUp] = genericContentHandler // accessory rarity up items
	userModelField[enum.ContentTypeAccessoryRarityUp] = "UserAccessoryRarityUpItemById"

	// generics exchange point (SBL / DLP)
	// also include channel exchanges
	contentHandlerByContentType[enum.ContentTypeExchangeEventPoint] = genericContentHandler
	userModelField[enum.ContentTypeExchangeEventPoint] = "UserExchangeEventPointById"

	contentHandlerByContentType[enum.ContentTypeLiveSkipTicket] = genericContentHandler // skip tickets
	userModelField[enum.ContentTypeLiveSkipTicket] = "UserLiveSkipTicketById"

	contentHandlerByContentType[enum.ContentTypeStoryEventUnlock] = genericContentHandler // event story unlock key
	userModelField[enum.ContentTypeStoryEventUnlock] = "UserStoryEventUnlockItemById"

	contentHandlerByContentType[enum.ContentTypeEventMarathonBooster] = genericContentHandler // marathon boosters
	userModelField[enum.ContentTypeEventMarathonBooster] = "UserEventMarathonBoosterById"

	contentHandlerByContentType[enum.ContentTypeRecoveryTowerCardUsedCount] = genericContentHandler // dlp water bottle
	userModelField[enum.ContentTypeRecoveryTowerCardUsedCount] = "UserRecoveryTowerCardUsedCountItemByRecoveryTowerCardUsedCountItemMasterId"
}

func genericContentByContentIdFinalizer(session *userdata.Session) {
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
				userdata.GenericDatabaseInsert(session, "u_content", content)
			}

			obj := reflect.New(rElementType)
			rElementFromContent.Func.Call([]reflect.Value{obj, reflect.ValueOf(content)})
			rDictionarySet.Func.Call([]reflect.Value{rDictionary.Addr(), reflect.ValueOf(content.ContentId), reflect.Indirect(obj)})
		}
	}
}

func genericContentByContentIdPopulator(session *userdata.Session) {
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

func init() {
	userdata.AddContentFinalizer(genericContentByContentIdFinalizer)
	userdata.AddContentPopulator(genericContentByContentIdPopulator)
}
