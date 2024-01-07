package serverdata

import (
	"elichika/client"
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
	InitTable("s_gacha_appeal", model.GachaAppeal{}, overwrite)
	InitTable("s_gacha_draw", model.GachaDraw{}, overwrite)
	InitTable("s_gacha", model.Gacha{}, overwrite)
	InitTable("s_gacha_group", model.GachaGroup{}, overwrite)
	InitTable("s_gacha_card", model.GachaCard{}, overwrite)
	InitTable("s_gacha_guarantee", model.GachaGuarantee{}, overwrite)
	InitTable("s_trade", model.Trade{}, overwrite)
	InitTable("s_trade_product", model.TradeProduct{}, overwrite)
	InitTable("s_login_bonus", client.LoginBonus{}, overwrite)
	InitTable("s_login_bonus_reward_day", client.LoginBonusRewardDay{}, overwrite)
	InitTable("s_login_bonus_reward_content", client.LoginBonusRewardContent{}, overwrite)

}

func AutoInsert() {
	session := Engine.NewSession()
	defer session.Close()
	session.Begin()
	total, err := session.Table("s_trade").Count()
	utils.CheckErr(err)
	if total > 0 { // already have something
		return
	}
	TradeCli(session, []string{"insert", config.ServerInitJsons + "trade.json"})
	GachaCli(session, []string{"init"})
	GachaCli(session, []string{"insert", config.ServerInitJsons + "gacha.json"})
	InitialiseLoginBonus(session)
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
