package userdata

import (
	"elichika/model"
	"elichika/utils"
)

func (session *Session) GetSubsriptionStatus() model.UserSubscriptionStatus {
	status := model.UserSubscriptionStatus{}
	exists, err := session.Db.Table("u_subscription_status").
		Where("user_id = ?", session.UserStatus.UserID).Get(&status)
	utils.CheckErr(err)
	if !exists {
		panic("subscription doesn't exist")
	}
	return status
}

func subscriptionStatusFinalizer(session *Session) {
	for _, userSubscriptionStatus := range session.UserModel.UserSubscriptionStatusByID.Objects {
		// userSubscriptionStatus.ExpireDate = (1 << 31) - 1 // patch it so we don't need to deal with expiration
		// userSubscriptionStatus.PlatformExpireDate = (1 << 31) - 1
		affected, err := session.Db.Table("u_subscription_status").
			Where("user_id = ? AND subscription_master_id = ?",
				session.UserStatus.UserID, userSubscriptionStatus.SubscriptionMasterID).
			AllCols().Update(userSubscriptionStatus)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_subscription_status").
				Insert(userSubscriptionStatus)
			utils.CheckErr(err)
		}
	}
}

func init() {
	addFinalizer(subscriptionStatusFinalizer)
	addGenericTableFieldPopulator("u_subscription_status", "UserSubscriptionStatusByID")
}
