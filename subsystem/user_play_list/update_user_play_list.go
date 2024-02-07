package user_play_list

import (
	"elichika/client"
	"elichika/userdata"
)

func UpdateUserPlayList(session *userdata.Session, groupNum, liveMasterId int32, isSet bool) {
	id := liveMasterId*10 + groupNum
	if isSet {
		session.UserModel.UserPlayListById.Set(id, client.UserPlayList{
			UserPlayListId: id,
			GroupNum:       groupNum,
			LiveId:         liveMasterId,
		})
	} else {
		session.UserModel.UserPlayListById.SetNull(id)
	}
}
