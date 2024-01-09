package client

import (
	"elichika/enum"

	"fmt"
)

type UserRecoveryLp struct {
	RecoveryLpMasterId int32 `json:"recovery_lp_master_id"`
	Amount             int32 `json:"amount"`
}

func (url *UserRecoveryLp) Id() int64 {
	return int64(url.RecoveryLpMasterId)
}
func (url *UserRecoveryLp) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeRecoveryLp { // 17
		panic(fmt.Sprintln("Wrong content for RecoveryLp: ", content))
	}
	url.RecoveryLpMasterId = content.ContentId
	url.Amount = content.ContentAmount
}
func (url *UserRecoveryLp) ToContent() Content {
	return Content{
		ContentType:   enum.ContentTypeRecoveryLp,
		ContentId:     url.RecoveryLpMasterId,
		ContentAmount: url.Amount,
	}
}
