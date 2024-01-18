package response

import (
	"elichika/client"
	"elichika/generic"
)

type FetchPresentResponse struct {
	PresentItems        generic.List[client.PresentItem]        `json:"present_items"`
	PresentHistoryItems generic.List[client.PresentHistoryItem] `json:"present_history_items"`
	PresentCount        int32                                   `json:"present_count"`
}
