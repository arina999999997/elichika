package response

import (
	"elichika/client"
)

type FetchLiveParntersResponse struct {
	PartnerSelectState client.PartnerSelectState `json:"partner_select_state"`
}
