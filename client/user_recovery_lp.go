package client

import (
	"elichika/enum"
	"elichika/log"

	"fmt"
)

type UserRecoveryLp struct {
	RecoveryLpMasterId int32 `json:"recovery_lp_master_id"`
	Amount             int32 `json:"amount"`
}

func (url *UserRecoveryLp) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeRecoveryLp { // 17
		log.Panic(fmt.Sprintln("Wrong content for RecoveryLp: ", content))
	}
	url.RecoveryLpMasterId = content.ContentId
	url.Amount = content.ContentAmount
}
func (url *UserRecoveryLp) ToContent(contentId int32) Content {
	return Content{
		ContentType:   enum.ContentTypeRecoveryLp,
		ContentId:     contentId,
		ContentAmount: url.Amount,
	}
}
