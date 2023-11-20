package userdata

import (
	"elichika/model"
	"elichika/utils"

	// "fmt"
	"time"
)

func (session *Session) GetAllUserAccessories() []model.UserAccessory {
	accessories := []model.UserAccessory{}
	err := session.Db.Table("u_accessory").Where("user_id = ?", session.UserStatus.UserID).
		Find(&accessories)
	utils.CheckErr(err)
	return accessories
}

func (session *Session) GetUserAccessory(userAccessoryID int64) model.UserAccessory {
	// if exists then reuse
	pos, exist := session.UserAccessoryMapping.Map[userAccessoryID]
	if exist {
		return session.UserModel.UserAccessoryByUserAccessoryID.Objects[pos]
	}

	// if not look in db
	accessory := model.UserAccessory{}
	exists, err := session.Db.Table("u_accessory").
		Where("user_id = ? AND user_accessory_id = ?", session.UserStatus.UserID, userAccessoryID).Get(&accessory)
	utils.CheckErr(err)
	if !exists {
		// if not exists, create new one
		accessory = model.UserAccessory{
			UserID:             session.UserStatus.UserID,
			UserAccessoryID:    userAccessoryID,
			Level:              1,
			PassiveSkill1Level: 1,
			PassiveSkill2Level: 1,
			IsNew:              true,
			AcquiredAt:         time.Now().Unix(),
		}
	}
	return accessory
}

func (session *Session) UpdateUserAccessory(accessory model.UserAccessory) {
	session.UserAccessoryMapping.SetList(&session.UserModel.UserAccessoryByUserAccessoryID).Update(accessory)
}

func accessoryFinalizer(session *Session) {
	for _, accessory := range session.UserModel.UserAccessoryByUserAccessoryID.Objects {
		if accessory.IsNull {
			affected, err := session.Db.Table("u_accessory").
				Where("user_id = ? AND user_accessory_id = ?", session.UserStatus.UserID, accessory.UserAccessoryID).
				Delete(&accessory)
			utils.CheckErr(err)
			if affected != 1 {
				panic("accessory doesn't exists")
			}
		} else {
			affected, err := session.Db.Table("u_accessory").
				Where("user_id = ? AND user_accessory_id = ?", session.UserStatus.UserID, accessory.UserAccessoryID).
				AllCols().Update(accessory)
			utils.CheckErr(err)
			if affected == 0 {
				_, err := session.Db.Table("u_accessory").AllCols().Insert(accessory)
				utils.CheckErr(err)
			}
		}
	}

}

func init() {
	addFinalizer(accessoryFinalizer)
	addGenericTableFieldPopulator("u_accessory", "UserAccessoryByUserAccessoryID")
}
