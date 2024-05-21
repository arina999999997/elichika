package serverdata

import (
	"elichika/client"
	"elichika/config"
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

	InitTable("s_gacha_guarantee", GachaGuarantee{}, overwrite)
	InitTable("s_gacha", ServerGacha{}, overwrite)
	InitTable("s_gacha_group", GachaGroup{}, overwrite)
	InitTable("s_gacha_card", GachaCard{}, overwrite)
	InitTable("s_trade", client.Trade{}, overwrite)
	InitTable("s_trade_product", client.TradeProduct{}, overwrite)
	InitTable("s_login_bonus", LoginBonus{}, overwrite)
	InitTable("s_login_bonus_reward_day", LoginBonusRewardDay{}, overwrite)
	InitTable("s_login_bonus_reward_content", LoginBonusRewardContent{}, overwrite)
	InitTable("s_ng_word", NgWord{}, overwrite)
}

func AutoInsert() {
	session := Engine.NewSession()
	defer session.Close()
	session.Begin()
	total, err := session.Table("s_trade").Count()
	utils.CheckErr(err)
	if total == 0 { // already have something
		TradeCli(session, []string{"insert", config.ServerInitJsons + "trade.json"})
	}

	total, err = session.Table("s_gacha").Count()
	utils.CheckErr(err)
	if total == 0 {
		GachaCli(session, []string{"init"})
		GachaCli(session, []string{"insert", config.ServerInitJsons + "gacha.json"})
	}
	total, err = session.Table("s_login_bonus").Count()
	utils.CheckErr(err)
	if total == 0 {
		InitialiseLoginBonus(session)
	}
	total, err = session.Table("s_ng_word").Count()
	utils.CheckErr(err)
	if total == 0 {
		InitialiseNgWord(session)
	}
	session.Commit()
}

func init() {
	var err error
	Engine, err = xorm.NewEngine("sqlite", config.ServerdataPath)
	utils.CheckErr(err)
	Engine.SetMaxOpenConns(50)
	Engine.SetMaxIdleConns(10)
	InitTables(false) // insert new tables if necessary
	AutoInsert()

}
