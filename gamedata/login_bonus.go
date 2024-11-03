package gamedata

import (
	"elichika/client"
	"elichika/generic"
	"elichika/serverdata"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type LoginBonus struct {
	LoginBonusId            int32                                  `xorm:"pk 'login_bonus_id'"`
	LoginBonusType          int32                                  `xorm:"'login_bonus_type'" enum:"LoginBonusType"`
	StartAt                 int64                                  `xorm:"'start_at'"`
	EndAt                   int64                                  `xorm:"'end_at'"`
	BackgroundId            int32                                  `xorm:"'background_id'"`
	WhiteboardTextureAsset  client.TextureStruktur                 `xorm:"'whiteboard_texture_asset'"`
	LoginBonusHandler       string                                 `xorm:"'login_bonus_handler'"`
	LoginBonusHandlerConfig string                                 `xorm:"'login_bonus_handler_config'"`
	LoginBonusRewards       generic.List[client.LoginBonusRewards] `xorm:"-"`
}

func (lb *LoginBonus) populate(gamedata *Gamedata) {
	rewardDays := []serverdata.LoginBonusRewardDay{}
	var err error
	gamedata.ServerdataDb.Do(func(session *xorm.Session) {
		err = session.Table("s_login_bonus_reward_day").Where("login_bonus_id = ?", lb.LoginBonusId).OrderBy("day").Find(&rewardDays)
	})
	utils.CheckErr(err)
	for _, day := range rewardDays {
		reward := client.LoginBonusRewards{
			Day:          day.Day,
			ContentGrade: generic.NewNullable(day.ContentGrade),
		}
		gamedata.ServerdataDb.Do(func(session *xorm.Session) {
			err = session.Table("s_login_bonus_reward_content").Where("login_bonus_id = ? AND day = ?", lb.LoginBonusId, day.Day).Find(&reward.LoginBonusContents.Slice)
		})
		utils.CheckErr(err)
		lb.LoginBonusRewards.Append(reward)
	}
}

func loadLoginBonus(gamedata *Gamedata) {
	fmt.Println("Loading LoginBonus")
	gamedata.LoginBonus = make(map[int32]*LoginBonus)
	var err error
	gamedata.ServerdataDb.Do(func(session *xorm.Session) {
		err = session.Table("s_login_bonus").Find(&gamedata.LoginBonus)
	})
	utils.CheckErr(err)
	for _, loginBonus := range gamedata.LoginBonus {
		loginBonus.populate(gamedata)
	}
}

func init() {
	addLoadFunc(loadLoginBonus)
}

func (lb *LoginBonus) NaviLoginBonus() client.NaviLoginBonus {
	return client.NaviLoginBonus{
		LoginBonusId:           lb.LoginBonusId,
		LoginBonusRewards:      lb.LoginBonusRewards.Copy(),
		BackgroundId:           lb.BackgroundId,
		WhiteboardTextureAsset: &lb.WhiteboardTextureAsset,
		StartAt:                lb.StartAt,
		EndAt:                  lb.EndAt,
	}
}
