package client

type UserPlayList struct {
	UserPlayListId int32 `xorm:"pk 'user_play_list_id'" json:"user_play_list_id"`
	GroupNum       int32 `xorm:"'group_num'" json:"group_num"` // UserPlayListId % 10
	LiveId         int32 `xorm:"'live_id'" json:"live_id"`     // UserPlayListId / 10
}
