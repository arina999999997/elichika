package gamedata

import (
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type ActivityPointRecoveryPrice struct {
	RecoveryCount int32 `xorm:"'recovery_count' pk"`
	Amount        int32 `xorm:"'amount'"`
}

// load into a prefix sum instead
func loadActivityPointRecoveryPrice(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading ActivityPointRecoveryPrice")
	err := masterdata_db.Table("m_activity_point_recovery_price").OrderBy("recovery_count").
		Find(&gamedata.ActivityPointRecoveryPrice)
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
