package gamedata

import (
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

func loadConstantInt(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading ConstantInt")
	type ConstantInt struct {
		Index int32 `xorm:"constant_int"`
		Value int32 `xorm:"value"`
	}
	constants := []ConstantInt{}

	err := masterdata_db.Table("m_constant_int").Find(&constants)
	utils.CheckErr(err)
	sz := int32(0)
	for _, c := range constants {
		for ; sz <= c.Index; sz++ {
			gamedata.ConstantInt = append(gamedata.ConstantInt, 0)
		}
		gamedata.ConstantInt[c.Index] = c.Value
	}
}

func init() {
	addLoadFunc(loadConstantInt)
}
