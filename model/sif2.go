package model

type UserSif2DataLink struct {
	UserID   int    `xorm:"pk 'user_id'" json:"-"`
	Sif2ID   int64  `xorm:"pk 'sif_2_id'" json:"sif_2_id"`
	Password string `xorm:"'password'" json:"password"`
}

func (usf2dl *UserSif2DataLink) ID() int64 {
	return int64(usf2dl.Sif2ID)
}

func init() {
	if TableNameToInterface == nil {
		TableNameToInterface = make(map[string]interface{})
	}
	TableNameToInterface["u_sif_2_data_link"] = UserSif2DataLink{}
}
