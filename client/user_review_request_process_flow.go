package client

type UserReviewRequestProcessFlow struct {
	ReviewRequestTriggerType int32 `xorm:"'review_request_trigger_type'" json:"review_request_trigger_type" enum:"ReviewRequestTriggerType"`
	ReviewRequestStatusType  int32 `xorm:"'review_request_status_type'" json:"review_request_status_type" enum:"ReviewRequestStatusType"`
}
