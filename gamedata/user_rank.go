package gamedata

import (
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type UserRank struct {
	Rank                     int32 `xorm:"pk"`
	Exp                      int32
	MaxLp                    int32
	MaxAp                    int32
	FriendLimit              int32
	AdditionalAccessoryLimit int32
}

func loadUserRank(gamedata *Gamedata) {
	fmt.Println("Loading UserRank")
	gamedata.UserRank = make(map[int32]*UserRank)
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_user_rank").Find(&gamedata.UserRank)
	})
	utils.CheckErr(err)
	gamedata.UserRankMax = 0
	for _, userRank := range gamedata.UserRank {
		if userRank.Rank > gamedata.UserRankMax {
			gamedata.UserRankMax = userRank.Rank
		}
	}
}

func init() {
	addLoadFunc(loadUserRank)
}
