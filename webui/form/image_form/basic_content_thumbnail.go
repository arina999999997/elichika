package image_form

import (
	"elichika/enum"
	"elichika/gamedata"
	"elichika/log"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

// common thumbnail html, used for items that have frame ready
// for items like cards, there are a frame around it depending on rarity / attribute that we have to setup (for now this is just ignored)
const contentThumbnailImageHTMLFormatString = `<img class="content-thumbnail-image" src="%s" />`

var cachedContentThumbnailImageHTML = map[int32]map[int32]*string{}
var cachedExchangeEventPointIconAssetPathHTML = map[int32]*string{}
var genericContentTypeToTableName = map[int32]string{}
var contentTypeToUiTextureKey = map[int32]int32{}

// create the map
func init() {
	cachedContentThumbnailImageHTML[enum.ContentTypeSnsCoin] = map[int32]*string{}
	cachedContentThumbnailImageHTML[enum.ContentTypeCard] = map[int32]*string{}
	cachedContentThumbnailImageHTML[enum.ContentTypeCardExp] = map[int32]*string{}
	cachedContentThumbnailImageHTML[enum.ContentTypeGachaPoint] = map[int32]*string{}
	cachedContentThumbnailImageHTML[enum.ContentTypeLessonEnhancingItem] = map[int32]*string{}
	cachedContentThumbnailImageHTML[enum.ContentTypeSuit] = map[int32]*string{}
	cachedContentThumbnailImageHTML[enum.ContentTypeVoice] = map[int32]*string{}
	cachedContentThumbnailImageHTML[enum.ContentTypeGachaTicket] = map[int32]*string{}
	cachedContentThumbnailImageHTML[enum.ContentTypeGameMoney] = map[int32]*string{}
	cachedContentThumbnailImageHTML[enum.ContentTypeTrainingMaterial] = map[int32]*string{}
	cachedContentThumbnailImageHTML[enum.ContentTypeGradeUpper] = map[int32]*string{}
	cachedContentThumbnailImageHTML[enum.ContentTypeGiftBox] = map[int32]*string{}
	cachedContentThumbnailImageHTML[enum.ContentTypeEmblem] = map[int32]*string{}
	cachedContentThumbnailImageHTML[enum.ContentTypeRecoveryAp] = map[int32]*string{}
	cachedContentThumbnailImageHTML[enum.ContentTypeRecoveryLp] = map[int32]*string{}
	cachedContentThumbnailImageHTML[enum.ContentTypeStorySide] = map[int32]*string{}
	cachedContentThumbnailImageHTML[enum.ContentTypeStoryMember] = map[int32]*string{}
	cachedContentThumbnailImageHTML[enum.ContentTypeExchangeEventPoint] = map[int32]*string{}
	cachedContentThumbnailImageHTML[enum.ContentTypeAccessory] = map[int32]*string{}
	cachedContentThumbnailImageHTML[enum.ContentTypeAccessoryLevelUp] = map[int32]*string{}
	cachedContentThumbnailImageHTML[enum.ContentTypeAccessoryRarityUp] = map[int32]*string{}
	cachedContentThumbnailImageHTML[enum.ContentTypeCustomBackground] = map[int32]*string{}
	cachedContentThumbnailImageHTML[enum.ContentTypeEventMarathonBooster] = map[int32]*string{}
	cachedContentThumbnailImageHTML[enum.ContentTypeLiveSkipTicket] = map[int32]*string{}
	cachedContentThumbnailImageHTML[enum.ContentTypeEventMiningBooster] = map[int32]*string{}
	cachedContentThumbnailImageHTML[enum.ContentTypeStoryEventUnlock] = map[int32]*string{}
	cachedContentThumbnailImageHTML[enum.ContentTypeRecoveryTowerCardUsedCount] = map[int32]*string{}
	cachedContentThumbnailImageHTML[enum.ContentTypeSubscriptionCoin] = map[int32]*string{}
	cachedContentThumbnailImageHTML[enum.ContentTypeMemberGuildSupport] = map[int32]*string{}

	genericContentTypeToTableName[enum.ContentTypeGachaPoint] = "m_gacha_point"
	genericContentTypeToTableName[enum.ContentTypeLessonEnhancingItem] = "m_lesson_enhancing_item"
	genericContentTypeToTableName[enum.ContentTypeGachaTicket] = "m_gacha_ticket"
	genericContentTypeToTableName[enum.ContentTypeTrainingMaterial] = "m_training_material"
	genericContentTypeToTableName[enum.ContentTypeGradeUpper] = "m_grade_upper"
	genericContentTypeToTableName[enum.ContentTypeRecoveryAp] = "m_recovery_ap"
	genericContentTypeToTableName[enum.ContentTypeRecoveryLp] = "m_recovery_lp"
	genericContentTypeToTableName[enum.ContentTypeExchangeEventPoint] = "m_exchange_event_point"
	genericContentTypeToTableName[enum.ContentTypeAccessory] = "m_accessory"
	genericContentTypeToTableName[enum.ContentTypeAccessoryLevelUp] = "m_accessory_level_up_item"
	genericContentTypeToTableName[enum.ContentTypeAccessoryRarityUp] = "m_accessory_rarity_up_item"
	genericContentTypeToTableName[enum.ContentTypeCustomBackground] = "m_custom_background"
	genericContentTypeToTableName[enum.ContentTypeEventMarathonBooster] = "m_event_marathon_booster_item"
	genericContentTypeToTableName[enum.ContentTypeLiveSkipTicket] = "m_live_skip_ticket"
	genericContentTypeToTableName[enum.ContentTypeEventMiningBooster] = "m_event_mining_booster_item"
	genericContentTypeToTableName[enum.ContentTypeStoryEventUnlock] = "m_story_event_unlock_item"
	genericContentTypeToTableName[enum.ContentTypeRecoveryTowerCardUsedCount] = "m_recovery_tower_card_used_count_item"
	genericContentTypeToTableName[enum.ContentTypeMemberGuildSupport] = "m_member_guild_support_item"

	contentTypeToUiTextureKey[enum.ContentTypeSnsCoin] = enum.UiTextureKeySnsCoinIcon
	contentTypeToUiTextureKey[enum.ContentTypeCardExp] = enum.UiTextureKeyCardExpIcon
	contentTypeToUiTextureKey[enum.ContentTypeGameMoney] = enum.UiTextureKeyGameMoneyIcon
	contentTypeToUiTextureKey[enum.ContentTypeVoice] = enum.UiTextureKeyVoiceIcon
	contentTypeToUiTextureKey[enum.ContentTypeStorySide] = enum.UiTextureKeyStoryMember
	contentTypeToUiTextureKey[enum.ContentTypeStoryMember] = enum.UiTextureKeyStoryMember
	contentTypeToUiTextureKey[enum.ContentTypeSubscriptionCoin] = enum.UiTextureKeySubscriptionCoinIcon
}

func GetContentThumbnailImageHTML(contentType, contentId int32) *string {
	ptr, exists := cachedContentThumbnailImageHTML[contentType][contentId]
	if exists {
		return ptr
	}
	db := gamedata.Instance.MasterdataDb
	var assetPath string
	// each content type has a table that we can fetch this data from
	switch contentType {
	// generic cases: these each have a table that has the following structure:
	// - id: column using contentId
	// - thumbnail_asset_path: the asset path directly
	case enum.ContentTypeGachaPoint: // m_gacha_point
		fallthrough
	case enum.ContentTypeLessonEnhancingItem: // m_lesson_enhancing_item
		fallthrough
	case enum.ContentTypeGachaTicket: // m_gacha_ticket
		fallthrough
	case enum.ContentTypeTrainingMaterial: // m_training_material
		fallthrough
	case enum.ContentTypeGradeUpper: // m_grade_upper
		fallthrough
	case enum.ContentTypeRecoveryAp: // m_recovery_ap
		fallthrough
	case enum.ContentTypeRecoveryLp: // m_recovery_lp
		fallthrough
	case enum.ContentTypeExchangeEventPoint: // m_exchange_event_point
		fallthrough
	case enum.ContentTypeAccessory: // m_accessory
		fallthrough
	case enum.ContentTypeAccessoryLevelUp: // m_accessory_level_up_item
		fallthrough
	case enum.ContentTypeAccessoryRarityUp: // m_accessory_rarity_up_item
		fallthrough
	case enum.ContentTypeCustomBackground: // m_custom_background
		fallthrough
	case enum.ContentTypeEventMarathonBooster: // m_event_marathon_booster_item
		fallthrough
	case enum.ContentTypeLiveSkipTicket: // m_live_skip_ticket
		fallthrough
	case enum.ContentTypeEventMiningBooster: // m_event_mining_booster_item
		fallthrough
	case enum.ContentTypeStoryEventUnlock: // m_story_event_unlock_item
		fallthrough
	case enum.ContentTypeRecoveryTowerCardUsedCount: // m_recovery_tower_card_used_count_item
		fallthrough
	case enum.ContentTypeMemberGuildSupport: // m_member_guild_support_item
		tableName := genericContentTypeToTableName[contentType]
		db.Do(func(session *xorm.Session) {
			exists, err := session.Table(tableName).Where("id = ?", contentId).Cols("thumbnail_asset_path").Get(&assetPath)
			utils.CheckErrMustExist(err, exists)
		})
	// ui texture keys:
	// - these are the texture that can be found in m_ui_texture
	// - the key is the enum UiTextureKey
	case enum.ContentTypeSnsCoin:
		fallthrough
	case enum.ContentTypeCardExp:
		fallthrough
	case enum.ContentTypeGameMoney:
		fallthrough
	case enum.ContentTypeVoice:
		fallthrough
	case enum.ContentTypeStorySide:
		fallthrough
	case enum.ContentTypeStoryMember:
		fallthrough
	case enum.ContentTypeSubscriptionCoin:
		textureKey := contentTypeToUiTextureKey[contentType]
		db.Do(func(session *xorm.Session) {
			exists, err := session.Table("m_ui_texture").Where("id = ?", textureKey).Cols("asset_path").Get(&assetPath)
			utils.CheckErrMustExist(err, exists)
		})
	// finally special case
	case enum.ContentTypeCard:
		// TODO(webui visual): card need proper frame
		assetPath, _ = GetCardAssetPaths(contentId)
	case enum.ContentTypeSuit:
		db.Do(func(session *xorm.Session) {
			exists, err := session.Table("m_suit").Where("id = ?", contentId).Cols("thumbnail_image_asset_path").Get(&assetPath)
			utils.CheckErrMustExist(err, exists)
		})
	case enum.ContentTypeEmblem:
		// TODO(webui visual): emblem is a bit special, we need to layer on top
		db.Do(func(session *xorm.Session) {
			exists, err := session.Table("m_emblem").Where("id = ?", contentId).Cols("emblem_asset_path").Get(&assetPath)
			utils.CheckErrMustExist(err, exists)
		})
	// and failure cases
	case enum.ContentTypeGiftBox:
		// this is unused
		log.Panic("ContentTypeGiftBox is unused")
	default:
		log.Panic("Unknown content type")
	}
	cachedContentThumbnailImageHTML[contentType][contentId] = new(string)
	*cachedContentThumbnailImageHTML[contentType][contentId] = fmt.Sprintf(contentThumbnailImageHTMLFormatString, GetImageByAssetPath(assetPath))
	return cachedContentThumbnailImageHTML[contentType][contentId]
}

func GetExchangeEventPointIconAssetPath(id int32) *string {
	ptr, exists := cachedExchangeEventPointIconAssetPathHTML[id]
	if exists {
		return ptr
	}
	db := gamedata.Instance.MasterdataDb
	assetPath := ""
	db.Do(func(session *xorm.Session) {
		exists, err := session.Table("m_exchange_event_point").Where("id = ?", id).Cols("icon_asset_path").Get(&assetPath)
		utils.CheckErrMustExist(err, exists)
	})
	cachedExchangeEventPointIconAssetPathHTML[id] = new(string)
	*cachedExchangeEventPointIconAssetPathHTML[id] = fmt.Sprintf(contentThumbnailImageHTMLFormatString, GetImageByAssetPath(assetPath))
	return cachedExchangeEventPointIconAssetPathHTML[id]
}
