package request

type BillingHistoryRequest struct {
	HistoryType int32 `json:"history_type" enum:"BillingHistoryType"`
	Page        int32 `json:"page"`
}
