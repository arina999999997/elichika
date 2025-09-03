package serverdata

import (
	"elichika/config"
	"elichika/log"
	"elichika/utils"

	"encoding/json"

	"xorm.io/xorm"
)

// TODO(extra): the topic rewards are all cards for existing events
// but depending on the client, they could be other things too
// main name subname stuff depends on the name of the member of the card, so no need to store them all
type EventTopicReward struct {
	EventId          int32 `xorm:"pk 'event_id'"`
	DisplayOrder     int32 `xorm:"pk 'display_order'"`
	RewardCardId     int32 `xorm:"reward_card_id"`
	RewardCardAmount int32 `xorm:"reward_card_amount"`
}

type EventMemberNameAsset struct {
	MemberMasterId          int32  `json:"member_master_id" xorm:"pk 'member_master_id'"`
	MainNameTopAssetPath    string `json:"main_name_top_asset_path" xorm:"'main_name_top_asset_path'"`
	MainNameBottomAssetPath string `json:"main_name_bottom_asset_path" xorm:"'main_name_bottom_asset_path'"`
	SubNameTopAssetPath     string `json:"sub_name_top_asset_path" xorm:"'sub_name_top_asset_path'"`
	SubNameBottomAssetPath  string `json:"sub_name_bottom_asset_path" xorm:"'sub_name_bottom_asset_path'"`
}

func initEventMemberNameAsset(session *xorm.Session) {
	path := config.AssetPath + "event/event_member_name.json"

	log.Printf("Parsing event member name asset file: %s\n", path)
	text := utils.ReadAllText(path)

	assets := []EventMemberNameAsset{}
	err := json.Unmarshal([]byte(text), &assets)
	utils.CheckErr(err)
	_, err = session.Table("s_event_member_name_asset").Insert(assets)
	utils.CheckErr(err)
}

func init() {
	addTable("s_event_member_name_asset", EventMemberNameAsset{}, initEventMemberNameAsset)
}
