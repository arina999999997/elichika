package client

import (
	"elichika/generic"
)

// TODO(gacha): This is the correct struct but it's not used right now
type GachaUnconfirmed struct {
	Gacha       Gacha                              `json:"gacha"`
	RetryGacha  RetryGacha                         `json:"retry_gacha"`
	ResultCards generic.List[AddedGachaCardResult] `json:"result_cards"`
}
