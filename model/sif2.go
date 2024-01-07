package model

import (
	"elichika/generic"
)

type UserSif2DataLink struct {
	Sif2Id   int64  `xorm:"pk 'sif_2_id'" json:"sif_2_id"`
	Password string `xorm:"'password'" json:"password"`
}

func (usf2dl *UserSif2DataLink) Id() int64 {
	return int64(usf2dl.Sif2Id)
}

func init() {
	TableNameToInterface["u_sif_2_data_link"] = generic.UserIdWrapper[UserSif2DataLink]{}
}
