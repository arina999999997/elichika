package userdata

import (
	"elichika/client"
	"elichika/generic"
	"elichika/utils"
)

func (session *Session) GetAllUserAccessories() []client.UserAccessory {
	accessories := []client.UserAccessory{}
	err := session.Db.Table("u_accessory").Where("user_id = ?", session.UserId).
		Find(&accessories)
	utils.CheckErr(err)
	return accessories
}

func (session *Session) GetUserAccessory(userAccessoryId int64) client.UserAccessory {
	ptr, exist := session.UserModel.UserAccessoryByUserAccessoryId.Get(userAccessoryId)
	if exist {
		return *ptr
	}
	accessory := client.UserAccessory{}
	// if not look in db
	exist, err := session.Db.Table("u_accessory").
		Where("user_id = ? AND user_accessory_id = ?", session.UserId, userAccessoryId).Get(&accessory)
	utils.CheckErr(err)
	if !exist {
		// if not exist, create new one
		accessory = client.UserAccessory{
			UserAccessoryId:    userAccessoryId,
			Level:              1,
			PassiveSkill1Level: generic.NewNullable(int32(1)),
			PassiveSkill2Level: generic.NewNullable(int32(1)),
			IsNew:              true,
			AcquiredAt:         session.Time.Unix(),
		}
	}
	return accessory
}

func (session *Session) UpdateUserAccessory(accessory client.UserAccessory) {
	session.UserModel.UserAccessoryByUserAccessoryId.Set(accessory.UserAccessoryId, accessory)
}

func (session *Session) DeleteUserAccessory(userAccessoryId int64) {
	session.UserModel.UserAccessoryByUserAccessoryId.SetNull(userAccessoryId)
}

func accessoryFinalizer(session *Session) {
	for accessoryId, accessory := range session.UserModel.UserAccessoryByUserAccessoryId.Map {
		if accessory != nil {
			affected, err := session.Db.Table("u_accessory").
				Where("user_id = ? AND user_accessory_id = ?", session.UserId, accessory.UserAccessoryId).
				AllCols().Update(*accessory)
			utils.CheckErr(err)
			if affected == 0 {
				GenericDatabaseInsert(session, "u_accessory", *accessory)
			}
		} else {
			affected, err := session.Db.Table("u_accessory").
				Where("user_id = ? AND user_accessory_id = ?", session.UserId, accessoryId).
				Delete(client.UserAccessory{})
			utils.CheckErr(err)
			if affected != 1 {
				panic("accessory doesn't exist")
			}
		}
	}

}

func init() {
	AddContentFinalizer(accessoryFinalizer)
}
