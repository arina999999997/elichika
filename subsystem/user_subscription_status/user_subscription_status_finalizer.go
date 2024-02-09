package user_subscription_status

import (
	"elichika/userdata"
	"elichika/utils"
)

func userSubscriptionStatusFinalizer(session *userdata.Session) {
	for _, userSubscriptionStatus := range session.UserModel.UserSubscriptionStatusById.Map {
		// userSubscriptionStatus.ExpireDate = (1 << 31) - 1 // patch it so we don't need to deal with expiration
		// userSubscriptionStatus.PlatformExpireDate = (1 << 31) - 1
		affected, err := session.Db.Table("u_subscription_status").
			Where("user_id = ? AND subscription_master_id = ?",
				session.UserId, userSubscriptionStatus.SubscriptionMasterId).
			AllCols().Update(*userSubscriptionStatus)
		utils.CheckErr(err)
		if affected == 0 {
			userdata.GenericDatabaseInsert(session, "u_subscription_status", *userSubscriptionStatus)
		}
	}
}

func init() {
	userdata.AddFinalizer(userSubscriptionStatusFinalizer)
}
