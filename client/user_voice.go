package client

type UserVoice struct {
	NaviVoiceMasterId int32 `xorm:"pk 'navi_voice_master_id'" json:"navi_voice_master_id"`
	IsNew             bool  `xorm:"'is_new'" json:"is_new"`
}
