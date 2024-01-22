package client

type UserSif2DataLink struct {
	Sif2Id   int32  `xorm:"pk 'sif_2_id'" json:"sif_2_id"`
	Password string `xorm:"'password'" json:"password"`
}
