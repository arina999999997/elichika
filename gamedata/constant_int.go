package gamedata

import (
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

func loadConstantInt(gamedata *Gamedata) {
	fmt.Println("Loading ConstantInt")
	type ConstantInt struct {
		Index int32 `xorm:"constant_int"`
		Value int32 `xorm:"value"`
	}
	constants := []ConstantInt{}

	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_constant_int").Find(&constants)
	})
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
