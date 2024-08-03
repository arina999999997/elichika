package serverdata

import (
	"elichika/client"
	"elichika/enum"
	"elichika/generic"
	"elichika/item"
	"elichika/utils"

	"xorm.io/xorm"
)

type LoginBonus struct {
	LoginBonusId            int32 `xorm:"pk"`
	LoginBonusType          int32 `xorm:"pk"`
	StartAt                 int64
	EndAt                   int64
	BackgroundId            int32
	WhiteboardTextureAsset  *client.TextureStruktur `xorm:"varchar(3)"`
	LoginBonusHandler       string
	LoginBonusHandlerConfig string
}

type LoginBonusRewardDay struct {
	LoginBonusId int32 `xorm:"pk"`
	Day          int32 `xorm:"pk"`
	ContentGrade int32 `enum:"LoginBonusContentGrade"`
}

type LoginBonusRewardContent struct {
	LoginBonusId int32
	Day          int32
	Content      client.Content `xorm:"extends"`
}

func InitializeLoginBonus(session *xorm.Session) {
	const BeginnerLoginBonusId = 1000001
	const NormalLoginBonusId = 1000002
	const BirthDayLoginBonusId = 1000003
	loginBonuses := []LoginBonus{}
	loginBonusRewardDays := []LoginBonusRewardDay{}
	loginBonusRewardContents := []LoginBonusRewardContent{}

	// beginner login bonus
	loginBonuses = append(loginBonuses, LoginBonus{
		LoginBonusId:   BeginnerLoginBonusId,
		LoginBonusType: enum.LoginBonusTypeNormal,
		BackgroundId:   100100700,
		WhiteboardTextureAsset: &client.TextureStruktur{
			V: generic.NewNullable(":7S"),
		},
		StartAt:                 0,
		EndAt:                   1<<31 - 1,
		LoginBonusHandler:       "beginner_login_bonus",
		LoginBonusHandlerConfig: "",
	})
	for day := 1; day <= 7; day++ {
		loginBonusRewardDays = append(loginBonusRewardDays, LoginBonusRewardDay{
			LoginBonusId: BeginnerLoginBonusId,
			Day:          int32(day),
			ContentGrade: enum.LoginBonusContentGradeNormal,
		})
	}
	loginBonusRewardContents = append(loginBonusRewardContents, LoginBonusRewardContent{
		LoginBonusId: BeginnerLoginBonusId,
		Day:          1,
		Content:      item.Gold.Amount(300000),
	})
	loginBonusRewardContents = append(loginBonusRewardContents, LoginBonusRewardContent{
		LoginBonusId: BeginnerLoginBonusId,
		Day:          2,
		Content:      item.SRScoutingTicket,
	})
	loginBonusRewardContents = append(loginBonusRewardContents, LoginBonusRewardContent{
		LoginBonusId: BeginnerLoginBonusId,
		Day:          3,
		Content:      item.EXP.Amount(300000),
	})
	loginBonusRewardContents = append(loginBonusRewardContents, LoginBonusRewardContent{
		LoginBonusId: BeginnerLoginBonusId,
		Day:          4,
		Content:      item.ShowCandy100,
	})
	loginBonusRewardContents = append(loginBonusRewardContents, LoginBonusRewardContent{
		LoginBonusId: BeginnerLoginBonusId,
		Day:          5,
		Content:      item.StarGem.Amount(50),
	})
	loginBonusRewardContents = append(loginBonusRewardContents, LoginBonusRewardContent{
		LoginBonusId: BeginnerLoginBonusId,
		Day:          6,
		Content:      item.MemoryKey.Amount(7),
	})
	loginBonusRewardContents = append(loginBonusRewardContents, LoginBonusRewardContent{
		LoginBonusId: BeginnerLoginBonusId,
		Day:          7,
		Content:      item.URScoutingTicket,
	})

	// normal login bonus
	loginBonuses = append(loginBonuses, LoginBonus{
		LoginBonusId:   NormalLoginBonusId,
		LoginBonusType: enum.LoginBonusTypeNormal,
		BackgroundId:   100100700,
		WhiteboardTextureAsset: &client.TextureStruktur{
			V: generic.NewNullable("/4n"),
		},
		StartAt:                 0,
		EndAt:                   1<<31 - 1,
		LoginBonusHandler:       "normal_login_bonus",
		LoginBonusHandlerConfig: "",
	})
	for day := 1; day <= 7; day++ {
		loginBonusRewardDays = append(loginBonusRewardDays, LoginBonusRewardDay{
			LoginBonusId: NormalLoginBonusId,
			Day:          int32(day),
			ContentGrade: enum.LoginBonusContentGradeNormal,
		})
	}
	loginBonusRewardContents = append(loginBonusRewardContents, LoginBonusRewardContent{
		LoginBonusId: NormalLoginBonusId,
		Day:          1,
		Content:      item.ShowCandy50,
	})
	loginBonusRewardContents = append(loginBonusRewardContents, LoginBonusRewardContent{
		LoginBonusId: NormalLoginBonusId,
		Day:          2,
		Content:      item.StarGem.Amount(10),
	})
	loginBonusRewardContents = append(loginBonusRewardContents, LoginBonusRewardContent{
		LoginBonusId: NormalLoginBonusId,
		Day:          3,
		Content:      item.TrainingTicket,
	})
	loginBonusRewardContents = append(loginBonusRewardContents, LoginBonusRewardContent{
		LoginBonusId: NormalLoginBonusId,
		Day:          4,
		Content:      item.StarGem.Amount(20),
	})
	loginBonusRewardContents = append(loginBonusRewardContents, LoginBonusRewardContent{
		LoginBonusId: NormalLoginBonusId,
		Day:          5,
		Content:      item.ShowCandy50,
	})
	loginBonusRewardContents = append(loginBonusRewardContents, LoginBonusRewardContent{
		LoginBonusId: NormalLoginBonusId,
		Day:          6,
		Content:      item.MemoryKey,
	})
	loginBonusRewardContents = append(loginBonusRewardContents, LoginBonusRewardContent{
		LoginBonusId: NormalLoginBonusId,
		Day:          7,
		Content:      item.StarGem.Amount(30),
	})

	// birthday login bonus
	loginBonuses = append(loginBonuses, LoginBonus{
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

func init() {
	addTable("s_login_bonus", LoginBonus{}, InitializeLoginBonus)
	addTable("s_login_bonus_reward_day", LoginBonusRewardDay{}, nil)
	addTable("s_login_bonus_reward_content", LoginBonusRewardContent{}, nil)
}
