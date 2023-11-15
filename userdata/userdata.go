package userdata

import (
	"elichika/config"
	"elichika/model"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

var (
	Engine *xorm.Engine
)

func InitTable(tableName string, structure interface{}, overwrite bool) {
	exist, err := Engine.Table(tableName).IsTableExist(tableName)
	utils.CheckErr(err)

	if !exist {
		fmt.Println("Creating new table:", tableName)
		err = Engine.Table(tableName).CreateTable(structure)
		utils.CheckErr(err)
	} else {
		if !overwrite {
			return
		}
		fmt.Println("Overwrite existing table:", tableName)
		err := Engine.DropTables(tableName)
		utils.CheckErr(err)
		err = Engine.Table(tableName).CreateTable(structure)
		utils.CheckErr(err)
	}
}

func InitTables(overwrite bool) {
	type DbUser struct {
		model.UserStatus           `xorm:"extends"`
		model.UserProfileLiveStats `xorm:"extends"`
	}

	InitTable("u_info", DbUser{}, overwrite)
	InitTable("u_custom_set_profile", model.UserSetProfile{}, overwrite)
	InitTable("u_card", model.UserCard{}, overwrite)
	InitTable("u_suit", model.UserSuit{}, overwrite)
	InitTable("u_accessory", model.UserAccessory{}, overwrite)
	InitTable("u_training_tree_cell", model.TrainingTreeCell{}, overwrite)

	type DbMember struct {
		model.UserMember          `xorm:"extends"`
		LovePanelLevel            int   `xorm:"'love_panel_level' default 1"`
		LovePanelLastLevelCellIds []int `xorm:"'love_panel_last_level_cell_ids' default '[]'"`
	}
	InitTable("u_member", DbMember{}, overwrite)
	InitTable("u_lesson_deck", model.UserLessonDeck{}, overwrite)
	InitTable("u_live_deck", model.UserLiveDeck{}, overwrite)
	InitTable("u_live_party", model.UserLiveParty{}, overwrite)
	InitTable("u_live_state", model.LiveState{}, overwrite)
	InitTable("u_play_list", model.UserPlayListItem{}, overwrite)
	type DbLiveRecord struct {
		model.UserLiveDifficulty `xorm:"extends"`
		Voltage                  int   `xorm:"'last_clear_voltage'" json:"voltage"`
		IsCleared                bool  `xorm:"'last_clear_is_cleared'" json:"is_cleared"`
		RecordedAt               int64 `xorm:"'last_clear_recorded_at'" json:"recorded_at"`
		CardWithSuitDict         []int `xorm:"'last_clear_cards_and_suits'" json:"card_with_suit_dict"`
		SquadDict                []any `xorm:"'squad_dict'" json:"squad_dict"`
	}
	InitTable("u_live_record", DbLiveRecord{}, overwrite)
	InitTable("u_trigger_basic", model.TriggerBasic{}, overwrite)
	InitTable("u_trigger_card_grade_up", model.TriggerCardGradeUp{}, overwrite)
	InitTable("u_trigger_member_love_level_up", model.TriggerMemberLoveLevelUp{}, overwrite)
	InitTable("u_resource", UserResource{}, overwrite)

	InitTable("u_trade_product", model.TradeProductUser{}, overwrite)
}

func init() {
	var err error
	Engine, err = xorm.NewEngine("sqlite", config.UserdataDbPath)
	utils.CheckErr(err)
	Engine.SetMaxOpenConns(50)
	Engine.SetMaxIdleConns(10)
	InitTables(false) // insert new tables if necessary
}
