package model

import (
	"elichika/generic"
)

// not sure what this thing is about but let's save it anyway
type UserReviewRequestProcessFlow struct {
	ReviewRequestId          int64 `xorm:"pk 'review_request_id'" json:"-"`
	ReviewRequestTriggerType int   `xorm:"'review_request_trigger_type'" json:"review_request_trigger_type"`
	ReviewRequestStatusType  int   `xorm:"'review_request_status_type'" json:"review_request_status_type"`
}

func (urrpf *UserReviewRequestProcessFlow) Id() int64 {
	return urrpf.ReviewRequestId
}
func (urrpf *UserReviewRequestProcessFlow) SetId(id int64) {
	urrpf.ReviewRequestId = id
}

func init() {

	TableNameToInterface["u_review_request_process_flow"] = generic.UserIdWrapper[UserReviewRequestProcessFlow]{}
}
