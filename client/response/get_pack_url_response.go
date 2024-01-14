package response

import (
	"elichika/generic"
)

type GetPackUrlResponse struct {
	UrlList generic.List[string] `json:"url_list"`
}
