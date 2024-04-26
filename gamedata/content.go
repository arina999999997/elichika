package gamedata

import (
	"elichika/client"
	"elichika/dictionary"
	"elichika/enum"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type Content struct {
	Content client.Content
	Name    string
}

func loadContent(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading Content")
	gamedata.Content = make(map[int32]map[int32]*Content)
	for contentType := range gamedata.ContentType {
		gamedata.Content[contentType] = make(map[int32]*Content)
	}
	genericLoading := func(contentType int32, table string) {
		type genericContent struct {
			Id   int32  `xorm:"pk 'id'"`
			Name string `xorm:"'name'"`
		}
		contents := []genericContent{}
		err := masterdata_db.Table(table).Find(&contents)
		utils.CheckErr(err)
		for _, content := range contents {
			gamedata.Content[contentType][content.Id] = &Content{
				Content: client.Content{
					ContentType: contentType,
					ContentId:   content.Id,
				},
				Name: dictionary.Resolve(content.Name),
			}
		}
	}
	genericLoading(enum.ContentTypeGachaPoint, "m_gacha_point_setting")
	genericLoading(enum.ContentTypeLessonEnhancingItem, "m_lesson_enhancing_item")
	genericLoading(enum.ContentTypeGachaTicket, "m_gacha_ticket")
	genericLoading(enum.ContentTypeTrainingMaterial, "m_training_material")
	genericLoading(enum.ContentTypeGradeUpper, "m_grade_upper")
	genericLoading(enum.ContentTypeRecoveryAp, "m_recovery_ap")
	genericLoading(enum.ContentTypeRecoveryLp, "m_recovery_lp")
	genericLoading(enum.ContentTypeExchangeEventPoint, "m_exchange_event_point")
	genericLoading(enum.ContentTypeAccessoryLevelUp, "m_accessory_level_up_item")
	genericLoading(enum.ContentTypeAccessoryRarityUp, "m_accessory_rarity_up_item")
	genericLoading(enum.ContentTypeEventMarathonBooster, "m_event_marathon_booster_item")
	genericLoading(enum.ContentTypeLiveSkipTicket, "m_live_skip_ticket")
	genericLoading(enum.ContentTypeStoryEventUnlock, "m_story_event_unlock_item")
	genericLoading(enum.ContentTypeRecoveryTowerCardUsedCount, "m_recovery_tower_card_used_count_item")

	gamedata.ContentsByContentType = make(map[int32][]*client.Content)
	for contentType := range gamedata.Content {
		for _, content := range gamedata.Content[contentType] {
			gamedata.ContentsByContentType[contentType] = append(gamedata.ContentsByContentType[contentType], &content.Content)
		}
	}
}

func init() {
	addLoadFunc(loadContent)
	addPrequisite(loadContent, loadContentType)
}
