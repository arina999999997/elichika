package response

import (
	"elichika/client"
)

// this is match client
type FetchLiveParntersResponse struct {
	PartnerSelectState client.PartnerSelectState `json:"partner_select_state"`
}
