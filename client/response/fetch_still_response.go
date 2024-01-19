package response

import (
	"elichika/generic"
)

type FetchStillResponse struct {
	NewStillList generic.List[int32] `json:"new_still_list"`
}
