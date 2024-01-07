package serverdata

import (
	"elichika/enum"
	"elichika/item"
	// "elichika/model"
	"elichika/client"
	"elichika/utils"

	"xorm.io/xorm"
)

func InitialiseLoginBonus(session *xorm.Session) {
	const BeginnerLoginBonusId = 1000001
	const NormalLoginBonusId = 1000002
	const BirthDayLoginBonusId = 1000003
	loginBonuses := []client.LoginBonus{}
	loginBonusRewardDays := []client.LoginBonusRewardDay{}
	loginBonusRewardContents := []client.LoginBonusRewardContent{}

	// beginner login bonus
	loginBonuses = append(loginBonuses, client.LoginBonus{
		LoginBonusId:   BeginnerLoginBonusId,
		LoginBonusType: enum.LoginBonusTypeNormal,
		BackgroundId:   100100700,
		WhiteboardTextureAsset: &client.TextureStruktur{
			V: ":7S",
		},
		StartAt:                 0,
		EndAt:                   1<<31 - 1,
		LoginBonusHandler:       "beginner_login_bonus",
		LoginBonusHandlerConfig: "",
	})
	for day := 1; day <= 7; day++ {
		loginBonusRewardDays = append(loginBonusRewardDays, client.LoginBonusRewardDay{
			LoginBonusId: BeginnerLoginBonusId,
			Day:          day,
			ContentGrade: enum.LoginBonusContentGradeNormal,
		})
	}
	loginBonusRewardContents = append(loginBonusRewardContents, client.LoginBonusRewardContent{
		LoginBonusId: BeginnerLoginBonusId,
		Day:          1,
		Content:      item.Gold.Amount(300000),
	})
	loginBonusRewardContents = append(loginBonusRewardContents, client.LoginBonusRewardContent{
		LoginBonusId: BeginnerLoginBonusId,
		Day:          2,
		Content:      item.SRScoutingTicket,
	})
	loginBonusRewardContents = append(loginBonusRewardContents, client.LoginBonusRewardContent{
		LoginBonusId: BeginnerLoginBonusId,
		Day:          3,
		Content:      item.EXP.Amount(300000),
	})
	loginBonusRewardContents = append(loginBonusRewardContents, client.LoginBonusRewardContent{
		LoginBonusId: BeginnerLoginBonusId,
		Day:          4,
		Content:      item.ShowCandy100,
	})
	loginBonusRewardContents = append(loginBonusRewardContents, client.LoginBonusRewardContent{
		LoginBonusId: BeginnerLoginBonusId,
		Day:          5,
		Content:      item.StarGem.Amount(50),
	})
	loginBonusRewardContents = append(loginBonusRewardContents, client.LoginBonusRewardContent{
		LoginBonusId: BeginnerLoginBonusId,
		Day:          6,
		Content:      item.MemoryKey.Amount(7),
	})
	loginBonusRewardContents = append(loginBonusRewardContents, client.LoginBonusRewardContent{
		LoginBonusId: BeginnerLoginBonusId,
		Day:          7,
		Content:      item.URScoutingTicket,
	})

	// normal login bonus
	loginBonuses = append(loginBonuses, client.LoginBonus{
		LoginBonusId:   NormalLoginBonusId,
		LoginBonusType: enum.LoginBonusTypeNormal,
		BackgroundId:   100100700,
		WhiteboardTextureAsset: &client.TextureStruktur{
			V: "/4n",
		},
		StartAt:                 0,
		EndAt:                   1<<31 - 1,
		LoginBonusHandler:       "normal_login_bonus",
		LoginBonusHandlerConfig: "",
	})
	for day := 1; day <= 7; day++ {
		loginBonusRewardDays = append(loginBonusRewardDays, client.LoginBonusRewardDay{
			LoginBonusId: NormalLoginBonusId,
			Day:          day,
			ContentGrade: enum.LoginBonusContentGradeNormal,
		})
	}
	loginBonusRewardContents = append(loginBonusRewardContents, client.LoginBonusRewardContent{
		LoginBonusId: NormalLoginBonusId,
		Day:          1,
		Content:      item.ShowCandy50,
	})
	loginBonusRewardContents = append(loginBonusRewardContents, client.LoginBonusRewardContent{
		LoginBonusId: NormalLoginBonusId,
		Day:          2,
		Content:      item.StarGem.Amount(10),
	})
	loginBonusRewardContents = append(loginBonusRewardContents, client.LoginBonusRewardContent{
		LoginBonusId: NormalLoginBonusId,
		Day:          3,
		Content:      item.TrainingTicket,
	})
	loginBonusRewardContents = append(loginBonusRewardContents, client.LoginBonusRewardContent{
		LoginBonusId: NormalLoginBonusId,
		Day:          4,
		Content:      item.StarGem.Amount(20),
	})
	loginBonusRewardContents = append(loginBonusRewardContents, client.LoginBonusRewardContent{
		LoginBonusId: NormalLoginBonusId,
		Day:          5,
		Content:      item.ShowCandy50,
	})
	loginBonusRewardContents = append(loginBonusRewardContents, client.LoginBonusRewardContent{
		LoginBonusId: NormalLoginBonusId,
		Day:          6,
		Content:      item.MemoryKey,
	})
	loginBonusRewardContents = append(loginBonusRewardContents, client.LoginBonusRewardContent{
		LoginBonusId: NormalLoginBonusId,
		Day:          7,
		Content:      item.StarGem.Amount(30),
	})

	// birthday login bonus
	loginBonuses = append(loginBonuses, client.LoginBonus{
		LoginBonusId:            BirthDayLoginBonusId,
		LoginBonusType:          enum.LoginBonusTypeBirthday,
		StartAt:                 0,
		EndAt:                   1<<31 - 1,
		LoginBonusHandler:       "birthday_login_bonus",
		LoginBonusHandlerConfig: "random", // can be set to latest to use the latest pair
	})

	for _, loginBonus := range loginBonuses {
		_, err := session.Table("s_login_bonus").Insert(loginBonus)
		utils.CheckErr(err)
	}

	for _, day := range loginBonusRewardDays {
		_, err := session.Table("s_login_bonus_reward_day").Insert(day)
		utils.CheckErr(err)
	}

	for _, content := range loginBonusRewardContents {
		_, err := session.Table("s_login_bonus_reward_content").Insert(content)
		utils.CheckErr(err)
	}

}
