package gamedata

import (
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type ActivityPointRecoveryPrice struct {
	RecoveryCount int32 `xorm:"'recovery_count' pk"`
	Amount        int32 `xorm:"'amount'"`
}

// load into a prefix sum instead
func loadActivityPointRecoveryPrice(gamedata *Gamedata) {
	fmt.Println("Loading ActivityPointRecoveryPrice")
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_activity_point_recovery_price").OrderBy("recovery_count").Find(&gamedata.ActivityPointRecoveryPrice)
	})
	utils.CheckErr(err)
	gamedata.ActivityPointRecoveryPrice = append([]ActivityPointRecoveryPrice{ActivityPointRecoveryPrice{}},
		gamedata.ActivityPointRecoveryPrice...)
	for i := range gamedata.ActivityPointRecoveryPrice {
		if i == 0 {
			continue
		}
		gamedata.ActivityPointRecoveryPrice[i].Amount += gamedata.ActivityPointRecoveryPrice[i-1].Amount
	}
}

func init() {
	addLoadFunc(loadActivityPointRecoveryPrice)
}
