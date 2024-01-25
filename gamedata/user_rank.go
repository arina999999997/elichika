package gamedata

import (
	"elichika/dictionary"
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

func loadUserRank(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading UserRank")
	gamedata.UserRank = make(map[int32]*UserRank)
	err := masterdata_db.Table("m_user_rank").Find(&gamedata.UserRank)
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
