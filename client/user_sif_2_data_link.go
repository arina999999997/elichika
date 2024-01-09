package client

type UserSif2DataLink struct {
	Sif2Id   int32  `xorm:"pk 'sif_2_id'" json:"sif_2_id"`
	Password string `xorm:"'password'" json:"password"`
}

func (us2dl *UserSif2DataLink) Id() int64 {
	return int64(us2dl.Sif2Id)
}
