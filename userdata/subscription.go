package userdata

import (
	"elichika/model"
	"elichika/utils"
)

func (session *Session) GetSubsriptionStatus() model.UserSubscriptionStatus {
	status := model.UserSubscriptionStatus{}
	exists, err := session.Db.Table("u_subscription_status").
		Where("user_id = ?", session.UserStatus.UserId).Get(&status)
	utils.CheckErr(err)
	if !exists {
		status = model.UserSubscriptionStatus{
			UserId:               session.UserStatus.UserId,
			SubscriptionMasterId: 13001,
			StartDate:            int(session.Time.Unix()),
			ExpireDate:           1<<31 - 1,
			PlatformExpireDate:   1<<31 - 1,
			SubscriptionPassId:   session.Time.UnixNano(),
			AttachId:             "miraizura",
			IsAutoRenew:          true,
			IsDoneTrial:          true,
		}
	}
	return status
}

func subscriptionStatusFinalizer(session *Session) {
	for _, userSubscriptionStatus := range session.UserModel.UserSubscriptionStatusById.Objects {
		// userSubscriptionStatus.ExpireDate = (1 << 31) - 1 // patch it so we don't need to deal with expiration
		// userSubscriptionStatus.PlatformExpireDate = (1 << 31) - 1
		affected, err := session.Db.Table("u_subscription_status").
			Where("user_id = ? AND subscription_master_id = ?",
				session.UserStatus.UserId, userSubscriptionStatus.SubscriptionMasterId).
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
	addGenericTableFieldPopulator("u_subscription_status", "UserSubscriptionStatusById")
}
