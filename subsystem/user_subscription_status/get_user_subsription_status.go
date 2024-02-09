package user_subscription_status

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func GetUserSubsriptionStatus(session *userdata.Session, subscriptionMasterId int32) client.UserSubscriptionStatus {
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
