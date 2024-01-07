package userdata

import (
	"elichika/model"
	"elichika/utils"
)

func (session *Session) GetAllUserAccessories() []model.UserAccessory {
	accessories := []model.UserAccessory{}
	err := session.Db.Table("u_accessory").Where("user_id = ?", session.UserStatus.UserId).
		Find(&accessories)
	utils.CheckErr(err)
	return accessories
}

func (session *Session) GetUserAccessory(userAccessoryId int64) model.UserAccessory {
	// if exist then reuse
	pos, exist := session.UserAccessoryMapping.SetList(&session.UserModel.UserAccessoryByUserAccessoryId).Map[userAccessoryId]
	if exist {
		return session.UserModel.UserAccessoryByUserAccessoryId.Objects[pos]
	}

	// if not look in db
	accessory := model.UserAccessory{}
	exist, err := session.Db.Table("u_accessory").
		Where("user_id = ? AND user_accessory_id = ?", session.UserStatus.UserId, userAccessoryId).Get(&accessory)
	utils.CheckErr(err)
	if !exist {
		// if not exist, create new one
		accessory = model.UserAccessory{
			UserId:             session.UserStatus.UserId,
			UserAccessoryId:    userAccessoryId,
			Level:              1,
			PassiveSkill1Level: 1,
			PassiveSkill2Level: 1,
			IsNew:              true,
			AcquiredAt:         session.Time.Unix(),
		}
	}
	return accessory
}

func (session *Session) UpdateUserAccessory(accessory model.UserAccessory) {
	session.UserAccessoryMapping.SetList(&session.UserModel.UserAccessoryByUserAccessoryId).Update(accessory)
}

func accessoryFinalizer(session *Session) {
	for _, accessory := range session.UserModel.UserAccessoryByUserAccessoryId.Objects {
		if accessory.IsNull {
			affected, err := session.Db.Table("u_accessory").
				Where("user_id = ? AND user_accessory_id = ?", session.UserStatus.UserId, accessory.UserAccessoryId).
				Delete(&accessory)
			utils.CheckErr(err)
			if affected != 1 {
				panic("accessory doesn't exist")
			}
		} else {
			affected, err := session.Db.Table("u_accessory").
				Where("user_id = ? AND user_accessory_id = ?", session.UserStatus.UserId, accessory.UserAccessoryId).
				AllCols().Update(accessory)
			utils.CheckErr(err)
			if affected == 0 {
				_, err = session.Db.Table("u_accessory").AllCols().Insert(accessory)
				utils.CheckErr(err)
			}
		}
	}

}

func init() {
	addFinalizer(accessoryFinalizer)
	addGenericTableFieldPopulator("u_accessory", "UserAccessoryByUserAccessoryId")
}
