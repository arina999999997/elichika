package gamedata

import (
	"elichika/client"
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type LoginBonus struct {
	LoginBonusId            int                        `xorm:"pk 'login_bonus_id'"`
	LoginBonusType          int                        `xorm:"'login_bonus_type'"`
	StartAt                 int64                      `xorm:"'start_at'"`
	EndAt                   int64                      `xorm:"'end_at'"`
	BackgroundId            int                        `xorm:"'background_id'"`
	WhiteboardTextureAsset  client.TextureStruktur     `xorm:"'whiteboard_texture_asset'"`
	LoginBonusHandler       string                     `xorm:"'login_bonus_handler'"`
	LoginBonusHandlerConfig string                     `xorm:"'login_bonus_handler_config'"`
	LoginBonusRewards       []client.LoginBonusRewards `xorm:"-"`
}

func (lb *LoginBonus) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	rewardDays := []client.LoginBonusRewardDay{}
	err := serverdata_db.Table("s_login_bonus_reward_day").Where("login_bonus_id = ?", lb.LoginBonusId).
		OrderBy("day").Find(&rewardDays)
	utils.CheckErr(err)
	for _, day := range rewardDays {
		reward := client.LoginBonusRewards{
			Day:          day.Day,
			ContentGrade: day.ContentGrade,
		}
		err = serverdata_db.Table("s_login_bonus_reward_content").
			Where("login_bonus_id = ? AND day = ?", lb.LoginBonusId, day.Day).Find(&reward.LoginBonusContents)
		utils.CheckErr(err)
		lb.LoginBonusRewards = append(lb.LoginBonusRewards, reward)
	}
}

func loadLoginBonus(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading LoginBonus")
	gamedata.LoginBonus = make(map[int]*LoginBonus)
	err := serverdata_db.Table("s_login_bonus").Find(&gamedata.LoginBonus)
	utils.CheckErr(err)
	for _, loginBonus := range gamedata.LoginBonus {
		loginBonus.populate(gamedata, masterdata_db, serverdata_db, dictionary)
	}
}

func init() {
	addLoadFunc(loadLoginBonus)
}

func (lb *LoginBonus) NaviLoginBonus() client.NaviLoginBonus {
	return client.NaviLoginBonus{
		LoginBonusId:           lb.LoginBonusId,
		LoginBonusRewards:      append([]client.LoginBonusRewards{}, lb.LoginBonusRewards...), // copy the list
		BackgroundId:           lb.BackgroundId,
		WhiteboardTextureAsset: &lb.WhiteboardTextureAsset,
		StartAt:                lb.StartAt,
		EndAt:                  lb.EndAt,
	}
}
