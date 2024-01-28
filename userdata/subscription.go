package userdata

import (
	"elichika/client"
	"elichika/utils"
)

func (session *Session) GetSubsriptionStatus(subscriptionMasterId int32) client.UserSubscriptionStatus {
	status := client.UserSubscriptionStatus{}
	exists, err := session.Db.Table("u_subscription_status").
		Where("user_id = ? AND subscription_master_id = ?", session.UserId, subscriptionMasterId).Get(&status)
	utils.CheckErr(err)
	if !exists {
		status = client.UserSubscriptionStatus{
			SubscriptionMasterId: subscriptionMasterId,
			StartDate:            session.Time.Unix(),
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
	for _, userSubscriptionStatus := range session.UserModel.UserSubscriptionStatusById.Map {
		// userSubscriptionStatus.ExpireDate = (1 << 31) - 1 // patch it so we don't need to deal with expiration
		// userSubscriptionStatus.PlatformExpireDate = (1 << 31) - 1
		affected, err := session.Db.Table("u_subscription_status").
			Where("user_id = ? AND subscription_master_id = ?",
				session.UserId, userSubscriptionStatus.SubscriptionMasterId).
			AllCols().Update(*userSubscriptionStatus)
		utils.CheckErr(err)
		if affected == 0 {
			GenericDatabaseInsert(session, "u_subscription_status", *userSubscriptionStatus)
		}
	}
}

func init() {
	AddContentFinalizer(subscriptionStatusFinalizer)
}
