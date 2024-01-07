package model

type UserSubscriptionStatus struct {
	UserId               int    `xorm:"pk 'user_id'" json:"-"`
	SubscriptionMasterId int    `xorm:"pk 'subscription_master_id'" json:"subscription_master_id"`
	StartDate            int    `xorm:"'start_date'" json:"start_date"`
	ExpireDate           int    `xorm:"'expire_date'" json:"expire_date"`
	PlatformExpireDate   int    `xorm:"'platform_expire_date'" json:"platform_expire_date"`
	RenewalCount         int    `xorm:"'renewal_count'" json:"renewal_count"`
	ContinueCount        int    `xorm:"'continue_count'" json:"continue_count"`
	SubscriptionPassId   int64  `xorm:"'subscription_pass_id'" json:"subscription_pass_id"`
	AttachId             string `xorm:"'attach_id'" json:"attach_id"`
	IsAutoRenew          bool   `xorm:"'is_auto_renew'" json:"is_auto_renew"`
	IsDoneTrial          bool   `xorm:"'is_done_trial'" json:"is_done_trial"`
}

func (uss *UserSubscriptionStatus) Id() int64 {
	return int64(uss.SubscriptionMasterId)
}

func init() {
	if TableNameToInterface == nil {
		TableNameToInterface = make(map[string]interface{})
	}
	TableNameToInterface["u_subscription_status"] = UserSubscriptionStatus{}
}
