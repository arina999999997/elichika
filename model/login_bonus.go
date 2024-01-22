package model

import (
	"elichika/client"
	// "elichika/generic"
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
