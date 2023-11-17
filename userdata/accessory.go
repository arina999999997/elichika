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
	accessory, exist := session.UserAccessoryDiffs[userAccessoryID]
	if exist {
		return accessory
	}

	// if not look in db
	accessory = model.UserAccessory{}
	exists, err := session.Db.Table("u_accessory").
		Where("user_id = ? AND user_accessory_id = ?", session.UserStatus.UserID, userAccessoryID).Get(&accessory)
	utils.CheckErr(err)
	if !exists {
		// if not exists, create new one
		accessory = model.UserAccessory{
			UserID:             session.UserStatus.UserID,
			UserAccessoryID:    userAccessoryID,
			Level:              1,
			Exp:                0,
			Grade:              0,
			Attribute:          0,
			PassiveSkill1ID:    0,
			PassiveSkill1Level: 1,
			PassiveSkill2ID:    nil,
			PassiveSkill2Level: 1,
			IsLock:             false,
			IsNew:              true,
			AcquiredAt:         time.Now().Unix(),
		}
	}
	return accessory
}

func (session *Session) UpdateUserAccessory(accessory model.UserAccessory) {
	session.UserAccessoryDiffs[accessory.UserAccessoryID] = accessory
}

func (session *Session) FinalizeUserAccessories() []any {
	accessoryByUserAccessoryID := []any{}
	for userAccessoryID, accessory := range session.UserAccessoryDiffs {
		accessoryByUserAccessoryID = append(accessoryByUserAccessoryID, userAccessoryID)
		session.UserModel.UserAccessoryByUserAccessoryID.PushBack(accessory)
		if accessory.IsNull {
			accessoryByUserAccessoryID = append(accessoryByUserAccessoryID, nil)
			affected, err := session.Db.Table("u_accessory").
				Where("user_id = ? AND user_accessory_id = ?", session.UserStatus.UserID, userAccessoryID).Delete(&accessory)
			utils.CheckErr(err)
			if affected != 1 {
				panic("accessories doesn't exists")
			}
		} else {
			accessoryByUserAccessoryID = append(accessoryByUserAccessoryID, accessory)
			affected, err := session.Db.Table("u_accessory").
				Where("user_id = ? AND user_accessory_id = ?", session.UserStatus.UserID, userAccessoryID).AllCols().Update(accessory)
			utils.CheckErr(err)
			if affected == 0 {
				_, err := session.Db.Table("u_accessory").AllCols().Insert(accessory)
				utils.CheckErr(err)
			}
		}
	}
	return accessoryByUserAccessoryID
}

func init() {
	addGenericTableFieldPopulator("u_accessory", "UserAccessoryByUserAccessoryID")
}
