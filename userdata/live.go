package userdata

import (
	"elichika/model"
	"elichika/utils"
)

func SaveLiveState(live model.LiveState) {
	// delete whatever is there
	_, err := Engine.Table("u_live_state").Where("user_id = ?", live.UserID).Delete(&model.LiveState{})
	if err != nil {
		panic(err)
	}
	affected, err := Engine.Table("u_live_state").AllCols().Insert(live)
	utils.CheckErr(err)
	if affected != 1 {
		panic("failed to insert")
	}
}

func LoadLiveState(userID int) (bool, model.LiveState) {
	live := model.LiveState{}
	exist, err := Engine.Table("u_live_state").Where("user_id = ?", userID).Get(&live)
	if err != nil {
		panic(err)
	}
	if exist {
		_, err = Engine.Table("u_live_state").Where("user_id = ?", userID).Delete(&model.LiveState{})
		if err != nil {
			panic(err)
		}
	}
	return exist, live
}
