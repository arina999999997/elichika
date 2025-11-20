package image_form

import (
	"elichika/locale"
	"elichika/utils"
)

var (
	CardAssetPathByCardMasterId       = map[int32]string{}
	CardAwakenAssetPathByCardMasterId = map[int32]string{}
)

func init() {
	type CardAppearance struct {
		CardMId            int32  `xorm:"'card_m_id'"`
		AppearanceType     int32  `xorm:"'appearance_type'"`
		ThumbnailAssetPath string `xorm:"'thumbnail_asset_path'"`
	}
	appearances := []CardAppearance{}
	path := locale.Locales["en"].Path
	err := locale.GetEngine(path + "masterdata.db").Table("m_card_appearance").Find(&appearances)
	utils.CheckErr(err)
	for _, ca := range appearances {
		if ca.AppearanceType == 1 {
			CardAssetPathByCardMasterId[ca.CardMId] = ca.ThumbnailAssetPath
		} else {
			CardAwakenAssetPathByCardMasterId[ca.CardMId] = ca.ThumbnailAssetPath
		}
	}
}

func GetCardAssetPaths(cardMasterId int32) (string, string) {
	return CardAssetPathByCardMasterId[cardMasterId], CardAwakenAssetPathByCardMasterId[cardMasterId]
}
