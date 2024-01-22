package client

import (
	"elichika/enum"

	"fmt"
)

type UserRecoveryAp struct {
	RecoveryApMasterId int32 `json:"recovery_ap_master_id"`
	Amount             int32 `json:"amount"`
}

func (ura *UserRecoveryAp) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeRecoveryAp { // 16
		panic(fmt.Sprintln("Wrong content for RecoveryAp: ", content))
	}
	ura.RecoveryApMasterId = content.ContentId
	ura.Amount = content.ContentAmount
}
func (ura *UserRecoveryAp) ToContent(contentId int32) Content {
	return Content{
		ContentType:   enum.ContentTypeRecoveryAp,
		ContentId:     contentId,
		ContentAmount: ura.Amount,
	}
}
