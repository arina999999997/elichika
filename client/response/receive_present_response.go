package response

import (
	"elichika/client"
	"elichika/generic"
)

type ReceivePresentResponse struct {
	UserModelDiff        *client.UserModel                       `json:"user_model_diff"`
	PresentItems         generic.List[client.PresentItem]        `json:"present_items"`
	PresentHistoryItems  generic.List[client.PresentHistoryItem] `json:"present_history_items"`
	ReceivedPresentItems generic.List[client.Content]            `json:"received_present_items"`
	LimitExceededItems   generic.List[client.PresentItem]        `json:"limit_exceeded_items"`
	CardGradeUpResult    generic.List[client.AddedCardResult]    `json:"card_grade_up_result"`
	PresentCount         int32                                   `json:"present_count"`
}
