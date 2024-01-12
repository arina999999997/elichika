package response

import (
	"elichika/client"
)

type BillingHistoryResponse struct {
	BillingBalanceHistoryList []client.BillingBalanceHistory `json:"billing_balance_history_list"`
	BillingDepositHistoryList []client.BillingDepositHistory `json:"billing_deposit_history_list"`
	BillingConsumeHistoryList []client.BillingConsumeHistory `json:"billing_consume_history_list"`
}
